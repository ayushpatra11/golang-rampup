package basics

import "fmt"

// ---- ARRAYS ----
// Fixed size, value type (assignment copies the whole array)

func arrayDemo() {
	var a [5]int               // zero-initialized: [0 0 0 0 0]
	b := [3]string{"x", "y", "z"}
	c := [...]int{1, 2, 3, 4}  // compiler infers length -> [4]int

	fmt.Println(a, b, c)
	fmt.Println(len(b)) // 3

	// arrays are comparable if element type is comparable
	d := [3]int{1, 2, 3}
	e := [3]int{1, 2, 3}
	fmt.Println(d == e) // true
}

// ---- SLICES ----
// Dynamic, reference type - backed by an array
// Slice header = (pointer to array) + length + capacity

func sliceDemo() {
	s := []int{10, 20, 30, 40, 50}

	// slicing: [low:high] -> low inclusive, high exclusive
	fmt.Println(s[1:3]) // [20 30]
	fmt.Println(s[:2])  // [10 20]
	fmt.Println(s[3:])  // [40 50]

	fmt.Println(len(s), cap(s)) // 5 5

	// nil slice: zero value, len=0, cap=0
	var ns []int
	fmt.Println(ns == nil, len(ns)) // true 0
}

// make: create slice with given length and capacity
func makeSlice() {
	s := make([]int, 3, 5) // len=3, cap=5, elements zero-initialized
	fmt.Println(s)         // [0 0 0]
	fmt.Println(len(s), cap(s)) // 3 5
}

// append grows the slice; may allocate a new backing array
// always use the returned slice - the original may be stale
func appendDemo() {
	var s []int
	for i := 0; i < 5; i++ {
		s = append(s, i*i)
	}
	fmt.Println(s) // [0 1 4 9 16]

	// append one slice to another using ... to unpack
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	a = append(a, b...)
	fmt.Println(a) // [1 2 3 4 5 6]
}

// copy: copies min(len(dst), len(src)) elements, returns count
func copyDemo() {
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	n := copy(dst, src)
	fmt.Println(dst, n) // [1 2 3] 3

	// copy between overlapping regions works correctly
}

// delete element at index i (order not preserved - fast)
func deleteUnordered(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// delete element at index i (order preserved)
func deleteOrdered(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}

// 2D slice (slice of slices)
func matrix(rows, cols int) [][]int {
	m := make([][]int, rows)
	for i := range m {
		m[i] = make([]int, cols)
	}
	return m
}

func reverseArray(nums []int, left int, right int) {
	// There is no default function for reverse, 
	// have to define it manually.
	// Recently used it in LC 189 to rotate an array.
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

/*
Important: slices share the underlying array.
Mutations through one slice are visible through another if they overlap.

    s := []int{1, 2, 3, 4, 5}
    a := s[:3]  // [1 2 3]
    b := s[2:]  // [3 4 5]
    a[2] = 99   // modifies index 2 of the backing array
    fmt.Println(b[0]) // 99  <- shared!

Once append exceeds capacity, a NEW array is allocated and slices diverge.
Always reassign: s = append(s, x)
*/
