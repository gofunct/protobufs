package cmd


import (
	"github.com/spf13/cobra"
	server "github.com/gofunct/grpclab/gen/user/server"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use: "serve",
	Short: "start a grpc and http server",
	Run: func(cmd *cobra.Command, args []string) {
		server.RunServer(defaultConfig, kitLog)
	},
}
