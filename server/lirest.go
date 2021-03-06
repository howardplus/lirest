package server

import (
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/config"
	"github.com/howardplus/lirest/inject"
	"github.com/howardplus/lirest/route"
	"github.com/howardplus/lirest/source"
	"io"
	"net/http"
	"os"
)

// DescMsg contains the route trie that have been modified or created
// to generate the new api routes
type DescMsg struct {
	trie *route.Trie
	err  error
}

const (
	localFilename = "lirest.des"
)

// Download description files from URL
func Download(url string, path string) error {
	resp, err := http.Get(url)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
			"url":   url,
		}).Error("http get error")
		return err
	}
	defer resp.Body.Close()

	// TODO: support more than single file?

	// remove files from the path
	os.Remove(path + localFilename)

	if err := os.MkdirAll(path, 0755); err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
			"path":  path,
		}).Error("path create error")
		return err
	}

	out, err := os.Create(path + localFilename)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
			"path":  path + localFilename,
		}).Error("file create error")
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		log.Error("io error")
		return err
	}

	return nil
}

// Run starts the lirest server
func Run(path string, noSysctl bool, watch bool) {

	routeChange := make(chan *DescMsg, 1)
	serverDone := make(chan int, 1)

	go DescriptionWatcher(routeChange, path, noSysctl, watch)

	go source.CacheManager()

	go inject.JobTracker()

	// wait and handle all the channel messagas
	var srv *http.Server
	for {
		select {
		case msg := <-routeChange:
			if srv != nil {
				log.Info("Shutting down server...")

				srv.Shutdown(nil)
				// now wait for server to shut down
				<-serverDone
			}

			if msg.err != nil {
				log.Error("Error with descriptions")
			} else {
				// create a new server
				srv = &http.Server{
					Addr:    config.GetConfig().Addr + ":" + config.GetConfig().Port,
					Handler: route.NewRouter(msg.trie),
				}
				log.Info("Running " + config.ProjectName + " server on " + srv.Addr)
				go func() {
					log.Info(srv.ListenAndServe())
					serverDone <- 1
				}()
			}
		}
	}
}
