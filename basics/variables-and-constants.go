package varsandconsts

const abc = "Hello"  // allowed

import (
	"fmt"
	"math/rand"
)

const RESCODES = (
	SUCCESS
	FAIL
	PENDING
)

// no overflow in Go

const Ln2= 0.693147180559945309417232121458176568075500134360255254120680009
const Log2E= 1/Ln2 // this is a precise reciprocal
const BILLION = 1e9 // float constant
const HARD_EIGHT = (1 << 100) >> 97

const CHICKEN, TWO, C = "meat", 2, "veg"

type Gender int
const (
  UNKNOWN Gender = iota
  FEMALE
  MALE
)

var num int = 5

num := 5 // same thing and mostly this is used

// pointers in go:

x := 10
p := &x   // p is a pointer to x

fmt.Println(*p) // 10

*p = 20
fmt.Println(x)  // 20 (modified via pointer)

// no memory management and no incr or decr operations on the pointer itself  

/*

There are a few types that are new from C++

For integers:

int8 (-128 to 127)
int16 (-32768 to 32767)
int32 (− 2,147,483,648 to 2,147,483,647)
int64 (− 9,223,372,036,854,775,808 to 9,223,372,036,854,775,807)

For unsigned integers:

uint8 (with the alias byte, 0 to 255)
uint16 (0 to 65,535)
uint32 (0 to 4,294,967,295)
uint64 (0 to 18,446,744,073,709,551,615)

For floats:

NOTE: There is NO float by it's own in Go

float32 
float64 


NOTE: As Go is strongly typed, the mixing of types is not allowed, as in the following program. However, constants are considered to have no type in this respect. Therefore, with constants, mixing is allowed.

*/

func mixTypes(a int, b int32) {
	b = b + 5 // allowed
	b = a + a // not allowed -> will throw a compiler error

	// we can however explicitly typecaste the exact same way we used to do in C++

	var n int16 = 34    // int16 variable
	var m int32         // int32 variable
	m = int32(n)        // explicit typing
}

func genRandomNumProd() int {
	
	a := rand.Int()
	b := rand.Intn(8) 		// will go from 0 till 8 (not including 8) -> [0, n)

	return a*b
}


// for characters as well, using byte/char/hexa value will lead to same assignment: 

var ch byte = 'A'
var ch byte = 65
var ch byte = '\x41'


// checking if isNum/letter and so on, we do it using Unicode -> verify if there has to be a utf8 package that has to be imported 

unicode.IsLetter(ch)
unicode.IsDigit(ch)
unicode.IsSpace(ch)


