package concurrency

import (
	"context"
	"fmt"
	"time"
)

// context.Context: carries deadlines, cancellation signals, and request-scoped values
// Pass ctx as the FIRST argument by convention

// Root contexts:
//   context.Background() - root, never cancelled (use at main/top level)
//   context.TODO()       - placeholder while refactoring, same as Background

// ---- WithCancel ----
// Returns a copy of the parent with a cancel function.
// Calling cancel() or the parent being cancelled triggers ctx.Done().

func withCancelDemo() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // always defer cancel to avoid resource leak

	done := make(chan struct{})
	go func() {
		defer close(done)
		select {
		case <-ctx.Done():
			fmt.Println("cancelled:", ctx.Err()) // context.Canceled
		case <-time.After(5 * time.Second):
			fmt.Println("finished naturally")
		}
	}()

	time.Sleep(100 * time.Millisecond)
	cancel()  // trigger cancellation
	<-done    // wait for goroutine to exit
}

// ---- WithTimeout / WithDeadline ----
// Automatically cancels after a duration or at an absolute time.

func withTimeoutDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("timeout:", ctx.Err()) // context.DeadlineExceeded
	}
}

// check if context is still valid before expensive work
func doExpensiveWork(ctx context.Context) error {
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			// do a chunk of work
			time.Sleep(10 * time.Millisecond)
		}
	}
	return nil
}

// ---- WithValue ----
// Attaches request-scoped values (trace IDs, user info, etc.)
// Use sparingly - not for passing optional function parameters.

// use unexported custom key types to avoid collisions between packages
type contextKey string

const (
	requestIDKey contextKey = "requestID"
	userIDKey    contextKey = "userID"
)

func withRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIDKey, id)
}

func getRequestID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(requestIDKey).(string)
	return id, ok
}

func handleRequest(ctx context.Context) {
	id, ok := getRequestID(ctx)
	if ok {
		fmt.Println("handling request:", id)
	}
}

// ---- propagating context through call chains ----

func A(ctx context.Context) error {
	return B(ctx)
}

func B(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		fmt.Println("B doing work")
		return nil
	}
}

/*
Rules for context:
1. Always pass ctx as the first parameter: func F(ctx context.Context, ...)
2. Never store ctx in a struct field - pass it explicitly
3. Never pass nil context - use context.Background() or context.TODO()
4. Cancel functions must always be called (defer cancel())
5. WithValue keys must be unexported custom types
6. Don't use context for passing optional function parameters

context.Err() returns:
  context.Canceled        if cancelled via cancel()
  context.DeadlineExceeded if deadline/timeout elapsed
*/
