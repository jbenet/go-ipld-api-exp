
type Block interface {
  CID() cid.CID
  // CID().Multihash()
  // CID().Multicodec()

  Data() []byte
}

type BlockMarshaler interface {
  MarshalBlock() (Block, error)
}

type BlockUnmarshaler interface {
  UnmarshalBlock(Block) error
}
