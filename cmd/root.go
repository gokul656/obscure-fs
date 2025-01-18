package cmd

import (
	"context"
	"log"
	"os"

	"github.com/gokul656/obscure-fs/internal/networking"
	"github.com/gokul656/obscure-fs/internal/storage"
	"github.com/gokul656/obscure-fs/utils"
	"github.com/spf13/cobra"
)

var (
	rootCmd    = &cobra.Command{Use: "obscure-fs"}
	network    *networking.Netowrk // Shared network instance
	store      *storage.FileStore  // Shared storage instance
	ctx        = context.Background()
	listenPort int
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVar(&listenPort, "port", 0, "Port to listen on")
	rootCmd.MarkPersistentFlagRequired("port")

	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		log.Println("Starting PersistentPreRun...")

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

		log.Println("PersistentPreRun complete.")
	}
}
