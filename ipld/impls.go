package ipld

import (
  "errors"

  mcp "github.com/multiformats/go-multicodec-packed"
  block "github.com/ipfs/go-ipld/exp/block"
)

// CodecTable maps from a multicodec-packed code to a Codec
// to use in constructing Nodes.
type CodecTable map[mcp.Code]Codec

// Codec represents an ipld implementation codec.
// This is used to construct ipld nodes from Blocks.
type Codec interface {
  Decode(b block.Block) (n Node, err error)
  Encode(n Node) (b block.Block, err error)
}




func BlockWithNode(n Node) (block.Block, error) {
  if nb, ok := n.(block.BlockMarshaler); ok {
    return nb.MarshalBlock()
  }

  // otherwise, do it with "sort of reflection"
  // with a default serialization type: cbor
  return cbor.MarshalBlock(n)
}

func NodeWithBlock(b block.Block) n Node {
  var n Node

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

// Transformation converts nodes from a to b. This a generic
// transformation, that loses type-specific information. In general
// use it with type assertions:
//
//   b := getABlockFromWire()
//
//   var n ipld.Node
//   n := ipld.NodeWithBlock(&n)
//
//
//
//
type Transformation func(a Node) (b Node, err error)

func Transform(src, Tranformation) Node {

}
