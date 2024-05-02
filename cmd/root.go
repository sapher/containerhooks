package cmd

import (
	"log"
	"os"

	"github.com/sapher/containerhooks/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "containerhooks",
	Short: "Tool to run githooks isolated in containers",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(onInitialize)
}

func onInitialize() {
	// Load configuration
	_, err := internal.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}
}
