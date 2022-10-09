// author: ashing
// time: 2020/7/12 2:44 下午
// mail: axingfly@gmail.com
// Less is more.

package cmd

import (
	"log"

	"github.com/ronething/grpc-sample/grpc-gateway-sample/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the gRPC hello-world server",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recover error : %v\n", err)
				return
			}
		}()

		server.Serve()
	},
}

func init() {
	serverCmd.Flags().StringVarP(&server.ServerPort, "port", "p", "50052", "server port")
	serverCmd.Flags().StringVarP(&server.CertPemPath, "cert-pem", "", "./certs/server.pem", "cert pem path")
	serverCmd.Flags().StringVarP(&server.CertKeyPath, "cert-key", "", "./certs/server.key", "cert key path")
	serverCmd.Flags().StringVarP(&server.CertName, "cert-name", "", "localhost", "server's hostname")
	rootCmd.AddCommand(serverCmd)
}
