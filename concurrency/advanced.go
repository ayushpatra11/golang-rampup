package concurrency

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// ---- sync.Cond: condition variable ----
// Used when a goroutine needs to wait for a condition to become true.
// Prefer channels for most cases; Cond is useful for broadcast scenarios.

type Queue struct {
	mu    sync.Mutex
	cond  *sync.Cond
	items []int
}

func NewQueue() *Queue {
	q := &Queue{}
	q.cond = sync.NewCond(&q.mu)
	return q
}

func (q *Queue) Push(v int) {
	q.mu.Lock()
	q.items = append(q.items, v)
	q.cond.Signal() // wake one waiting goroutine
	q.mu.Unlock()
}

func (q *Queue) Pop() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	for len(q.items) == 0 {
		q.cond.Wait() // atomically releases lock and suspends
	}
	v := q.items[0]
	q.items = q.items[1:]
	return v
}

// cond.Broadcast() wakes ALL waiting goroutines (vs Signal which wakes one)

// ---- sync.Map ----
// Thread-safe map optimized for specific patterns:
// - Write-once, read-many keys (stable key set)
// - Disjoint key sets accessed by different goroutines
// For general use, a regular map + RWMutex is usually clearer.

func syncMapDemo() {
	var m sync.Map

	m.Store("key", 42)

	val, ok := m.Load("key")
	fmt.Println(val, ok) // 42 true

	// LoadOrStore: load if exists, otherwise store and return new value
	actual, loaded := m.LoadOrStore("key", 99)
	fmt.Println(actual, loaded) // 42 true (was already there)

	m.Delete("key")

	// Range: iterate over all key-value pairs
	m.Store("a", 1)
	m.Store("b", 2)
	m.Range(func(k, v any) bool {
		fmt.Println(k, v)
		return true // return false to stop iteration
	})
}

// ---- Timeouts on goroutine operations ----

func withOperationTimeout(timeout time.Duration, fn func() error) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	done := make(chan error, 1)
	go func() {
		done <- fn()
	}()

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ---- Limiting goroutine lifetime ----
// Always ensure goroutines can be stopped to prevent leaks.

func startWorker(ctx context.Context, jobs <-chan int) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("worker shutting down:", ctx.Err())
				return
			case j, ok := <-jobs:
				if !ok {
					fmt.Println("jobs channel closed")
					return
				}
				fmt.Println("processing job:", j)
			}
		}
	}()
}

// ---- Channel ownership ----
// Owner: the goroutine that creates, writes to, and closes a channel
// Consumer: reads from the channel

// Clear ownership prevents:
// - Multiple goroutines closing the same channel (panic)
// - Sending on a closed channel (panic)
// - Goroutine leaks (writer can't send because reader quit)

func ownedChannel() <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)           // owner closes
		for i := 0; i < 5; i++ {
			out <- i
		}
	}()
	return out // consumer gets read-only channel
}

/*
Goroutine leak checklist:
1. Every goroutine needs a way to exit (context, done channel, or channel close)
2. Goroutines blocked on channels exit when the channel is closed
3. Use go test -race to detect data races
4. goleak (github.com/uber-go/goleak) detects leaked goroutines in tests

Rule: "Never start a goroutine without knowing how it will stop."
*/
