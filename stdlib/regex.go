package stdlib

import (
	"fmt"
	"regexp"
)

// regexp package: regular expressions (RE2 syntax)
// Note: Go uses RE2 - no backreferences, no lookaheads

// MustCompile: panics on invalid pattern (use at package level for constants)
var (
	emailRe  = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	digitRe  = regexp.MustCompile(`\d+`)
	wordsRe  = regexp.MustCompile(`\b\w+\b`)
)

// Compile: returns error on invalid pattern (use at runtime/dynamic patterns)
func compileDemo(pattern string) (*regexp.Regexp, error) {
	return regexp.Compile(pattern)
}

func matchDemo() {
	// MatchString: quick check without compiling
	ok, _ := regexp.MatchString(`^\d+$`, "12345")
	fmt.Println(ok) // true

	// compiled pattern methods
	fmt.Println(emailRe.MatchString("user@example.com")) // true
	fmt.Println(emailRe.MatchString("not-an-email"))     // false
}

func findDemo(s string) {
	// Find: first match
	first := digitRe.FindString("abc123def456")
	fmt.Println(first) // "123"

	// FindAll: all matches (n=-1 means all)
	all := digitRe.FindAllString("abc123def456ghi789", -1)
	fmt.Println(all) // [123 456 789]

	// FindStringIndex: returns [start, end] indices
	idx := digitRe.FindStringIndex("abc123")
	fmt.Println(idx) // [3 6]

	// FindAllStringIndex
	indices := digitRe.FindAllStringIndex("abc123def456", -1)
	fmt.Println(indices) // [[3 6] [9 12]]
}

// capturing groups
func captureDemo() {
	re := regexp.MustCompile(`(\w+)@(\w+)\.(\w+)`)
	s := "user@example.com"

	// FindStringSubmatch: full match + capture groups
	match := re.FindStringSubmatch(s)
	fmt.Println(match) // [user@example.com user example com]
	// match[0] = full match, match[1..] = capture groups

	// named captures
	namedRe := regexp.MustCompile(`(?P<user>\w+)@(?P<domain>\w+)\.(?P<tld>\w+)`)
	m := namedRe.FindStringSubmatch(s)
	for i, name := range namedRe.SubexpNames() {
		if i != 0 && name != "" {
			fmt.Printf("%s: %s\n", name, m[i])
		}
	}
}

func replaceDemo(s string) string {
	// ReplaceAllString: replace all matches
	clean := digitRe.ReplaceAllString(s, "#")
	fmt.Println(clean)

	// ReplaceAllStringFunc: replace with computed value
	upper := wordsRe.ReplaceAllStringFunc(s, func(w string) string {
		return "[" + w + "]"
	})
	fmt.Println(upper)

	return clean
}

func splitDemo(s string) {
	// Split: split by regex separator
	re := regexp.MustCompile(`\s*,\s*`)
	parts := re.Split("a, b,c,  d", -1)
	fmt.Println(parts) // [a b c d]
}

/*
RE2 cheat sheet:
  .         any char except newline
  \d        digit [0-9]
  \w        word char [0-9A-Za-z_]
  \s        whitespace
  ^         start of string (or line with (?m))
  $         end of string (or line with (?m))
  *         0 or more
  +         1 or more
  ?         0 or 1
  {n,m}     between n and m
  (...)     capture group
  (?P<name>...)  named capture group
  (?:...)   non-capturing group
  [abc]     character class
  [^abc]    negated class

Flags (prefix with (?flags)):
  i  case insensitive
  m  multiline (^ and $ match line boundaries)
  s  . matches \n too
*/
