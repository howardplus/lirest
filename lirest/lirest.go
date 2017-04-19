package lirest

import (
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/config"
	"github.com/howardplus/lirest/describe"
	"github.com/howardplus/lirest/route"
	"net/http"
)

func Run(path string) error {

	// retrieve descriptions which tell us how to build the routes
	desc, err := describe.ReadDescriptionPath(path)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	// create a route trie for all the paths
	trie := route.NewTrie()
	standards := desc.DescriptionMap[describe.DescTypeStandard]
	for _, s := range standards {
		output := s.Output
		if err := trie.AddPath(output.Path, s); err != nil {
			log.WithFields(log.Fields{
				"path": output.Path,
			}).Error(err.Error())
			continue
		}
	}

	log.WithFields(log.Fields{
		"depth": trie.Depth(),
		"count": trie.Count(),
	}).Debug("Description path trie created")

	// create routes based on the trie
	r := route.NewRouter(trie)

	// all done, start the server
	log.Info("Running liRest server on " + config.Config.Addr + ":" + config.Config.Port)
	log.Fatal(http.ListenAndServe(config.Config.Addr+":"+config.Config.Port, r))

	return nil
}
