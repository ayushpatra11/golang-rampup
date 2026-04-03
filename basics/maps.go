package basics

import "fmt"

// map: built-in hash map (like unordered_map in C++)
// zero value of a map is nil - nil map reads return zero value, writes panic

func mapDemo() {
	// map literal
	m := map[string]int{
		"alice": 30,
		"bob":   25,
	}

	// read - returns zero value for missing keys (no error)
	age := m["alice"] // 30
	fmt.Println(age)

	// ok idiom: distinguish "missing key" from "stored zero value"
	val, ok := m["charlie"]
	if !ok {
		fmt.Println("not found, zero value:", val) // 0
	}

	// write / update
	m["charlie"] = 40
	m["alice"] = 31

	// delete (no-op if key doesn't exist)
	delete(m, "bob")

	fmt.Println(len(m)) // 2

	// iteration order is NOT guaranteed - randomized by design
	for k, v := range m {
		fmt.Println(k, v)
	}
}

// make: create empty map
func makeMap() map[string][]string {
	return make(map[string][]string)
}

// common pattern: group items by key
func groupByLength(words []string) map[int][]string {
	groups := make(map[int][]string)
	for _, w := range words {
		n := len(w)
		groups[n] = append(groups[n], w)
	}
	return groups
}

// counting occurrences
func wordCount(words []string) map[string]int {
	freq := make(map[string]int)
	for _, w := range words {
		freq[w]++ // zero value of int is 0, so this always works
	}
	return freq
}

// set: there's no built-in set in Go; use map[T]struct{}
// struct{} takes 0 bytes (empty struct)
func toSet(items []string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, item := range items {
		set[item] = struct{}{}
	}
	return set
}

func inSet(set map[string]struct{}, key string) bool {
	_, ok := set[key]
	return ok
}

// maps are NOT safe for concurrent reads+writes
// use sync.Map or protect with sync.RWMutex

/*
Compared to C++ std::unordered_map:
- No .find() - use the two-value assignment: v, ok := m[k]
- No iterator invalidation concept (GC handles memory)
- Cannot take address of map value: &m["key"] is invalid
- Map values are not addressable; copy out to modify:
    type Counter struct{ n int }
    m := map[string]Counter{"a": {}}
    c := m["a"]; c.n++; m["a"] = c   <- must copy back
*/
