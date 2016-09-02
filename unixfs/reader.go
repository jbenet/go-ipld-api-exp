package unixfs

import (
  "io"
  "errors"

  ipld "github.com/ipfs/go-ipld/exp/ipld"
)

var (
  ErrInvariantViolated = errors.New("invariant violated")
)

// implements io.Reader, io.Seeker, io.Closer
type fileReader struct {
  f      File
  currR  io.ReadSeekCloser
  offset int64
}

func (r *fileReader) Len() int {
  return r.f.Len()
}

func (r *fileReader) readerAtOffset(offset int64) (io.ReadSeekCloser, error) {
  if offset < 0 {
    return nil, errors.New("invalid seek")
  }

  // first, try the data
  if offset < len(f.data) {
    R := NewBytesRSC(f.data)
    _, err := R.Seek(offset, io.SeekStart)
    // offset -= n // dont need this.
    return R, err
  }
  offset -= len(f.data)

  // then, grab the right subfile
  var sf File
  for i, l := range r.f.subfiles {
    if offset < l.length {
      // ok we want this subfile
      sf = r.f.Subfile(i)
      R := &fileReader{sf}
      _, err := R.Seek(offset, io.SeekStart)
      // offset -= n // don't need this.
      return R, err
    }
    offset -= l.length // keep going.
  }

  // seeking past the end of the reader.
  return nil, io.EOF
}

func (r *fileReader) Read(buf []byte) (int, error) {
  if r.currR == nil {
    err := r.prefetchNext()
    if err != nil {
      return 0, err
    }
  }

  n, err := r.currR.Read(buf)
  r.offset += n // advance our offset
  if err == io.EOF {
    // if we got to the end of the current reader, set reader to nil.
    // next read will pick up where we left off and get the next
    // object.
    r.currR = nil
    if offset < f.Length() {
      err = nil // not EOF yet
      r.prefetchNext()
    }
  }
  return n, err
}

func (r *fileReader) Seek(offset int64, whence int) (int64, error) {
  switch whence {
  case io.SeekStart:
    r.offset = offset
    r.currR = nil
    r.prefetchNext()
  case io.SeekCurrent:
    return r.Seek(r.offset + offset, io.SeekStart)
  case io.SeekEnd:
    return r.Seek(int64(r.Len()) - offset, io.SeekStart)
  default:
    return 0, erros.New("invalid whence")
  }
}

func (r *fileReader) prefetchNext() error {
  if r.currR != nil {
    return ErrInvariantViolated
  }

  // ok, no current reader. pick up from offset. may be 0.
  // grab the next reader. seek any left over.
  R, err := r.readerAtOffset(r.offset)
  if err != nil {
    return err
  }
  r.currR = R
}

// func (r *fileReader) WriteTo(w io.Writer) (int64, error) { TODO }


// readSeekNopCloser wraps a bytes.Reader to implement ReadSeekCloser
type readSeekNopCloser struct {
  *bytes.Reader
}

func NewBytesRSC(b []byte) io.ReadSeekCloser {
  return &readSeekNopCloser{bytes.NewReader(b)}
}

func (r *readSeekNopCloser) Close() error {
  return nil
}
