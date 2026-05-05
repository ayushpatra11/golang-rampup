// this is the file to learn strings in go

// string data type  

// length in go - len(str)

// can also use utf8.RuneCountInString(str) for the same thing

package string

import (
	"unicode/utf8"
	"fmt"
)

func main() {
	var str = "This is a tutorial personal repo for go"

	var length = len(str)

	fmt.Println("The length of the string is : ", length)
}

// we can also concat using +

// s := s1 + s2






