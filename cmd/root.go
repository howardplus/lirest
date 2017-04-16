package cmd

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/config"
	"github.com/howardplus/lirest/lirest"
	"github.com/howardplus/lirest/util"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)

	RootCmd.PersistentFlags().BoolVarP(&config.Config.Verbose, "verbose", "v", false, "verbose output")
	RootCmd.PersistentFlags().BoolVarP(&config.Config.Quiet, "quiet", "q", false, "quiet output")
	RootCmd.PersistentFlags().StringVarP(&config.Config.Addr, "ip", "i", "localhost", "IP address to listen on")
	RootCmd.PersistentFlags().StringVarP(&config.Config.Port, "port", "p", "8080", "Port to listen on")
	RootCmd.PersistentFlags().StringVarP(&config.Config.DescPath, "desc-path", "d", "./describe/", "Description file path")
}

var RootCmd = &cobra.Command{
	Use:   "lirest",
	Short: "liRest exposes Linux operating system using REST API",
	Long:  "liRest exposes Linux operating system using REST API",
	RunE:  RunRootCmd,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of liRest",
	Long:  `All software has versions. This is liRest's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("liRest v%d.%d\n", util.Version.Major, util.Version.Minor)
	},
}

func RunRootCmd(cmd *cobra.Command, args []string) error {
	if config.Config.Verbose {
		log.SetLevel(log.DebugLevel)
	} else if config.Config.Quiet {
		log.SetLevel(log.FatalLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	lirest.Run()
	return nil
}
