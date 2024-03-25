package cmd

import (
	"github.com/mazikai/web-crawler/cmd/master"
	"github.com/mazikai/web-crawler/cmd/worker"

	"github.com/spf13/cobra"
)

func Execute() {
	var rootCmd = &cobra.Command{Use: "web-crawler"}
	rootCmd.AddCommand(master.MasterCmd, worker.WorkerCmd)
	rootCmd.Execute()
}
