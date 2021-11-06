package cmd

import (
	"github.com/spf13/cobra"
)

var (
	modeCmd = &cobra.Command{
		Use:   "mode",
		Short: "debug mode [true|false] default is false",
		Run:   Runner.mode,
	}
)

func init() {
	modeCmd.Flags().BoolVarP(&debug, "debug", "v", false, "verbose mode")
}

func (c *command) mode(cmd *cobra.Command, args []string) {
	c.run(cmd, args)
}
