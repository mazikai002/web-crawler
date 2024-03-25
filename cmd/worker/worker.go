package worker

import "github.com/spf13/cobra"

var ServiceName string = "go.micro.server.worker"

var WorkerCmd = &cobra.Command{
	Use:   "worker",
	Short: "run worker service.",
	Long:  "run worker service.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

var cluster bool
var workerID string
var HTTPListenAddress string
var GRPCListenAddress string
var PProfListenAddress string
var podIP string

func init() {
	WorkerCmd.Flags().StringVar(
		&workerID, "id", "", "set worker id")

	WorkerCmd.Flags().StringVar(
		&podIP, "podip", "", "set worker id")

	WorkerCmd.Flags().StringVar(
		&HTTPListenAddress, "http", ":8080", "set HTTP listen address")

	WorkerCmd.Flags().StringVar(
		&GRPCListenAddress, "grpc", ":9090", "set GRPC listen address")

	WorkerCmd.Flags().StringVar(
		&PProfListenAddress, "pprof", ":9981", "set pprof address")

	WorkerCmd.Flags().BoolVar(
		&cluster, "cluster", true, "run mode")
}

func Run() {

}
