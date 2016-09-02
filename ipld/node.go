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

func Block(n Node) (block.Block, error) {
  if nb, ok := n.(block.BlockMarshaler); ok {
    return nb.MarshalBlock()
  }

  // otherwise, do it with "sort of reflection"
  // with a default serialization type: cbor
  return cbor.MarshalBlock(n)
}

func UnmarshalBlockTo(b block.Block, v interface{}) error {
  var n Node
  err := block.Unmarshal(b, n)
  if err != nil {
    return err
  }

  return n.UnmarshalTo(v)
}

func MarshalBlockFrom(v interface{}) (block.Block, error) {
  if mv, ok := v.(block.BlockMarshaler); ok {
    return mv.MarshalBlock()
  }

  // try to use reflection?
  // try to use a "node.Subgraph" or "node.BlockGraph" for limits?
  return nil, errors.New("not a BlockMarshaler")
}
