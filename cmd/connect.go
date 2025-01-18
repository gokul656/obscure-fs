package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect [multiaddress]",
	Short: "Connect to a peer using its multiaddress",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := network.ConnectToPeer(args[0])
		if err != nil {
			log.Fatalf("Failed to connect to peer: %s\n", err)
		}
		fmt.Println("Successfully connected to the peer.")
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
}
