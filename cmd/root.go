package cmd

import (
	"log"
	"os"

	"github.com/d-fal/bverify/application"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var wlog = log.New(os.Stdout, "[ cmd ] ", log.Lshortfile|log.Ltime)

type CommandLine interface {
	version(cmd *cobra.Command, args []string)
	run(cmd *cobra.Command, args []string)
	mode(cmd *cobra.Command, args []string)
}
type command struct{}

var (
	debug        bool
	src          string      // source file path
	dst          string      // destination path
	diffFilePath string      // diff file path
	Runner       CommandLine = &command{}
	rootCmd                  = &cobra.Command{
		Use:              "DB Backup verifier ",
		Short:            "DB Backup Verifier for HATCH",
		Long:             `DB Backup verifier comes into play when dba wants to make sure if their database backups are imported properly.`,
		Run:              Runner.run,
		TraverseChildren: true,
	}
)

func init() {

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(modeCmd)

	rootCmd.PersistentFlags().StringVarP(&src, "source", "s", "./sample/original.json", "set source path")
	rootCmd.PersistentFlags().StringVarP(&dst, "destination", "d", "./sample/duplicate.json", "set destination path")
	rootCmd.PersistentFlags().StringVarP(&diffFilePath, "diff-path", "f", "", "set destination diff path")

}

func Execute() error {
	return rootCmd.Execute()
}

func (c *command) run(cmd *cobra.Command, args []string) {

	var opts []application.Opt

	opts = append(opts, application.WithConsolePrint())

	if diffFilePath != "" {
		opts = append(opts, application.WithDiffSave(diffFilePath))
	}

	// start the app
	match, err := application.Start(src, dst,
		opts...,
	)
	if err != nil {
		wlog.Fatalf("error in starting the app. Reason: %v\n", err.Error())
	}

	if !match {
		wlog.Printf("%s and %s are not matched\n", aurora.Red(src), aurora.Yellow(dst))
	}

}
