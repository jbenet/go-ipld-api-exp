package ipld

import (
  "errors"

  block "github.com/ipfs/go-ipld/exp/block"
)

var (
  ErrParse = errors.New("ipld graph parse error")
  ErrNotFound = errors.New("ipld path not found")
)

type Node interface {
  // Block returns a low-level format serialized
  // version of this Node. Use this to dump out
  // an IPLD graph into disk, the wire, and so on.
  // Block() block.Block

  // Links returns a set of edges (path components)
  // or names of links. It is only the first links.
  Links() []string

  // GetNode returns a node by walking path p from
  // this node.
  GetNode(p Path) (Node, error)

  // Get unmarshals values from the underlying graph
  // into the given value val.
  // Get(p Path, val interface{}) (error)

  // UnmarshalTo unmarshals
  UnmarshalTo(interface{}) error
}

func GetValue(n Node, p Path, v interface) error {
  n2, err := n.GetNode(p)
  if err != nil {
    return err
  }
  return n2.UnmarshalTo(v)
}

func Subgraph(g Graph, p Path) Graph {

}

func AdjacentNodes(n Node) map[string]Node {
  m := map[string]Node{}
  for _, l := range n.Links() {
    m[l] = n.GetNode(l)
  }
  return m
}

func Advance(n Node, p Path) Node {
  return n.GetNode(p)
}

