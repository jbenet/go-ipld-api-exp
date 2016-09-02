func newFile(n dag.Node) (*file, error) {
  f := &file{n}

  // des  erialize method 1
  f2 := &fileParse{}
  err := n.Unmarshal(f2)

  // deserialize method 2
  var links []*fileLink
  err := n.Get("subfiles", &links)
  if err != nil {
    return nil, dag.ErrParse
  }

  // deserialize method 3
  var links interface{}
  err := n.Get("subfiles", &links)
  if err != nil {
    return nil, dag.ErrParse
  }
  linksMapArr, ok := link.([]interface{})
  if !ok {
    return nil, dag.ErrParse
  }
  for i := 0; i < len(linksMapArr) i++ {
    l, err := n.Get(fmt.Sprint("subfiles/", i))
    if err != nil {
      return nil, dag.ErrParse
    }

    m, _ := linksMapArr[i].(map[string]interface{})
    l := &fileLink{
      l:      dag.LinkFromRaw(m["link"]),
      length: int(m["len"].(float64)),
    }

    links[i] = fileLink{l.link, l.len}
  }


  // deserialize method 4
  var links Node
  err := n.Get("subfiles", &links)
  if err != nil {
    return nil, dag.ErrParse
  }

  linksArr := links.(NodeArray)
  for i := 0; i < linksArr.Length(); i++ {
    l, err := linksArr.Get(PathIndex(i))
  }


  f.links = links
  f.subfiles = make([]*file, len(links))
  return f, nil
}


`
{
  data: "foobarbaz",
  subfiles: [
    {
      len: 1000,
      link: {/: /ipld/A...},
    },
    {
      len: 1000,
      link: {/: /ipld/B...},
    }
  ]
}
`

`
---
subfiles:
- len: 1000
  link: #{/ipld/A...}
- len: 1000
  link: #{/ipld/B...}
`

`
[
  {
    len: 1000,
    link: {/: /ipld/A...},
  },
  {
    len: 1000,
    link: {/: /ipld/B...},
  }
]
`

// directories

`
{
  a: {
    size: 1000,
    link: {/: /ipld/A...},
  },
  b: {
    size: 1000,
    link: {/: /ipld/B...}
  },
}
`

