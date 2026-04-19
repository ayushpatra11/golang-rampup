package stdlib

// test files MUST end in _test.go
// build tags can restrict which tests run
// run tests: go test ./...
// run verbose: go test -v ./...
// run specific: go test -run TestFoo ./...
// with race detector: go test -race ./...
// coverage: go test -cover ./...
// coverage HTML: go test -coverprofile=cov.out ./... && go tool cover -html=cov.out

// ---- Basic test ----
// Every test function must be TestXxx(t *testing.T)

/*
func TestAdd(t *testing.T) {
    got := Add(2, 3)
    want := 5
    if got != want {
        t.Errorf("Add(2,3) = %d; want %d", got, want)
    }
}
*/

// t.Error / t.Errorf:  mark failed, continue running
// t.Fatal / t.Fatalf:  mark failed, stop this test immediately
// t.Log / t.Logf:      only printed with -v or on failure
// t.Skip / t.Skipf:    skip test (e.g. slow or integration)

// ---- Table-driven tests ----
// Idiomatic Go: one function, many sub-cases.

/*
func TestAddTable(t *testing.T) {
    cases := []struct {
        name    string
        a, b    int
        want    int
    }{
        {"positive", 2, 3, 5},
        {"negative", -1, -2, -3},
        {"zero", 0, 0, 0},
        {"mixed", 10, -3, 7},
    }
    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            got := Add(tc.a, tc.b)
            if got != tc.want {
                t.Errorf("Add(%d,%d) = %d; want %d", tc.a, tc.b, got, tc.want)
            }
        })
    }
}
*/

// t.Run creates a sub-test that can be run individually:
//   go test -run TestAddTable/positive

// ---- Benchmarks ----
// BenchmarkXxx(b *testing.B)
// run: go test -bench=. -benchmem ./...

/*
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ { // b.N is tuned by the framework
        Add(2, 3)
    }
}
*/

// b.ResetTimer() if setup before the loop is expensive
// b.ReportAllocs() or use -benchmem for allocation stats

// ---- Test helpers ----
/*
func assertEqual(t *testing.T, got, want any) {
    t.Helper() // marks this function as a helper - error line shows the caller
    if got != want {
        t.Errorf("got %v; want %v", got, want)
    }
}
*/

// t.Helper() ensures error lines point to the caller, not the helper.

// ---- TestMain: global setup/teardown ----
/*
func TestMain(m *testing.M) {
    setup()
    code := m.Run()
    teardown()
    os.Exit(code)
}
*/

// ---- Fuzz testing (Go 1.18+) ----
/*
func FuzzReverse(f *testing.F) {
    f.Add("hello") // seed corpus
    f.Fuzz(func(t *testing.T, s string) {
        rev := Reverse(s)
        if Reverse(rev) != s {
            t.Errorf("double reverse of %q != original", s)
        }
    })
}
// run: go test -fuzz=FuzzReverse ./...
*/

// ---- Testable examples ----
// Example functions serve as documentation AND are compiled/run as tests.
/*
func ExampleAdd() {
    fmt.Println(Add(2, 3))
    // Output:
    // 5
}
*/
// If Output comment doesn't match, the test fails.

/*
Third-party testing libraries (commonly used):
- github.com/stretchr/testify/assert  (cleaner assertions)
- github.com/stretchr/testify/mock    (mocking)
- github.com/google/go-cmp/cmp       (deep equality, diffs)

Recommended practice:
- Use table-driven tests for anything with multiple cases
- Use t.Helper() in any test helper function
- Prefer small, focused tests over large integration tests
- Run with -race to catch concurrency bugs
*/
