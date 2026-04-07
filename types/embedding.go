package types

import "fmt"

// embedding: include one type inside another without naming it
// embedded fields have their methods and fields promoted to the outer type

// ---- struct embedding ----

type Logger struct {
	prefix string
}

func (l Logger) Log(msg string) {
	fmt.Printf("[%s] %s\n", l.prefix, msg)
}

type Server struct {
	Logger        // embedded - not a named field
	host   string
	port   int
}

func embeddedStructDemo() {
	s := Server{
		Logger: Logger{prefix: "SERVER"},
		host:   "localhost",
		port:   8080,
	}

	s.Log("started") // promoted: s.Logger.Log("started")
	s.Logger.Log("explicit access also works")
	fmt.Println(s.prefix) // Logger.prefix is promoted too
}

// outer type can override a promoted method
type VerboseServer struct {
	Logger
	name string
}

func (v VerboseServer) Log(msg string) {
	// override the embedded method
	fmt.Printf("[VERBOSE][%s] %s\n", v.name, msg)
	v.Logger.Log(msg) // still callable via explicit access
}

// ---- interface embedding ----

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

// ReadWriter embeds both interfaces
type ReadWriter interface {
	Reader
	Writer
}

// a type satisfying ReadWriter must implement both Read and Write
type Buffer struct {
	data []byte
}

func (b *Buffer) Read(p []byte) (int, error) {
	n := copy(p, b.data)
	b.data = b.data[n:]
	return n, nil
}

func (b *Buffer) Write(p []byte) (int, error) {
	b.data = append(b.data, p...)
	return len(p), nil
}

// Buffer satisfies ReadWriter
func interfaceEmbedDemo() {
	var rw ReadWriter = &Buffer{}
	rw.Write([]byte("hello"))

	buf := make([]byte, 5)
	rw.Read(buf)
	fmt.Println(string(buf)) // hello
}

// multiple embedding - diamond inheritance is handled by name collision
// if two embedded types have the same method, it's ambiguous -> compile error
// unless the outer type overrides it

type A struct{}
type B struct{}

func (A) Hello() string { return "A" }
func (B) Hello() string { return "B" }

type C struct {
	A
	B
}

func (c C) Hello() string {
	return c.A.Hello() + c.B.Hello() // must disambiguate
}

/*
Embedding vs inheritance:
- Go embedding is NOT inheritance - no polymorphism via the outer type
- The embedded type doesn't know it's embedded
- Methods promoted, but the outer type is NOT a subtype of the embedded type
- c.Log() works, but func f(l Logger) { f(s) } does NOT work (s is not a Logger)
- For polymorphism, use interfaces
*/
