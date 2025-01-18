package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all stored files and their CIDs",
	Run: func(cmd *cobra.Command, args []string) {
		if len(store.ListFiles()) == 0 {
			fmt.Println("No files stored yet.")
			return
		}

		fmt.Println("Stored files:")
		for cid, path := range store.ListFiles() {
			fmt.Printf("CID: %s -> File: %s\n", cid, path)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
