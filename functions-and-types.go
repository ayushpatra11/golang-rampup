package functions

import "fmt"

/* syntax - 

func functionName(param1, param2, ...) {    -> not having this bracket here will give an error
	...
}

func functionName(param1, ...) type1 {
	...
}

func functionName(param1, ...) (ret1 type1, ret2 type2) {
	...
}


 main is required as a starting point for a compilable -> has no params and no return values  




 */

func sumOfTwoNumbers (a int, b int) int {
	return a+b
}

/*
Types

var var1 type1

func FunctionName (a typea, b typeb) typeFunc
return var1

func FunctionName (a typea, b typeb) (t1 type1, t2 type2)
return var1, var2

 CAN Also declare an alias for a type -

type IZ int 

var a IZ = 5 // a of type int 

type (
	a int
	b float
	c bool
	s string
)

can also type cast using ()

*/

