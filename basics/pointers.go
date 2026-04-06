package basics

import "fmt"

// Go has pointers but NO pointer arithmetic (safer than C/C++)

func pointerBasics() {
	x := 42
	p := &x        // p is *int; & takes the address
	fmt.Println(*p) // 42; * dereferences

	*p = 100
	fmt.Println(x) // 100; x was modified through the pointer
}

// new: allocates zeroed memory, returns pointer
func newDemo() {
	p := new(int) // *int, value is 0
	*p = 5
	fmt.Println(*p) // 5
}

// pointers allow functions to modify their arguments
func increment(n *int) {
	*n++
}

func incrementDemo() {
	x := 10
	increment(&x)
	fmt.Println(x) // 11
}

// pointer to struct: fields accessed with . (no -> like in C++)
type Point struct{ X, Y int }

func moveRight(p *Point, dx int) {
	p.X += dx // Go auto-dereferences: equivalent to (*p).X += dx
}

func pointerStructDemo() {
	pt := Point{1, 2}
	moveRight(&pt, 5)
	fmt.Println(pt) // {6 2}
}

// returning a pointer to a local variable is SAFE in Go
// the compiler promotes the variable to the heap ("escape analysis")
func newPoint(x, y int) *Point {
	return &Point{x, y} // safe - Go manages memory
}

// nil pointer: zero value of any pointer type
func nilPointerDemo() {
	var p *int
	fmt.Println(p == nil) // true
	// *p = 5 would panic: nil pointer dereference
}

// pointer vs value receivers (see functions/ for full methods coverage)
// pointer receiver: can mutate, avoids copy for large structs
// value receiver: cannot mutate the original, gets a copy

/*
Key differences from C++:
- No pointer arithmetic (p++ is a compile error)
- No manual malloc/free - GC handles memory
- Local variable addresses are safe to return (escape analysis)
- Dereferencing a nil pointer panics (rather than undefined behavior)
- . operator works for both pointer and value (no ->)
*/
