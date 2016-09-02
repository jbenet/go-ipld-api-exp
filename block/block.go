
type Block interface {
  CID() cid.CID
  // CID().Multihash()
  // CID().Multicodec()

  Data() []byte
}
