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

	RootCmd.PersistentFlags().BoolVarP(&config.GetConfig().NoSysctl, "no-sysctl", "s", false, "Disable sysctl routes")
	RootCmd.PersistentFlags().BoolVarP(&config.GetConfig().Verbose, "verbose", "v", false, "verbose output")
	RootCmd.PersistentFlags().BoolVarP(&config.GetConfig().Quiet, "quiet", "q", false, "quiet output")
	RootCmd.PersistentFlags().StringVarP(&config.GetConfig().Addr, "ip", "i", "localhost", "IP address to listen on")
	RootCmd.PersistentFlags().StringVarP(&config.GetConfig().Port, "port", "p", "8080", "Port to listen on")
	RootCmd.PersistentFlags().StringVarP(&config.GetConfig().DescPath, "desc-path", "d", "./descriptions/", "Description file path")
	RootCmd.PersistentFlags().BoolVarP(&config.GetConfig().Watch, "watch", "w", false, "Watch for changes in description files")
}

// Define the root command
var RootCmd = &cobra.Command{
	Use:   "lirest",
	Short: "liRest exposes Linux operating system using REST API",
	Long:  "liRest exposes Linux operating system using REST API",
	RunE:  runRootCmd,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of liRest",
	Long:  `All software has versions. This is liRest's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("liRest v%d.%d\n", util.GetVersion().Major, util.GetVersion().Minor)
	},
}

func runRootCmd(cmd *cobra.Command, args []string) error {
	if config.GetConfig().Verbose {
		log.SetLevel(log.DebugLevel)
	} else if config.GetConfig().Quiet {
		log.SetLevel(log.FatalLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	lirest.Run(config.GetConfig().DescPath, config.GetConfig().NoSysctl, config.GetConfig().Watch)
	return nil
}
