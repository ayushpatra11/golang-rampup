package basics

// A collection of Go idioms and patterns that come up frequently.
// Based on Effective Go, Google's Go style guide, and community conventions.

/*
---- NAMING ----
- Short, clear names for local variables (i, v, n, buf, err)
- Exported names should be descriptive: UserService not USrvc
- Acronyms stay consistent case: URL, HTTP, ID (not Url, Http, Id)
- Package name = last element of import path; don't repeat in exports:
  bufio.Reader not bufio.BufReader
- Getter: Name() not GetName(); Setter: SetName() is fine
- Interface ending in -er: Reader, Writer, Stringer, Handler
*/

/*
---- COMMENTS ----
- Doc comment on every exported symbol: starts with the symbol name
  // Server handles incoming HTTP connections.
  type Server struct { ... }

- go doc reads these; godoc renders them as HTML
- Use // not /* */ for doc comments
*/

/*
---- EARLY RETURNS ----
Go prefers "happy path not indented" - return early on errors.

// preferred
func process(id int) (string, error) {
    if id <= 0 {
        return "", errors.New("invalid id")
    }
    // ... happy path continues at top level
    return result, nil
}

// avoid deeply nested else chains
*/

/*
---- GUARD CLAUSES ----
Validate all preconditions first, then do the work.

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
*/

/*
---- ACCEPT INTERFACES, RETURN STRUCTS ----
Parameters: accept interfaces (more flexible, easier to test)
Return types: return concrete types (gives callers access to full API)

// DO
func Save(w io.Writer, data []byte) error { ... }

// DON'T (unnecessarily restrictive return)
func NewBuffer() io.Writer { return &bytes.Buffer{} }
// Caller can't call Bytes() or Reset() which are on *bytes.Buffer
*/

/*
---- ZERO VALUES ----
Design types so the zero value is useful.

sync.Mutex   zero value is an unlocked mutex - no constructor needed
bytes.Buffer zero value is an empty, ready-to-use buffer
sync.Once    zero value is ready to use

type Config struct {
    Timeout time.Duration // zero = no timeout (useful!)
    MaxSize int           // zero = unlimited (useful if documented)
}
*/

/*
---- GOROUTINE NAMING / TRACING ----
Goroutines are anonymous but runtime/pprof labels can add context.

func labeledGoroutine(ctx context.Context) {
    go func() {
        labels := pprof.Labels("handler", "user-request")
        pprof.Do(ctx, labels, func(ctx context.Context) {
            // work here shows as "user-request" in profiles
        })
    }()
}
*/

/*
---- PACKAGE DESIGN ----
- One package per directory (except _test packages)
- Package provides a single, cohesive abstraction
- Avoid circular imports (restructure if needed)
- Internal packages (internal/) restrict importers to parent subtree
- cmd/ for main packages (executables)
- pkg/ is optional and increasingly avoided

Typical layout:
  myapp/
    cmd/
      myapp/main.go
    internal/
      db/
      config/
    pkg/
      api/
    go.mod
*/

/*
---- COMMON GOTCHAS ----
1. Loop variable capture in closures:
   for _, v := range items { go func() { use(v) }() }  <- all goroutines use last v
   Fix: go func(v T) { use(v) }(v)  or  v := v  inside loop

2. Nil interface vs nil pointer:
   var p *MyType = nil
   var i MyInterface = p
   i == nil  is FALSE (interface holds a non-nil type descriptor)

3. Defer in a loop doesn't run until function returns:
   for _, f := range files { defer f.Close() }  <- files stay open until function exits
   Fix: wrap body in a closure: func() { defer f.Close(); process(f) }()

4. Slice sharing surprises:
   a := s[:3]; b := s[1:4]  <- share backing array; a[2] == b[1]

5. Map iteration is random:
   Never rely on map iteration order.

6. Goroutine leaks from forgotten channels:
   Always ensure every goroutine has a way to exit.
*/
