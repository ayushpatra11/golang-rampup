package errorhandling

import (
	"errors"
	"fmt"
)

// error is just an interface:
//   type error interface { Error() string }
// any type implementing Error() string satisfies it

// sentinel errors: package-level vars for expected conditions
var (
	ErrNotFound   = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
)

// errors.New: simple string error
func findUser(id int) (string, error) {
	if id <= 0 {
		return "", ErrNotFound
	}
	return "alice", nil
}

// fmt.Errorf with %w: wraps an error, preserving the chain
func fetchUserProfile(id int) (string, error) {
	name, err := findUser(id)
	if err != nil {
		return "", fmt.Errorf("fetchUserProfile(id=%d): %w", id, err)
	}
	return "profile:" + name, nil
}

// errors.Is: checks whether any error in the chain matches target
// works through wrapping (unlike == comparison)
func checkIs() {
	_, err := fetchUserProfile(-1)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("resource not found") // prints this
	}
}

// custom error type: carries structured data
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s - %s", e.Field, e.Message)
}

func validateAge(age int) error {
	if age < 0 || age > 150 {
		return &ValidationError{Field: "age", Message: "must be 0-150"}
	}
	return nil
}

// errors.As: checks if any error in the chain matches the target type
// and unwraps into it
func checkAs() {
	err := validateAge(-5)
	var ve *ValidationError
	if errors.As(err, &ve) {
		fmt.Println("bad field:", ve.Field) // age
	}
}

// wrapping preserves both the chain and the type
func wrappedValidation(age int) error {
	if err := validateAge(age); err != nil {
		return fmt.Errorf("user creation failed: %w", err)
	}
	return nil
}

// multiple error wrapping (Go 1.20+): fmt.Errorf with multiple %w
func multiWrap(e1, e2 error) error {
	return fmt.Errorf("two errors: %w and %w", e1, e2)
}

// errors.Join (Go 1.20+): combines multiple errors
func joinErrors(errs ...error) error {
	return errors.Join(errs...)
}

/*
Go error handling philosophy:
- Errors are explicit return values - no hidden exceptions
- Check each error where it can arise
- Wrap with context using %w to preserve the chain
- Use errors.Is for sentinel comparison (not ==)
- Use errors.As for typed extraction
- Return nil when there is no error (don't create zero-value error structs)
*/
