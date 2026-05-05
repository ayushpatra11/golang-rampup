package functions

import "fmt"

// ---- init functions ----
// init() runs automatically before main(), after package-level var declarations
// - takes no arguments, returns nothing
// - multiple init() functions allowed per file (and per package)
// - init() in imported packages runs before the importing package's init()
// - execution order within a package: declaration order in files, file order is alphabetical

var config = loadConfig() // package-level vars initialized first

func loadConfig() string {
	return "config-loaded"
}

func init() {
	// runs after all var declarations in this package
	fmt.Println("init: package initialized with config:", config)
}

// multiple init functions in one file (unusual but valid)
func init() {
	fmt.Println("init: second init in the same file")
}

// ---- main function ----
// Entry point for executable programs (package main only)
// No parameters, no return value
// os.Args[0] = program name, os.Args[1:] = arguments
// os.Exit(code) for non-zero exit; defers do NOT run after os.Exit

/*
func main() {
    args := os.Args[1:]
    if len(args) == 0 {
        fmt.Fprintln(os.Stderr, "usage: program <name>")
        os.Exit(1)
    }
    fmt.Println("Hello,", args[0])
}
*/

// ---- blank identifier ----
// _ discards values you don't need
func blankIdentifier() {
	// discard index in range
	for _, v := range []int{1, 2, 3} {
		fmt.Println(v)
	}

	// discard error (only do this when you're sure it can't fail)
	n, _ := fmt.Println("hello") // discard byte count
	_ = n

	// import for side effects only (runs init but nothing else is used)
	// import _ "net/http/pprof"
}

// ---- multiple return values ----
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// named return values: can improve readability for complex functions
func minMax(nums []int) (min, max int) {
	if len(nums) == 0 {
		return // returns zero values for min and max
	}
	min, max = nums[0], nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return // naked return - returns named values
}

/*
Package initialization order:
1. Imported packages are initialized first (recursively)
2. Package-level variables are initialized in declaration order
   (with dependency resolution: if var A depends on var B, B is initialized first)
3. init() functions run in source file order (alphabetical by filename)
4. main() runs last in the main package

Use init() sparingly - it makes the startup sequence harder to follow.
Prefer explicit initialization over implicit init() side effects.
*/
