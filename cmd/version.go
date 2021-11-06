package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Generate version",
	Run:   Runner.version,
}

func (c *command) version(cmd *cobra.Command, args []string) {
	fmt.Println("App version : 0.0.1")
}
