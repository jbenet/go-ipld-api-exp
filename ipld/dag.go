package ipld

type Graph interface {
  Nodes() <-chan Node
}
