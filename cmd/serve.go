package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/gokul656/obscure-fs/internal/networking"
	"github.com/gokul656/obscure-fs/internal/storage"
	"github.com/gokul656/obscure-fs/utils"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the node and start listening for connections",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Starting the node...")

		if store == nil {
			log.Println("Initializing file store...")
			store = storage.NewFileStore()
			log.Println("Sucessfully initialzied file store...")
		}

		if network == nil {
			log.Println("Initializing network...")
			if listenPort <= 0 {
				log.Fatalf("Invalid port: %d\n", listenPort)
			}
			network = networking.NewNetwork(ctx, listenPort, pkey, bootstrapNodes, store)
			network.StartSimpleProtocol(utils.ProtocolID)
			log.Printf("Node ID: %s\n", network.GetHost().ID().String())
			network.ConnectToBootstrapNodes()
			network.AnnounceToPeers(network.GetHost().ID().String(), fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", listenPort))
		}

		log.Printf("Node is listening on port %d. Press Ctrl+C to stop.\n", listenPort)

		registry := networking.NewNodeRegistry()
		router := gin.Default()
		gin.SetMode(gin.ReleaseMode)

		router.Use(func(c *gin.Context) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			c.Next()
		})

		router.POST("/nodes/register", func(c *gin.Context) {
			var node networking.Node
			if err := c.ShouldBindJSON(&node); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
				return
			}

			registry.RegisterNode(node)
			c.JSON(http.StatusOK, gin.H{"message": "Node registered successfully", "node_id": node.ID})
		})

		router.GET("/nodes", func(c *gin.Context) {
			nodes := registry.GetAllNodes()
			c.JSON(http.StatusOK, gin.H{"nodes": nodes})
		})

		router.POST("/files/upload", func(c *gin.Context) {
			file, err := c.FormFile("file")
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload file"})
				return
			}

			uploadDir := "./uploads"
			filePath := fmt.Sprintf("%s/%s", uploadDir, file.Filename)
			if err := c.SaveUploadedFile(file, filePath); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
				return
			}

			cid, err := network.ShareFile(filePath)
			if err != nil {
				return
			}

			log.Printf("file uploaded: %s (CID: %s)\n", filePath, cid)
			c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "cid": cid})
		})

		router.GET("/files/:cid", func(c *gin.Context) {
			cid := c.Param("cid")

			tempFilePath := fmt.Sprintf("./temp/%s", cid)

			err := network.RetrieveFile(cid, tempFilePath)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
				return
			}

			c.File(tempFilePath)
		})

		go func() {
			if err := router.Run(fmt.Sprintf(":%d", apiPort)); err != nil {
				log.Fatalf("Failed to start HTTP server: %v", err)
			}
		}()

		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, os.Interrupt)
		<-shutdown

		log.Println("Shutting down...")
		if err := network.Shutdown(); err != nil {
			log.Printf("Failed to shut down network: %v", err)
		}
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
