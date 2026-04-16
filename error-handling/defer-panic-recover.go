package errorhandling

import (
	"fmt"
	"os"
)

// ---- DEFER ----
// defer schedules a function call to run when the surrounding function returns
// (whether by normal return, panic, or os.Exit)
// deferred calls run in LIFO order (stack)

func deferOrder() {
	fmt.Println("start")
	defer fmt.Println("first")  // runs 3rd
	defer fmt.Println("second") // runs 2nd
	defer fmt.Println("third")  // runs 1st
	fmt.Println("end")
	// output: start, end, third, second, first
}

// primary use: cleanup that must happen regardless of return path
func writeToFile(path, content string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close() // guaranteed even if we return early below

	_, err = fmt.Fprintln(f, content)
	return err
}

// defer + named return: deferred function can modify the named return value
func safeRead(path string) (content string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer func() {
		if cerr := f.Close(); cerr != nil && err == nil {
			err = cerr // set named return error on close failure
		}
	}()
	buf := make([]byte, 1024)
	n, err := f.Read(buf)
	content = string(buf[:n])
	return
}

// defer evaluates arguments immediately (but runs the call later)
func deferArgEval() {
	x := 10
	defer fmt.Println("deferred x:", x) // captures x=10 now
	x = 20
	fmt.Println("current x:", x) // 20
	// deferred prints 10
}

// ---- PANIC ----
// panic: stop normal execution and unwind the stack running deferred functions
// use only for programming errors / invariant violations, NOT for expected errors

func mustPositive(n int) int {
	if n <= 0 {
		panic(fmt.Sprintf("expected positive, got %d", n))
	}
	return n
}

// ---- RECOVER ----
// recover: catches a panic; must be called inside a deferred function
// returns the value passed to panic, or nil if no panic in progress

func safeDiv(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from panic: %v", r)
		}
	}()
	result = a / b // panics if b == 0
	return
}

func recoverDemo() {
	result, err := safeDiv(10, 0)
	fmt.Println(result, err) // 0, recovered from panic: runtime error: integer divide by zero
}

// package-level recovery: wrap external-facing code in recover to convert panics to errors
func safeDo(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	fn()
	return
}

/*
When to use each:
- defer:   cleanup (Close, Unlock, wg.Done, span.End)
- panic:   programming bugs, invariant violations, unrecoverable state
- recover: at package/service boundaries to convert panics into errors
           (e.g., HTTP handler middleware, test harnesses)

Don't use panic/recover as a general exception mechanism.
Prefer explicit error returns for expected failure cases.
*/
