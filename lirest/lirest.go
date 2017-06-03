package lirest

import (
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/config"
	"github.com/howardplus/lirest/describe"
	"github.com/howardplus/lirest/route"
	"github.com/howardplus/lirest/util"
	"net/http"
	_ "strings"
)

// Run starts the lirest server
func Run(path string, noSysctl bool) error {

	var defn describe.DescDefn

	// retrieve descriptions which tell us how to build the routes
	if err := describe.ReadDescriptionPath(path, &defn); err != nil {
		log.Fatal(err.Error())
		return err
	}

	// create a route trie for all the paths
	trie := route.NewTrie()
	defns := defn.DescriptionMap[describe.DescTypeStandard]

	if noSysctl == false {
		// retrieve built-in descriptions for the /proc/sys directory
		describe.ReadSysctlDescriptions(&defn)
		defns = append(defns, defn.DescriptionMap[describe.DescTypeSysctl]...)
	}

	for _, s := range defns {
		api := s.Api

		vars := make(map[string]string, 0)
		for _, v := range s.Vars {
			vars[v.Name] = v.DataType
		}

		path := util.PathAddType(api.Path, vars)
		if err := trie.AddPath(path, s); err != nil {
			log.WithFields(log.Fields{
				"path": path,
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
