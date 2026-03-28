//package: a library, module or namespace  

// basically, each file is being constructed as a package  

package main // this just means that we are exporting the main package (where the code execution will start)

// package utils -> this would mean that we are exporting a utils package.

// you can import more packages and use them inside this package  

// every go file is a single package basically

/*

Think of a Go application as a book:

The book is your application.
Each chapter is a package.
Each page is a source file.
Even if the entire book is just one chapter (package main), you don’t have to write the whole chapter on one page. You can spread it across many pages (files), making it easier to read and edit.

If you have other chapters (pack1, pack2), they are separate sections that can be compiled independently and then linked together to form the complete book.

NOTE: main package will always give an executable, but it can be produced from different source files together to compile from a more modular code  

*/
// #### IMPORT ####

import "fmt"
import fm "fmt" // to give an alias to fmt
import "os" 
import "pack1"

// or can be written as 

import (
	"fmt"
	"os"
)

// #### VISIBILTY ####

// Capital letters are visible externally, others are not 

pack1.Object // to call Object  
