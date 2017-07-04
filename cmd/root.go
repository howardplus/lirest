package cmd

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/config"
	"github.com/howardplus/lirest/server"
	"github.com/howardplus/lirest/util"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(clientCmd)

	RootCmd.PersistentFlags().BoolVarP(&config.GetConfig().Verbose, "verbose", "v", false, "verbose output")
	RootCmd.PersistentFlags().BoolVarP(&config.GetConfig().Quiet, "quiet", "q", false, "quiet output")
	RootCmd.PersistentFlags().BoolVarP(&config.GetConfig().Watch, "watch", "w", false, "Watch for changes in description files")
	RootCmd.PersistentFlags().BoolVarP(&config.GetConfig().Pretty, "pretty", "y", false, "Pretty-print JSON output")

	RootCmd.Flags().BoolVarP(&config.GetConfig().NoSysctl, "no-sysctl", "s", false, "Disable sysctl routes")
	RootCmd.Flags().StringVarP(&config.GetConfig().Addr, "ip", "i", "localhost", "IP address to listen on")
	RootCmd.Flags().StringVarP(&config.GetConfig().Port, "port", "p", "8080", "Port to listen on")
	RootCmd.Flags().StringVarP(&config.GetConfig().DescPath, "desc-path", "d", "./descriptions/", "Description file path")
	RootCmd.Flags().StringVarP(&config.GetConfig().DescUrl, "desc-url", "u", "", "Download URL for description files")
}

// RootCmd define the root command
var RootCmd = &cobra.Command{
	Use:   config.ProjectName,
	Short: config.ProjectName + " exposes Linux operating system using REST API",
	Long:  config.ProjectName + " exposes Linux operating system using REST API",
	RunE:  runRootCmd,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of " + config.ProjectName,
	Long:  `All software has versions. This is ` + config.ProjectName + `'s`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s v%d.%d\n", config.ProjectName, util.GetVersion().Major, util.GetVersion().Minor)
	},
}

func runRootCmd(cmd *cobra.Command, args []string) error {

	conf := config.GetConfig()

	if conf.Verbose {
		log.SetLevel(log.DebugLevel)
	} else if conf.Quiet {
		log.SetLevel(log.FatalLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	if conf.DescUrl != "" {
		log.WithFields(log.Fields{
			"url": conf.DescUrl,
		}).Info("description download URL")

		if err := server.Download(conf.DescUrl, conf.DescPath+".lirest/"); err != nil {
			log.WithFields(log.Fields{
				"error": err.Error(),
			}).Error("description download failed")
		} else {
			server.Run(conf.DescPath+".lirest/", conf.NoSysctl, conf.Watch)
		}
	} else {
		server.Run(conf.DescPath, conf.NoSysctl, conf.Watch)
	}
	return nil
}
