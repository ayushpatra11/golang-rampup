package types

import "fmt"

// iota: auto-incrementing integer constant generator
// resets to 0 at each const block

type Direction int

const (
	North Direction = iota // 0
	East                   // 1
	South                  // 2
	West                   // 3
)

// implement fmt.Stringer so Direction prints as a name, not a number
func (d Direction) String() string {
	switch d {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	default:
		return fmt.Sprintf("Direction(%d)", int(d))
	}
}

// iota with expressions
type ByteSize float64

const (
	_           = iota // discard first value (0)
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
)

// iota with bit flags: each constant is a power of 2
type Permission uint

const (
	Read    Permission = 1 << iota // 1
	Write                          // 2
	Execute                        // 4
)

func (p Permission) String() string {
	s := ""
	if p&Read != 0 {
		s += "r"
	} else {
		s += "-"
	}
	if p&Write != 0 {
		s += "w"
	} else {
		s += "-"
	}
	if p&Execute != 0 {
		s += "x"
	} else {
		s += "-"
	}
	return s
}

func permissionDemo() {
	p := Read | Write
	fmt.Println(p)           // rw-
	fmt.Println(p&Execute != 0) // false
	p |= Execute
	fmt.Println(p)           // rwx
}

// sentinel zero value: use iota starting at 1 so the zero value is "unknown"
type Status int

const (
	StatusUnknown Status = iota // 0 - zero value, safe default
	StatusPending               // 1
	StatusActive                // 2
	StatusClosed                // 3
)

// go:generate stringer -type=Status
// running `go generate` creates status_string.go with a String() method
// install: go install golang.org/x/tools/cmd/stringer@latest

// typed constants prevent mixing unrelated iota sequences
func isMixSafe() {
	var d Direction = North
	var s Status = StatusPending
	// d == s  <- compile error: mismatched types Direction and Status
	_ = d
	_ = s
}

/*
iota patterns:
  iota                 0, 1, 2, 3 ...
  iota + 1             1, 2, 3, 4 ...
  1 << iota            1, 2, 4, 8 ... (bit flags)
  iota * 100           0, 100, 200 ...

Use _ to skip values:
  const (
      _ = iota  // skip 0
      One       // 1
      Two       // 2
  )

go:generate + stringer is the idiomatic way to auto-generate String() methods
for large enums instead of writing them by hand.
*/
