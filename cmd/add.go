package cmd

import (
	"fmt"
	"log"

	"github.com/gokul656/obscure-fs/internal/hashing"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [file-path]",
	Short: "Add a file to the IPFS clone",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		cid, err := hashing.HashFile(filePath)
		if err != nil {
			log.Fatalf("Failed to hash file: %s\n", err)
		}

		err = store.StoreFile(filePath, cid)
		if err != nil {
			log.Fatalf("Failed to store file: %s\n", err)
		}

		fmt.Printf("File added with CID: %s\n", cid)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
