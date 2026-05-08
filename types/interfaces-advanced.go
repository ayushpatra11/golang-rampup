package types

import (
	"fmt"
	"io"
	"math"
	"sort"
)

// ---- Interface satisfaction is implicit ----
// No "implements" keyword - a type satisfies an interface by having the methods

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct{ Width, Height float64 }
type CircleShape struct{ Radius float64 }

func (r Rect) Area() float64       { return r.Width * r.Height }
func (r Rect) Perimeter() float64  { return 2 * (r.Width + r.Height) }
func (c CircleShape) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c CircleShape) Perimeter() float64 { return 2 * math.Pi * c.Radius }

func totalArea(shapes []Shape) float64 {
	total := 0.0
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}

// ---- Small interfaces are better ----
// "The bigger the interface, the weaker the abstraction." - Rob Pike
// Prefer single-method interfaces when possible

type Stringer interface{ String() string }
type Sizer interface{ Size() int }

// ---- Interface composition ----
type ReadWriteCloser interface {
	io.Reader
	io.Writer
	io.Closer
}

// ---- Interface values ----
// An interface value holds (type, value) - the dynamic type and its value
// Two interface values are equal if both have the same dynamic type AND value
// A nil interface != an interface holding a nil pointer (common gotcha!)

func nilInterfaceGotcha() {
	var err error              // nil interface (both type and value are nil)
	fmt.Println(err == nil)    // true

	var p *fmt.Stringer         // (*fmt.Stringer)(nil)
	var i interface{} = p       // interface holds (type=*fmt.Stringer, value=nil)
	fmt.Println(i == nil)       // FALSE! type part is non-nil
}

// ---- Implementing standard interfaces ----

// fmt.Stringer: controls fmt.Print output
type Temperature struct {
	Celsius float64
}

func (t Temperature) String() string {
	return fmt.Sprintf("%.1f°C", t.Celsius)
}

// sort.Interface: allows sort.Sort to sort your type
type ByLength []string

func (b ByLength) Len() int           { return len(b) }
func (b ByLength) Less(i, j int) bool { return len(b[i]) < len(b[j]) }
func (b ByLength) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func sortByLength(words []string) {
	sort.Sort(ByLength(words))
	fmt.Println(words)
}

// error interface
type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

// ---- Interface for testing / dependency injection ----
// Accept interfaces, return structs

type EmailSender interface {
	Send(to, subject, body string) error
}

type UserService struct {
	mailer EmailSender
}

func NewUserService(mailer EmailSender) *UserService {
	return &UserService{mailer: mailer}
}

func (s *UserService) Register(email string) error {
	// ... create user ...
	return s.mailer.Send(email, "Welcome!", "Thanks for signing up.")
}

// MockMailer for tests
type MockMailer struct {
	Sent []string
}

func (m *MockMailer) Send(to, subject, body string) error {
	m.Sent = append(m.Sent, to)
	return nil
}

/*
Interface design principles:
1. Keep interfaces small (1-3 methods ideally)
2. Accept interfaces; return concrete types
3. Define interfaces where they're used, not where types are defined
4. Don't create interfaces "for the future" - wait until you have two implementations
5. Naming: single-method interfaces use method name + "er" (Reader, Writer, Stringer)
*/
