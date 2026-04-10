package concurrency

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// goroutine: lightweight thread managed by the Go runtime
// Go can run hundreds of thousands of goroutines concurrently
// "go" keyword starts a new goroutine

func goroutineBasic() {
	go fmt.Println("Hello from goroutine")
	time.Sleep(10 * time.Millisecond) // naive sync - don't do this in real code
}

// WaitGroup: wait for a collection of goroutines to finish
func waitGroupDemo() {
	var wg sync.WaitGroup
	names := []string{"alice", "bob", "charlie"}

	for _, name := range names {
		wg.Add(1) // increment before starting goroutine
		go func(n string) {
			defer wg.Done() // decrement when goroutine exits
			fmt.Println("processing:", n)
		}(name) // pass name as arg to avoid closure capture bug
	}

	wg.Wait() // block until count reaches zero
	fmt.Println("all done")
}

// Mutex: protect shared state from concurrent access
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func mutexDemo() {
	c := &SafeCounter{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	wg.Wait()
	fmt.Println(c.Value()) // 1000
}

// RWMutex: multiple readers OR one writer (more efficient for read-heavy workloads)
type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.data[key]
	return v, ok
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

// atomic: lock-free operations on individual values (faster than mutex for simple cases)
type AtomicCounter struct {
	count atomic.Int64
}

func (c *AtomicCounter) Inc()          { c.count.Add(1) }
func (c *AtomicCounter) Value() int64  { return c.count.Load() }

// once: run initialization exactly once (goroutine safe)
var (
	once     sync.Once
	instance *Cache
)

func getInstance() *Cache {
	once.Do(func() {
		instance = &Cache{data: make(map[string]string)}
	})
	return instance
}

/*
Go's concurrency mantra (Rob Pike):
"Do not communicate by sharing memory;
instead, share memory by communicating."

Prefer channels when goroutines need to pass data.
Use mutexes when goroutines share a data structure they both read/write.

Detecting data races: go test -race ./...
*/
