package concurrency

import (
	"fmt"
	"time"
)

// channel: typed conduit between goroutines
// zero value is nil - sending/receiving on nil channel blocks forever
// sending on a closed channel panics; receiving returns zero value + false

// make(chan T)     -> unbuffered: send blocks until receiver is ready
// make(chan T, n)  -> buffered:   send blocks only when buffer is full

func unbufferedChannel() {
	ch := make(chan int)
	go func() {
		ch <- 42 // blocks until main receives
	}()
	val := <-ch // blocks until goroutine sends
	fmt.Println(val)
}

func bufferedChannel() {
	ch := make(chan string, 3)
	ch <- "a" // doesn't block (buffer not full)
	ch <- "b"
	ch <- "c"
	// ch <- "d" would block (buffer full until someone reads)

	fmt.Println(<-ch) // "a"
	fmt.Println(<-ch) // "b"
}

// close signals that no more values will be sent
// receivers can detect closure with the ok idiom or range
func closeDemo() {
	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)

	// ok idiom
	for {
		v, ok := <-ch
		if !ok {
			break // channel closed and empty
		}
		fmt.Println(v)
	}
}

// range over channel: exits when channel is closed
func rangeChannel() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) // MUST close to exit the range loop
	}()
	for v := range ch {
		fmt.Println(v)
	}
}

// directional channel types enforced at compile time
func producer(out chan<- int, n int) { // send-only
	for i := 0; i < n; i++ {
		out <- i
	}
	close(out)
}

func consumer(in <-chan int) { // receive-only
	for v := range in {
		fmt.Println("received:", v)
	}
}

// select: wait on multiple channel operations simultaneously
// picks a ready case at random if multiple are ready
func selectDemo() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	go func() { time.Sleep(1 * time.Millisecond); ch1 <- "from ch1" }()
	go func() { time.Sleep(2 * time.Millisecond); ch2 <- "from ch2" }()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		}
	}
}

// non-blocking receive with default
func nonBlocking(ch <-chan int) {
	select {
	case v := <-ch:
		fmt.Println("received:", v)
	default:
		fmt.Println("no value ready")
	}
}

// timeout with time.After
func withTimeout(ch <-chan string, timeout time.Duration) {
	select {
	case msg := <-ch:
		fmt.Println("got:", msg)
	case <-time.After(timeout):
		fmt.Println("timed out")
	}
}

// done channel: cancellation signal (prefer context.Context in real code)
func doWork(done <-chan struct{}, jobs <-chan int) {
	for {
		select {
		case <-done:
			fmt.Println("cancelled")
			return
		case j, ok := <-jobs:
			if !ok {
				return
			}
			fmt.Println("working on", j)
		}
	}
}

/*
Channel direction summary:
  chan T    bidirectional (full access)
  chan<- T  send-only
  <-chan T  receive-only

Closing rules:
- Only the sender should close a channel
- Never close a channel more than once (panic)
- Never send on a closed channel (panic)
- Use sync.WaitGroup + a single "done" goroutine to coordinate closing
*/
