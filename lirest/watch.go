package lirest

import (
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/describe"
	"github.com/howardplus/lirest/route"
	"github.com/howardplus/lirest/util"
	"github.com/howeyc/fsnotify"
)

func watchDescriptionChange(path string) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	err = watcher.Watch(path)
	if err != nil {
		return err
	}

	ev := <-watcher.Event
	log.WithFields(log.Fields{
		"event": ev,
	}).Debug("got event")
	return nil
}

// DescriptionWatcher monitors changes on description
func DescriptionWatcher(ch chan *DescMsg, path string, noSysctl bool, watch bool) {
	for {
		if trie, err := generateRouteTrie(path, noSysctl); err != nil {
			ch <- &DescMsg{err: err}
		} else {
			ch <- &DescMsg{trie: trie}
		}

		// not watching, we do it once and done
		if watch == false {
			return
		}

		// otherwise wait for changes
		if err := watchDescriptionChange(path); err != nil {
			ch <- &DescMsg{err: err}
		}
	}
}

func generateRouteTrie(path string, noSysctl bool) (*route.Trie, error) {
	var defn describe.DescDefn

	// retrieve descriptions which tell us how to build the routes
	if err := describe.ReadDescriptionPath(path, &defn); err != nil {
		log.Fatal(err.Error())
		return nil, err
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
		for _, v := range s.Api.Vars {
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

	log.WithFields(log.Fields{
		"depth": trie.Depth(),
		"count": trie.Count(),
	}).Debug("Description path trie created")

	return trie, nil
}
