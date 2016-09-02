package mdagv1

import (
  ipld "github.com/ipfs/go-ipld/exp/ipld"
  mdag "github.com/ipfs/go-ipfs/merkledag"
)

type mdagIPLD struct {
  dn mdag.Node
}


func (m *mdagIPLD) Links() []string {
  s := make([]string, len(m.dn.Links))
  for i, l := range m.dn.Links {
    s[i] = l.Name
  }
  return s
}

func (m *mdagIPLD) GetNode(p Path) (Node, error) {
  s := PathSplit(p)
  c := s[0]

  l, err := m.dn.GetNodeLink(c)
  if err != nil {
    return nil, err
  }

  return mdagIPLD.graph.Resolve(l.Hash)
}

func (m *mdagIPLD) UnmarshalTo(v interface{}) error {
  // we need a way to set keys into v, depending on the struct
  // that it is. we probably need to use reflection here.
  // because we should be able to set on any of:
  //
  // v == interface{}
  // v == map[string]interface{}
  // v == struct{
  //   Data []byte
  //   Links map[string]struct{
  //     Name string
  //     Hash []byte
  //     Size int
  //   }
  // }
}
