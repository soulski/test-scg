package main

import (
	"os"

	"github.com/soulski/test-scg/pkg/scg"
	"github.com/spf13/cobra"
)

func main() {
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start server",

		Run: func(cmd *cobra.Command, args []string) {
			server := scg.NewServer()
			server.Start()
		},
	}

	rootCmd := &cobra.Command{Use: "scg"}
	rootCmd.AddCommand(startCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
