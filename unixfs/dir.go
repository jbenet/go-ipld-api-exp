package unixfs

import (
  "errors"

  ipld "github.com/ipfs/go-ipld/exp/ipld"
)

var (
  ErrNoEntry = errors.New("no such directory entry")
)

type Dir interface {
  ipld.Node

  Entries() (map[string]DirEntry, error)
}

type dir struct {
  entries map[string]DirEntry
}

type DirEntry struct {
  link ipld.Link // link to the content of the entry
  size int       // total size of the file or dir
}

func newDir(n dag.Node) (*dir, error) {
  d := &dir{}
  err := n.UnmarshalTo(d)
  if err != nil {
    return nil, err
  }
  return d, nil
}

func (d *dir) Entry(name string) (d DirEntry, err error) {
  d = d.entries[name]
  if d == nil {
    err = ErrNoEntry
  }
  return d, err
}

func (d *dir) Entries() map[string]DirEntry {
  // make a copy
  ents := map[string]DirEntry{}
  for n, e := range d.entries {
    ents[n] = e
  }
  return ents
}

// Implement the ipld.Node interface

// Links returns a set of edges (path components)
// or names of links. It is only the first links.
func (d *dir) Links() []string {
  ls := make([]string, len(d.entries))
  for n, _ := d.entries {
    ls = append(ls, n)
  }
  return ls
}

// GetNode returns a node by walking path p from
// this node.
func (d *dir) GetNode(p Path) (Node, error) {
  if p == "" {
    return d
  }

  s := PathSplit(p)
  e, found := d.entries[s[0]]
  if !found {
    return nil, ipld.NotFound
  }
  return e.link.GetNode(s[1:])
}

// UnmarshalTo unmarshals
func (d *dir) UnmarshalTo(v interface{}) error {
  switch v := v.(type) {
  case []string: // list of filenames
    v = d.Links()
    return nil
  case map[string]DirEntry, interface{}:
    v = d.Entries()
    return nil
  default:
    return ipld.ErrParse
  }
}
