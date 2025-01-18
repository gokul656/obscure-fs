package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listPeersCmd = &cobra.Command{
	Use:   "list-peers",
	Short: "List all connected peers",
	Run: func(cmd *cobra.Command, args []string) {
		for _, peer := range network.GetHost().Peerstore().Peers() {
			fmt.Printf("Peer ID: %s\n", peer.String())
		}
	},
}

func init() {
	rootCmd.AddCommand(listPeersCmd)
}
