package types

import "fmt"

// type assertion: extract the concrete value from an interface
// syntax: value.(ConcreteType)
// panics if the interface doesn't hold that type

func typeAssertionBasic(i interface{}) {
	s := i.(string) // panics if i is not a string
	fmt.Println(s)
}

// safe form: two-value assignment (ok idiom)
func typeAssertionSafe(i interface{}) {
	s, ok := i.(string)
	if ok {
		fmt.Println("string:", s)
	} else {
		fmt.Println("not a string")
	}
}

// type switch: test against multiple types
// preferred over a chain of if-else type assertions
func describe(i interface{}) string {
	switch v := i.(type) {
	case int:
		return fmt.Sprintf("int: %d", v)
	case string:
		return fmt.Sprintf("string: %q (len=%d)", v, len(v))
	case bool:
		return fmt.Sprintf("bool: %v", v)
	case []int:
		return fmt.Sprintf("[]int with %d elements", len(v))
	case nil:
		return "nil"
	default:
		return fmt.Sprintf("unknown type: %T", v)
	}
}

// interface{} (any) can hold any value
func anyDemo() {
	values := []interface{}{42, "hello", true, 3.14, nil}
	for _, v := range values {
		fmt.Println(describe(v))
	}
}

// type assertion to check if a value implements an interface
type Stringer interface {
	String() string
}

func printIfStringer(v interface{}) {
	if s, ok := v.(Stringer); ok {
		fmt.Println("Stringer:", s.String())
	} else {
		fmt.Printf("not a Stringer: %T\n", v)
	}
}

// error type switch: common pattern for handling different error types
type NotFoundError struct{ Name string }
type PermissionError struct{ User string }

func (e *NotFoundError) Error() string   { return e.Name + " not found" }
func (e *PermissionError) Error() string { return e.User + " has no permission" }

func handleErr(err error) {
	switch e := err.(type) {
	case *NotFoundError:
		fmt.Println("missing resource:", e.Name)
	case *PermissionError:
		fmt.Println("access denied for:", e.User)
	default:
		fmt.Println("unexpected error:", err)
	}
}

/*
Prefer errors.As() over type switches for error handling (handles wrapped errors).
Type switches are useful for general interface{}/any dispatch.

%T verb prints the dynamic type of an interface value.
*/
