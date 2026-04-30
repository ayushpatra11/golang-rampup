package stdlib

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// ---- strings package ----

func stringsDemo() {
	s := "Hello, Go World!"

	// searching
	fmt.Println(strings.Contains(s, "Go"))       // true
	fmt.Println(strings.HasPrefix(s, "Hello"))   // true
	fmt.Println(strings.HasSuffix(s, "!"))       // true
	fmt.Println(strings.Count(s, "o"))           // 2
	fmt.Println(strings.Index(s, "Go"))          // 7
	fmt.Println(strings.LastIndex(s, "o"))       // 14

	// transforming
	fmt.Println(strings.ToUpper(s))              // HELLO, GO WORLD!
	fmt.Println(strings.ToLower(s))              // hello, go world!
	fmt.Println(strings.Title("hello world"))    // Hello World (deprecated; use golang.org/x/text)
	fmt.Println(strings.TrimSpace("  hi  "))     // "hi"
	fmt.Println(strings.Trim("!!hi!!", "!"))     // "hi"
	fmt.Println(strings.TrimLeft("xxhello", "x"))  // "hello"
	fmt.Println(strings.TrimPrefix("hello.go", "hello.")) // "go"

	// splitting / joining
	fmt.Println(strings.Split("a,b,c", ","))          // [a b c]
	fmt.Println(strings.Fields("  foo  bar  baz  "))  // [foo bar baz] (whitespace split)
	fmt.Println(strings.Join([]string{"a", "b"}, "-")) // a-b
	fmt.Println(strings.Repeat("ab", 3))               // ababab

	// replacing
	fmt.Println(strings.Replace(s, "o", "0", -1))  // replace all; use 1 for first only
	fmt.Println(strings.ReplaceAll(s, "o", "0"))    // same as Replace with -1

	// checking
	fmt.Println(strings.ContainsAny(s, "aeiou"))    // true (any vowel)
	fmt.Println(strings.ContainsRune(s, 'G'))       // true

	// Map: apply function to each rune
	rot13 := strings.Map(func(r rune) rune {
		switch {
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		}
		return r
	}, "Hello")
	fmt.Println(rot13) // Uryyb

	// strings.Builder: efficient string concatenation
	var b strings.Builder
	for i := 0; i < 5; i++ {
		fmt.Fprintf(&b, "%d", i)
	}
	fmt.Println(b.String()) // 01234

	// strings.Reader: io.Reader over a string
	r := strings.NewReader("hello")
	buf := make([]byte, 3)
	r.Read(buf)
	fmt.Println(string(buf)) // hel
}

// ---- unicode package ----

func unicodeDemo() {
	for _, r := range "Hello, 世界 42!" {
		switch {
		case unicode.IsLetter(r):
			fmt.Printf("%c is a letter\n", r)
		case unicode.IsDigit(r):
			fmt.Printf("%c is a digit\n", r)
		case unicode.IsSpace(r):
			fmt.Printf("space\n")
		}
	}
}

// ---- strconv: string <-> primitive type conversions ----

func strconvDemo() {
	// int <-> string
	s := strconv.Itoa(42)      // int to string: "42"
	n, err := strconv.Atoi("123") // string to int
	if err == nil {
		fmt.Println(n) // 123
	}

	// invalid conversion returns error
	_, err = strconv.Atoi("abc")
	fmt.Println(err) // strconv.Atoi: parsing "abc": invalid syntax

	// float <-> string
	fs := strconv.FormatFloat(3.14159, 'f', 2, 64)  // "3.14"
	f, _ := strconv.ParseFloat("2.718", 64)
	fmt.Println(fs, f)

	// bool <-> string
	strconv.FormatBool(true)    // "true"
	strconv.ParseBool("false")  // false, nil

	// generic Format/Parse with base
	strconv.FormatInt(255, 16)  // "ff" (hex)
	strconv.FormatInt(255, 2)   // "11111111" (binary)
	strconv.ParseInt("ff", 16, 64) // 255, nil
}

// ---- sort package ----

func sortDemo() {
	// built-in sort helpers
	nums := []int{5, 2, 8, 1, 9, 3}
	sort.Ints(nums)
	fmt.Println(nums) // [1 2 3 5 8 9]

	words := []string{"banana", "apple", "cherry"}
	sort.Strings(words)
	fmt.Println(words) // [apple banana cherry]

	// custom sort with sort.Slice
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{{"Bob", 30}, {"Alice", 25}, {"Charlie", 35}}
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people) // [{Alice 25} {Bob 30} {Charlie 35}]

	// sort.SliceStable preserves original order of equal elements
	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})

	// binary search (slice must be sorted)
	idx, found := sort.Find(len(nums), func(i int) int {
		return nums[i] - 8 // target - element
	})
	fmt.Println(idx, found) // 5 true

	fmt.Println(sort.IntsAreSorted(nums)) // true
}

// ---- math package ----

func mathDemo() {
	fmt.Println(math.Abs(-3.5))           // 3.5
	fmt.Println(math.Sqrt(16))            // 4
	fmt.Println(math.Pow(2, 10))          // 1024
	fmt.Println(math.Floor(3.7))          // 3
	fmt.Println(math.Ceil(3.2))           // 4
	fmt.Println(math.Round(3.5))          // 4
	fmt.Println(math.Log(math.E))         // 1
	fmt.Println(math.Log2(8))             // 3
	fmt.Println(math.Log10(100))          // 2
	fmt.Println(math.MaxInt64)            // max int64 value
	fmt.Println(math.IsInf(1/0.0, 1))    // true (positive infinity)
	fmt.Println(math.IsNaN(math.NaN()))   // true
}
