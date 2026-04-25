package basics

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// io.Reader and io.Writer are the most important interfaces in Go's stdlib
// Most I/O functions accept these rather than concrete types

// io.Reader: anything that can be read from
// type Reader interface { Read(p []byte) (n int, err error) }
//   Read fills p, returns bytes read and nil (or io.EOF when done, or error)

// io.Writer: anything that can be written to
// type Writer interface { Write(p []byte) (n int, err error) }

// ---- Composing readers/writers ----

func copyDemo() {
	src := strings.NewReader("hello, world")
	var dst bytes.Buffer
	n, err := io.Copy(&dst, src)
	fmt.Println(n, err, dst.String()) // 12 <nil> hello, world
}

// io.TeeReader: reads from r, writes a copy to w
func teeDemo(r io.Reader) {
	var buf bytes.Buffer
	tee := io.TeeReader(r, &buf)

	data, _ := io.ReadAll(tee)
	fmt.Println("read:", string(data))
	fmt.Println("copy:", buf.String()) // same content
}

// io.MultiReader: reads from multiple readers sequentially
func multiReaderDemo() {
	r := io.MultiReader(
		strings.NewReader("hello "),
		strings.NewReader("world"),
	)
	data, _ := io.ReadAll(r)
	fmt.Println(string(data)) // hello world
}

// io.MultiWriter: writes to multiple writers simultaneously
func multiWriterDemo() {
	var a, b bytes.Buffer
	w := io.MultiWriter(&a, &b)
	fmt.Fprint(w, "broadcast")
	fmt.Println(a.String(), b.String()) // broadcast broadcast
}

// io.LimitReader: wraps a reader to read at most n bytes
func limitDemo(r io.Reader) {
	limited := io.LimitReader(r, 5)
	data, _ := io.ReadAll(limited)
	fmt.Println(string(data)) // first 5 bytes
}

// io.Pipe: synchronous in-memory pipe connecting a writer to a reader
func pipeDemo() {
	pr, pw := io.Pipe()

	go func() {
		fmt.Fprint(pw, "hello from pipe")
		pw.Close()
	}()

	data, _ := io.ReadAll(pr)
	fmt.Println(string(data)) // hello from pipe
}

// bytes.Buffer: in-memory read/write buffer implementing Reader and Writer
func bufferDemo() {
	var buf bytes.Buffer

	fmt.Fprint(&buf, "hello")
	buf.WriteString(", world")
	buf.WriteByte('!')

	fmt.Println(buf.String())    // hello, world!
	fmt.Println(buf.Len())       // remaining bytes
	fmt.Println(buf.Bytes())     // []byte
	buf.Reset()                  // clear without releasing memory
}

// strings.Builder: write-only, efficient string construction
func builderDemo() {
	var sb strings.Builder
	for i := 0; i < 5; i++ {
		fmt.Fprintf(&sb, "item%d ", i)
	}
	result := strings.TrimSpace(sb.String())
	fmt.Println(result)
}

/*
Key io interfaces:
  io.Reader        Read(p []byte) (n int, err error)
  io.Writer        Write(p []byte) (n int, err error)
  io.Closer        Close() error
  io.Seeker        Seek(offset int64, whence int) (int64, error)
  io.ReadWriter    Reader + Writer
  io.ReadCloser    Reader + Closer
  io.WriteCloser   Writer + Closer
  io.ReadWriteCloser Reader + Writer + Closer

io.Discard: writer that discards all writes (like /dev/null)
io.NopCloser: wraps a Reader to add a no-op Close method
*/
