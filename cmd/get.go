package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [CID]",
	Short: "Retrieve a file by its CID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cid := args[0]

		filePath, err := store.GetFile(cid)
		if err != nil {
			log.Fatalf("Failed to retrieve file: %s\n", err)
		}

		fmt.Printf("File retrieved: %s\n", filePath)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
