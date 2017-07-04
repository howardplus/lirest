package util

import (
	"bytes"
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/config"
	"io"
	"io/ioutil"
	"net/http"
)

// Get gets cmd via rest
func Get(cmd string) (io.Reader, error) {
	addr := config.GetClientConfig().Addr
	port := config.GetClientConfig().Port

	resp, err := http.Get("http://" + addr + ":" + port + "/cmd/" + cmd)
	if err != nil {
		log.WithFields(log.Fields{
			"cmd": cmd,
		}).Error(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"cmd": cmd,
		}).Error(err.Error())
		return nil, err
	}

	return bytes.NewReader(body), nil
}

// Put sets data via rest
func Put() {
}
