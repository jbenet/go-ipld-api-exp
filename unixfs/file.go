package unixfs

import (
  ipld "github.com/ipfs/go-ipld/exp/ipld"
)

type File interface {
  Len()    int
  Reader() ReadSeekCloser
}

func Open(p Path) File { ... }


type file struct {
  data     []byte
  subfiles []fileLink
}

type fileLink struct {
  link   dag.Link
  length int
}

func newFile(n dag.Node) (*file, error) {
  f := &file{}
  err := n.UnmarshalTo(f)
  if err != nil {
    return nil, err
  }
  return f, nil
}

func (f *file) Len() int {
  total := len(f.data)
  for _, s := range f.subfiles {
    total += s.length
  }
  return total
}

func (f *file) Subfile(i int) (f2 *file, err error) {
  if i < 0 || i >= len(f.subfiles) {
    return nil, ipld.ErrParse // invalid.
  }
  return newFile(f.subfiles[i].GetNode())
}

func (f *file) Reader() Reader {
  return &Reader{f}
}

// Implement the ipld.Node interface

// Links returns a set of edges (path components)
// or names of links. It is only the first links.
func (f *file) Links() []string {
  return nil // no links.
}

// GetNode returns a node by walking path p from
// this node.
func (f *file) GetNode(p Path) (Node, error) {
  if p == "" {
    return f
  }

  return nil, ipld.ErrParse
}

// UnmarshalTo
func (f *file) UnmarshalTo(v interface{}) error {
  switch v := v.(type) {
  case []byte:
    _, err := io.ReadFull(f, v)
    return err

  default:
    return ipld.ErrParse
  }
}

// UnmarshalBlock
func (f *file) UnmarshalBlock(b block.Block) error {
  return ipld.UnmarshalBlockTo(b, f)
}
