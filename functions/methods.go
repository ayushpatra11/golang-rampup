package functions

import (
	"fmt"
	"math"
)

// methods: functions with a receiver argument
// receiver can be any named type defined in the same package

type Circle struct {
	Radius float64
}

// value receiver: works on a copy, cannot modify the original
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// pointer receiver: can modify the original; also avoids copying large structs
func (c *Circle) Scale(factor float64) {
	c.Radius *= factor
}

func methodDemo() {
	c := Circle{Radius: 5}
	fmt.Printf("Area: %.2f\n", c.Area()) // ~78.54

	c.Scale(2)                            // Go auto-takes address: (&c).Scale(2)
	fmt.Println(c.Radius)                 // 10

	// pointer variable also works
	cp := &Circle{Radius: 3}
	fmt.Printf("Area: %.2f\n", cp.Area()) // cp is auto-dereferenced: (*cp).Area()
}

// methods on non-struct named types
type Celsius float64
type Fahrenheit float64

func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func (f Fahrenheit) ToCelsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// string representation - implement Stringer interface
func (c Circle) String() string {
	return fmt.Sprintf("Circle(r=%.2f)", c.Radius)
}

// method expressions and method values
func methodValueDemo() {
	c := Circle{Radius: 4}

	// method value: bound to a specific receiver
	areaFn := c.Area        // type: func() float64
	fmt.Println(areaFn())   // same as c.Area()

	// method expression: receiver passed explicitly
	areaExpr := Circle.Area // type: func(Circle) float64
	fmt.Println(areaExpr(c))
}

// embedding promotes methods to the outer type
type Animal struct {
	Name string
}

func (a Animal) Describe() string {
	return "I am " + a.Name
}

type Dog struct {
	Animal        // embedded (anonymous field) - not a named field
	Breed string
}

func embeddingDemo() {
	d := Dog{Animal: Animal{Name: "Rex"}, Breed: "Labrador"}
	fmt.Println(d.Describe()) // promoted: d.Animal.Describe()
	fmt.Println(d.Name)       // promoted: d.Animal.Name
}

/*
Rule of thumb: if ANY method on a type has a pointer receiver,
make ALL methods pointer receivers. Mixed receivers cause subtle
bugs when the type is used as an interface value.

Pointer receiver is needed when:
- The method mutates the receiver
- The receiver is a large struct (avoid copying)
- Consistency with other methods on the same type
*/
