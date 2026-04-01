package basics

import "fmt"

// if/else - parentheses not required, braces are mandatory
func classify(n int) string {
	if n < 0 {
		return "negative"
	} else if n == 0 {
		return "zero"
	} else {
		return "positive"
	}
}

// if with init statement - scoped to the if/else block
func checkDiv(a, b int) {
	if result := a / b; result > 10 {
		fmt.Println("large:", result)
	} else {
		fmt.Println("small:", result)
	}
}

// switch - no fallthrough by default (unlike C++)
// no break needed; each case is implicitly broken
func dayType(day string) string {
	switch day {
	case "Saturday", "Sunday":
		return "weekend"
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		return "weekday"
	default:
		return "unknown"
	}
}

// switch with no expression - acts like an if/else chain
func signOf(n int) string {
	switch {
	case n < 0:
		return "negative"
	case n > 0:
		return "positive"
	default:
		return "zero"
	}
}

// fallthrough: explicitly fall through to the next case body
func switchFallthrough(n int) {
	switch n {
	case 1:
		fmt.Println("one")
		fallthrough
	case 2:
		fmt.Println("one or two")
	case 3:
		fmt.Println("three")
	}
}

// for loop - the ONLY loop in Go (no while or do-while keywords)
func sumUpTo(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}

// while-style for (just omit init and post)
func countDown(n int) {
	for n > 0 {
		fmt.Println(n)
		n--
	}
}

// infinite loop with break
func firstEvenAbove(threshold int) int {
	n := threshold + 1
	for {
		if n%2 == 0 {
			break
		}
		n++
	}
	return n
}

// range over slice: i = index, v = value (copy)
func printAll(items []string) {
	for i, v := range items {
		fmt.Println(i, v)
	}
}

// range over map - iteration order is NOT guaranteed
func printMap(m map[string]int) {
	for k, v := range m {
		fmt.Println(k, "->", v)
	}
}

// range over string iterates runes (Unicode code points), not bytes
func printRunes(s string) {
	for i, r := range s {
		fmt.Printf("index %d: %c\n", i, r)
	}
}

// continue skips to the next iteration
func printOdds(n int) {
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

/*
Key differences from C++:
- No parentheses around if/for/switch conditions
- Opening brace MUST be on the same line (the compiler inserts semicolons)
- for is the only loop keyword (while/do-while don't exist)
- switch cases don't fall through by default
- range gives you a copy of the value; use index to modify the slice
*/
