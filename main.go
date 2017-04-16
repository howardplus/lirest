package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/cmd"

	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(-1)
	}
}
