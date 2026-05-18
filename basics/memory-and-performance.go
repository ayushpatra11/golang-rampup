package basics

// Notes on Go's memory model and performance considerations.
// Not runnable code - these are annotated examples and explanations.

/*
---- GARBAGE COLLECTOR ----
Go has a concurrent, tri-color mark-and-sweep GC.
- You don't manage memory manually
- GC runs concurrently with your program (low pause times)
- Tune with GOGC env var (default 100 = run GC when heap doubles)
- GOMEMLIMIT (Go 1.19+): cap total memory usage

Allocation tips to reduce GC pressure:
- Reuse objects with sync.Pool
- Prefer value types (structs) over pointers for small objects
- Pre-allocate slices with make([]T, 0, expectedLen)
- Pre-allocate maps with make(map[K]V, expectedSize)
- Use []byte + bytes.Buffer instead of string concatenation
- Avoid unnecessary heap escapes (see escape analysis below)
*/

/*
---- ESCAPE ANALYSIS ----
When a variable's address is taken and outlives the function,
the compiler moves it to the heap. Otherwise it stays on the stack.

- Stack allocation: fast, no GC overhead
- Heap allocation: GC must track it

Check with: go build -gcflags="-m" ./...

Causes of heap allocation:
- Returning a pointer to a local variable
- Storing a pointer in an interface{}
- Slices/maps (backing array is always heap)
- Closures capturing variables
- Variables too large for the stack

Often not worth optimizing unless profiling shows it's a bottleneck.
*/

/*
---- sync.Pool ----
Reuse temporary objects to reduce allocator pressure.
Objects may be GC'd at any time; don't store state that must persist.

var pool = sync.Pool{
    New: func() any { return make([]byte, 4096) },
}

func handler(w http.ResponseWriter, r *http.Request) {
    buf := pool.Get().([]byte)
    defer pool.Put(buf)
    // use buf...
}
*/

/*
---- PROFILING ----
Go has built-in profiling tools.

CPU profile:
  go test -cpuprofile=cpu.out ./...
  go tool pprof cpu.out

Memory profile:
  go test -memprofile=mem.out ./...
  go tool pprof mem.out

HTTP endpoint (add to your server):
  import _ "net/http/pprof"
  http.ListenAndServe(":6060", nil)
  Then: go tool pprof http://localhost:6060/debug/pprof/heap

Trace (goroutine scheduling):
  import "runtime/trace"
  trace.Start(f); defer trace.Stop()
  go tool trace trace.out

Benchmarks: go test -bench=. -benchmem -count=5 ./...
*/

/*
---- COMPILER DIRECTIVES ----

//go:noinline   prevents function inlining
//go:nosplit    prevents goroutine stack growth (low-level use only)
//go:generate  marks commands to run with `go generate`

Example:
//go:generate stringer -type=Direction
type Direction int
const (North Direction = iota; South; East; West)
// generates String() method for Direction via `go generate`
*/

/*
---- COMMON PERFORMANCE ANTI-PATTERNS ----

1. String concatenation in a loop:
   BAD:  for _, s := range strs { result += s }
   GOOD: strings.Join(strs, "") or strings.Builder

2. Append without pre-allocation:
   BAD:  var s []int; for i := 0; i < 1000; i++ { s = append(s, i) }
   GOOD: s := make([]int, 0, 1000)

3. Unnecessary interface boxing:
   BAD:  func sum(nums []interface{}) int64 { ... }
   GOOD: func sum(nums []int64) int64 { ... }  or generics

4. Copying large values in loops:
   BAD:  for _, item := range largeStructSlice { process(item) }
   GOOD: for i := range largeStructSlice { process(&largeStructSlice[i]) }

5. Not reusing HTTP clients:
   BAD:  client := &http.Client{}  (inside handler function)
   GOOD: var client = &http.Client{Timeout: 10*time.Second}  (package level)
*/
