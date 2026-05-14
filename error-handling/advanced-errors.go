package errorhandling

import (
	"errors"
	"fmt"
	"log/slog"
)

// ---- Structured error types ----
// Carry rich context for logging and programmatic handling

type ErrorCode string

const (
	ErrCodeNotFound    ErrorCode = "NOT_FOUND"
	ErrCodeUnauth      ErrorCode = "UNAUTHORIZED"
	ErrCodeInternal    ErrorCode = "INTERNAL"
	ErrCodeBadRequest  ErrorCode = "BAD_REQUEST"
)

type ServiceError struct {
	Code    ErrorCode
	Message string
	Cause   error
}

func (e *ServiceError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func (e *ServiceError) Unwrap() error { return e.Cause }

func NewServiceError(code ErrorCode, msg string, cause error) *ServiceError {
	return &ServiceError{Code: code, Message: msg, Cause: cause}
}

// ---- Error handling at service boundaries ----

func getItem(id int) error {
	if id == 0 {
		return NewServiceError(ErrCodeBadRequest, "id must be non-zero", nil)
	}
	if id > 1000 {
		return NewServiceError(ErrCodeNotFound, fmt.Sprintf("item %d not found", id), nil)
	}
	return nil
}

func handleServiceError(err error) {
	var se *ServiceError
	if errors.As(err, &se) {
		switch se.Code {
		case ErrCodeNotFound:
			fmt.Println("404:", se.Message)
		case ErrCodeBadRequest:
			fmt.Println("400:", se.Message)
		default:
			fmt.Println("500:", se.Message)
		}
	} else {
		fmt.Println("unexpected:", err)
	}
}

// ---- Error groups (Go 1.20+) ----
// Collect multiple errors from concurrent operations

func validateAll(inputs []string) error {
	var errs []error
	for _, s := range inputs {
		if len(s) == 0 {
			errs = append(errs, fmt.Errorf("empty string not allowed"))
		}
		if len(s) > 100 {
			errs = append(errs, fmt.Errorf("string too long: %q", s))
		}
	}
	return errors.Join(errs...) // nil if errs is empty
}

// ---- Logging errors with slog (Go 1.21) ----

func loggedOperation(id int) error {
	err := getItem(id)
	if err != nil {
		var se *ServiceError
		if errors.As(err, &se) {
			slog.Error("operation failed",
				"code", se.Code,
				"message", se.Message,
				"id", id,
			)
		}
		return fmt.Errorf("loggedOperation: %w", err)
	}
	return nil
}

// ---- Error middleware pattern ----
// Wrapping errors at each layer with context

type Layer string

const (
	LayerHTTP    Layer = "http"
	LayerService Layer = "service"
	LayerDB      Layer = "db"
)

func wrapLayer(layer Layer, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", layer, err)
}

func dbQuery(id int) error {
	if id < 0 {
		return errors.New("negative id")
	}
	return nil
}

func serviceCall(id int) error {
	if err := dbQuery(id); err != nil {
		return wrapLayer(LayerDB, err)
	}
	return nil
}

func httpHandler(id int) error {
	if err := serviceCall(id); err != nil {
		return wrapLayer(LayerService, err)
	}
	return nil
}

/*
Error chain built by the above: "service: db: negative id"
errors.Is still traverses the full chain regardless of wrapping.

Key rules for error handling at scale:
1. Add context at each layer (fmt.Errorf + %w)
2. Use structured error types for machine-readable codes
3. Log once at the top-level handler, not at every layer
4. Return errors up; don't swallow them silently
5. errors.Join is great for validation (collect all errors, not just the first)
*/
