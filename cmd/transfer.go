/*
Copyright © 2021 Dennis Thisner <dthisner@protonmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/dthisner/m3u-to-drive/cmd/transfer"
)

// transferCmd represents the root transfer command
var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Transfer your music",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

// startCmd will start transfering files
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start transfering your music",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("random called")
		if err := transfer.StartTransfer(); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(transferCmd)
	transferCmd.AddCommand(startCmd)
}
