package cmd

import (
	"context"
	"log"
	"os"

	"github.com/gokul656/obscure-fs/internal/networking"
	"github.com/gokul656/obscure-fs/internal/storage"
	"github.com/spf13/cobra"
)

var (
	rootCmd    = &cobra.Command{Use: "obscure-fs"}
	network    *networking.Network // Shared network instance
	store      *storage.FileStore  // Shared storage instance
	ctx        = context.Background()
	listenPort int
	apiPort    int
	pkey       string

	bootstrapNodes = []string{
		"/ip4/127.0.0.1/tcp/5001/p2p/12D3KooWPvMQD2LhyeqGUsNK7GaJzA8N9XEvrbxqLPq21oBdUbfm",
		"/ip4/127.0.0.1/tcp/5002/p2p/12D3KooWGBGQhc2oMoRUBkKHHGrqQYBskTXMmWAgcwfSiLbAtE1g",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVar(&listenPort, "port", 0, "Port to listen on")
	rootCmd.PersistentFlags().IntVar(&apiPort, "api-port", 8080, "Port for the REST API")
	rootCmd.PersistentFlags().StringVar(&pkey, "pkey", "", "Private key path")

	rootCmd.MarkPersistentFlagRequired("port")
	rootCmd.MarkPersistentFlagRequired("api-port")
	rootCmd.MarkPersistentFlagRequired("pkey")
}
