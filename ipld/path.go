package ipld

import (
  "strings"
)

// Path represents an ipld path.
//
// It may be an absolute path, for example:
//   /ipld/<cid>/a/b/c
// Or a partial path:
//   a/b/c
type Path string

func PathSplit(p Path) []string {
  s := strings.Split(p, "/")
  if s[0] == "" && len(s) > 1{
    s = s[1:]
  }
  return s
}
