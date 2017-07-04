package route

import (
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/util"
	"strings"
)

// Trie is a data structure to keep the route relationships
// this is created from the descriptions
// and used to create routes
type Trie struct {
	Val   interface{}
	Nodes map[string]*Trie
}

const (
	routeDefaultCount = 5
)

// NewTrie create the trie root
func NewTrie() *Trie {
	return &Trie{
		Nodes: make(map[string]*Trie, routeDefaultCount)}
}

// AddPath add a path to the Trie
// a path contains slash separated strings such as
// /a/b/c/d/e
func (t *Trie) AddPath(path string, val interface{}) error {
	log.Debug("Adding path ", path)
	tokens := strings.Split(path, "/")

	if len(tokens) == 0 {
		return util.NewError("Unknown path")
	}

	n := t
	for i, key := range tokens {
		// empty path, can be either "//", which we omit
		// or the front of the path, which is not a key
		if key == "" {
			continue
		}

		// on leaf node, assign value
		var v interface{}
		if i == len(tokens)-1 {
			v = val
		}

		if elem, found := n.Nodes[key]; found == true {
			n = elem
			if v != nil {
				// duplicate leaf node
				return util.NewError("Duplicate path")
			}
		} else {
			// not found, create it
			n.Nodes[key] = &Trie{
				Val:   v,
				Nodes: make(map[string]*Trie, routeDefaultCount)}
			n = n.Nodes[key]
		}
	}

	return nil
}

func (t *Trie) depthN(depth int) int {
	// no more nodes after this
	if len(t.Nodes) == 0 {
		return depth
	}

	// find the maximum of all paths
	max := depth
	for _, v := range t.Nodes {
		tmp := v.depthN(depth + 1)
		if tmp > max {
			max = tmp
		}
	}

	return max
}

// Depthfind the depth of the trie
func (t *Trie) Depth() int {
	return t.depthN(0)
}

// Count counts the number of nodes in the trie
func (t *Trie) Count() int {
	i := 1
	for _, v := range t.Nodes {
		i += v.Count()
	}
	return i
}
