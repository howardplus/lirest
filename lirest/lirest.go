package lirest

import (
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/config"
	"github.com/howardplus/lirest/route"
	"github.com/howardplus/lirest/source"
	"net/http"
	_ "time"
)

type DescMsg struct {
	trie *route.Trie
	err  error
}

// Run starts the lirest server
func Run(path string, noSysctl bool, watch bool) error {

	routeChange := make(chan *DescMsg, 1)
	serverDone := make(chan int, 1)

	go DescriptionWatcher(routeChange, path, noSysctl, watch)

	go source.CacheManager()

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
				log.Info("Error with descriptions")
				return msg.err
			}

			// create a new server
			srv = &http.Server{
				Addr:    config.GetConfig().Addr + ":" + config.GetConfig().Port,
				Handler: route.NewRouter(msg.trie),
			}
			log.Info("Running liRest server on " + srv.Addr)
			go func() {
				log.Info(srv.ListenAndServe())
				serverDone <- 1
			}()
		}
	}

	return nil
}
