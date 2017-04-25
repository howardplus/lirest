package lirest

import (
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/config"
	"github.com/howardplus/lirest/describe"
	"github.com/howardplus/lirest/route"
	"net/http"
)

// Run starts the lirest server
func Run(path string) error {

	var defn describe.DescDefn

	// retrieve descriptions which tell us how to build the routes
	if err := describe.ReadDescriptionPath(path, &defn); err != nil {
		log.Fatal(err.Error())
		return err
	}

	// retrieve built-in descriptions for the /proc/sys directory
	describe.ReadSysctlDescriptions(&defn)

	// create a route trie for all the paths
	trie := route.NewTrie()
	defns := defn.DescriptionMap[describe.DescTypeStandard]
	defns = append(defns, defn.DescriptionMap[describe.DescTypeSysctl]...)
	for _, s := range defns {
		api := s.Api
		if err := trie.AddPath(api.Path, s); err != nil {
			log.WithFields(log.Fields{
				"path": api.Path,
			}).Error(err.Error())
			continue
		}
	}

	// TODO: exclude paths

	log.WithFields(log.Fields{
		"depth": trie.Depth(),
		"count": trie.Count(),
	}).Debug("Description path trie created")

	// create routes based on the trie
	r := route.NewRouter(trie)

	// all done, start the server
	log.Info("Running liRest server on " + config.GetConfig().Addr + ":" + config.GetConfig().Port)
	log.Fatal(http.ListenAndServe(config.GetConfig().Addr+":"+config.GetConfig().Port, r))

	return nil
}
