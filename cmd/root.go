package cmd

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/config"
	"github.com/howardplus/lirest/lirest"
	"github.com/howardplus/lirest/util"
	"github.com/spf13/cobra"
)

// flags
var Verbose bool
var Quiet bool

func init() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(portCmd)
	RootCmd.AddCommand(ipCmd)

	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	RootCmd.PersistentFlags().BoolVarP(&Quiet, "quiet", "q", false, "quiet output")
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

var portCmd = &cobra.Command{
	Use:   "port",
	Short: "listening port of liRest",
	Long:  `listening port of liRest`,
	Run: func(cmd *cobra.Command, args []string) {
		config.Config.Port = args[0]
	},
}

var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "IP of liRest",
	Long:  `IP of liRest`,
	Run: func(cmd *cobra.Command, args []string) {
		config.Config.Addr = args[0]
	},
}

func RunRootCmd(cmd *cobra.Command, args []string) error {
	if Verbose {
		log.SetLevel(log.DebugLevel)
	} else if Quiet {
		log.SetLevel(log.FatalLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	log.Info("Running liRest in standalone mode")
	lirest.Run()
	return nil
}
