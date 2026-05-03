package stdlib

import (
	"cmp"
	"fmt"
	"maps"
	"slices"
)

// slices package (Go 1.21): generic functions for slice operations
// replaces many sort.Slice + manual loops

func slicesPkgDemo() {
	nums := []int{3, 1, 4, 1, 5, 9, 2, 6}

	// sort
	slices.Sort(nums)
	fmt.Println(nums) // [1 1 2 3 4 5 6 9]

	// sort with custom comparator
	words := []string{"banana", "apple", "cherry"}
	slices.SortFunc(words, func(a, b string) int {
		return cmp.Compare(a, b) // cmp.Compare returns -1, 0, 1
	})
	fmt.Println(words)

	// search (binary search on sorted slice)
	idx, found := slices.BinarySearch(nums, 5)
	fmt.Println(idx, found) // 5 true

	// contains
	fmt.Println(slices.Contains(nums, 9)) // true

	// index: first occurrence
	fmt.Println(slices.Index(nums, 1)) // 0

	// min/max
	fmt.Println(slices.Min(nums)) // 1
	fmt.Println(slices.Max(nums)) // 9

	// reverse (in place)
	slices.Reverse(nums)
	fmt.Println(nums) // [9 6 5 4 3 2 1 1]

	// compact: remove consecutive duplicates (like uniq)
	slices.Sort(nums)
	unique := slices.Compact(nums)
	fmt.Println(unique) // [1 2 3 4 5 6 9]

	// equal
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	fmt.Println(slices.Equal(a, b)) // true

	// clip: reduce capacity to length (free unused backing array space)
	s := make([]int, 3, 100)
	clipped := slices.Clip(s)
	fmt.Println(len(clipped), cap(clipped)) // 3 3

	// grow: ensure capacity without changing length
	grown := slices.Grow(s, 50)
	fmt.Println(len(grown), cap(grown)) // 3 >= 53

	// delete: remove elements at index range
	nums = []int{1, 2, 3, 4, 5}
	nums = slices.Delete(nums, 1, 3) // remove indices 1,2
	fmt.Println(nums)                // [1 4 5]

	// insert
	nums = slices.Insert(nums, 1, 10, 20) // insert 10,20 at index 1
	fmt.Println(nums)                     // [1 10 20 4 5]

	// concat
	x := slices.Concat([]int{1, 2}, []int{3, 4}, []int{5})
	fmt.Println(x) // [1 2 3 4 5]
}

// maps package (Go 1.21): generic functions for map operations
func mapsPkgDemo() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	// keys and values (order not guaranteed)
	keys := slices.Sorted(maps.Keys(m))     // sorted for determinism
	vals := slices.Sorted(maps.Values(m))
	fmt.Println(keys) // [a b c]
	fmt.Println(vals) // [1 2 3]

	// clone: shallow copy
	m2 := maps.Clone(m)
	m2["d"] = 4
	fmt.Println(len(m), len(m2)) // 3 4

	// copy: copy key-value pairs from src into dst (overwrites existing keys)
	dst := map[string]int{"a": 99}
	maps.Copy(dst, m) // m overwrites dst
	fmt.Println(dst["a"]) // 1 (from m)

	// delete by condition
	maps.DeleteFunc(m2, func(k string, v int) bool {
		return v%2 == 0 // delete even values
	})
	fmt.Println(m2)

	// equal
	fmt.Println(maps.Equal(m, m2)) // false

	// collect from iter.Seq2 (Go 1.23+)
	// m3 := maps.Collect(iter.Map2(...))
}

/*
slices and maps packages avoid error-prone manual loop patterns.
cmp.Compare(a, b) is the standard generic comparator: returns -1/0/1.
cmp.Less(a, b) is equivalent to a < b for ordered types.

These packages require Go 1.21+.
*/
