package cmd

import (
	"fmt"
	"github.com/howardplus/lirest/config"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "LiRest client commands",
	Long:  `LiRest's client commands`,
	RunE:  runClientCmd,
}

var jobCmd = &cobra.Command{
	Use:   "jobs",
	Short: "LiRest client commands",
	Long:  `LiRest's client commands`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("TODO: get jobs from rest api")
	},
}

func init() {
	clientCmd.PersistentFlags().StringVarP(&config.GetClientConfig().Addr, "ip", "i", "localhost", "IP address to connect to")
	clientCmd.PersistentFlags().StringVarP(&config.GetClientConfig().Port, "port", "p", "8080", "Port to connect on")

	clientCmd.AddCommand(jobCmd)
}

func runClientCmd(cmd *cobra.Command, args []string) error {
	return nil
}
