package cmd

import (
	"log"

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

		if network == nil {
			log.Println("Initializing network...")
			if listenPort <= 0 {
				log.Fatalf("Invalid port: %d\n", listenPort)
			}
			network = networking.NewNetwork(ctx, listenPort)
			network.StartSimpleProtocol(utils.ProtocolID)
			log.Printf("Node ID: %s\n", network.GetHost().ID().String())
		}

		if store == nil {
			log.Println("Initializing file store...")
			store = storage.NewFileStore()
		}

		log.Printf("Node is listening on port %d. Press Ctrl+C to stop.\n", listenPort)
		select {} // Block forever
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

