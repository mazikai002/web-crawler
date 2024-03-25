package master

import (
	"net/http"

	"github.com/spf13/cobra"
)

var MasterCmd = &cobra.Command{
	Use:   "master",
	Short: "run master service.",
	Long:  "run master service.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

var masterID string
var cfgFile string
var HTTPListenAddress string
var GRPCListenAddress string
var PProfListenAddress string
var podIP string

func init() {
	MasterCmd.Flags().StringVar(
		&masterID, "id", "", "set master id")
	MasterCmd.Flags().StringVar(
		&cfgFile, "config", "config.toml", "set master id")
	MasterCmd.Flags().StringVar(
		&HTTPListenAddress, "http", ":8081", "set HTTP listen address")
	MasterCmd.Flags().StringVar(
		&podIP, "podip", "", "set worker id")
	MasterCmd.Flags().StringVar(
		&GRPCListenAddress, "grpc", ":9091", "set GRPC listen address")
	MasterCmd.Flags().StringVar(
		&PProfListenAddress, "pprof", ":9981", "set pprof address")
}

func Run() {
	// start pprof，换成gin框架试试看
	go func() {
		if err := http.ListenAndServe(PProfListenAddress, nil); err != nil {
			panic(err)
		}
	}()

	// load config

	// log init

	// init tasks

	// start http proxy to GRPC

	// start grpc server
}
