package cmd

import (
	"context"
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
	debug   bool
	counter int
	src     string      // source file path
	dst     string      // destination path
	Runner  CommandLine = &command{}
	rootCmd             = &cobra.Command{
		Use:              "DB Backup verifier ",
		Short:            "DB Backup Verifier for HATCH",
		Long:             `DB Backup verifier comes into play when dba wants to make sure if their database backups are imported properly.`,
		Run:              Runner.run,
		TraverseChildren: true,
	}

	systemwideContext context.Context
	cancelFunc        context.CancelFunc
)

func init() {

	systemwideContext, cancelFunc = context.WithCancel(context.Background())
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(modeCmd)

	rootCmd.PersistentFlags().StringVarP(&src, "source", "s", "./sample/original.json", "set source path")
	rootCmd.PersistentFlags().StringVarP(&dst, "destination", "d", "./sample/duplicate.json", "set destination path")

}

func Execute() error {
	return rootCmd.Execute()
}

func (c *command) run(cmd *cobra.Command, args []string) {

	// start the app
	match, err := application.Start(src, dst,
		application.WithConsolePrint(),
		application.WithDiffSave("./new.txt"),
	)
	if err != nil {
		wlog.Fatalf("error in starting the app. Reason: %v\n", err.Error())
	}

	if !match {
		wlog.Fatalf("%s and %s are not matched\n", aurora.Red(src), aurora.Yellow(dst))
	}

}
