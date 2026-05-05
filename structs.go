package main

import (
	"fmt"
)

type car struct {
	Make string
	Model string
	Height int 
	Width int
}

func test (c car){
	fmt.Println(c.Make, c.Model, c.Height, c.Width)
}

func main (){
	test(car{
		Make: "Toyota",
		Model: "Camry",
		Height: 2300,
		Width: 2311,
	})
}

/*

We can also have nested structs -> we can have variable objects of other structs inside other structs

accessing it is like the same as in C.

*/

// anonymous structs : struct instances that do not have a name

myCar := struct {
	/*
	.
	.
	.
	*/
}

// this is the same as having an object being declared of the Car struct. 

// generally avoid anon structs, unless it is an ephemeral object that has to be used only once. 


