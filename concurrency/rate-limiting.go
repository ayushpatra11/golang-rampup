package concurrency

import (
	"fmt"
	"time"
)

// Rate limiting: control how frequently an operation can happen.
// Common use cases: API calls, DB queries, outbound requests.

// ---- Simple rate limiter with time.Ticker ----
// Process at most one event per interval.

func tickerRateLimiter() {
	requests := make(chan int, 10)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)

	// one request allowed every 200ms
	limiter := time.NewTicker(200 * time.Millisecond)
	defer limiter.Stop()

	for req := range requests {
		<-limiter.C // block until the next tick
		fmt.Println("request", req, "processed at", time.Now().Format("15:04:05.000"))
	}
}

// ---- Bursty rate limiter ----
// Allow short bursts, then fall back to a steady rate.
// Implemented as a buffered channel pre-filled with tokens.

func burstyLimiter(burstSize int, refillInterval time.Duration) <-chan time.Time {
	limiter := make(chan time.Time, burstSize)

	// pre-fill the burst capacity
	for i := 0; i < burstSize; i++ {
		limiter <- time.Now()
	}

	// refill one token at a time on a ticker
	go func() {
		ticker := time.NewTicker(refillInterval)
		defer ticker.Stop()
		for t := range ticker.C {
			limiter <- t // blocks when buffer is full (at capacity)
		}
	}()

	return limiter
}

func burstyDemo() {
	limiter := burstyLimiter(3, 200*time.Millisecond)

	for i := 0; i < 7; i++ {
		<-limiter
		fmt.Println("request", i, "at", time.Now().Format("15:04:05.000"))
		// first 3 requests go through immediately (burst)
		// remaining requests are throttled to one per 200ms
	}
}

// ---- Semaphore-based concurrency limiter ----
// Limit how many goroutines run simultaneously (not time-based).

func concurrencyLimiter(maxConcurrent int, tasks []func()) {
	sem := make(chan struct{}, maxConcurrent)
	done := make(chan struct{}, len(tasks))

	for _, task := range tasks {
		task := task
		go func() {
			sem <- struct{}{}        // acquire slot
			defer func() {
				<-sem               // release slot
				done <- struct{}{}
			}()
			task()
		}()
	}

	for range tasks {
		<-done
	}
}

// ---- golang.org/x/time/rate (token bucket) ----
// The standard choice for production rate limiting.
// Not in stdlib but maintained by the Go team.
//
// import "golang.org/x/time/rate"
//
// limiter := rate.NewLimiter(rate.Limit(10), 5)
//   rate.Limit(10) = 10 tokens per second
//   5              = burst capacity
//
// // blocking wait
// if err := limiter.Wait(ctx); err != nil { return err }
//
// // non-blocking check
// if !limiter.Allow() { return ErrRateLimited }
//
// // reserve and get wait duration
// r := limiter.Reserve()
// time.Sleep(r.Delay())

/*
Choosing an approach:
- time.Ticker:           simple throughput cap, no burst
- buffered channel:      allows burst, easy to reason about
- semaphore (chan):       limit concurrency, not rate
- golang.org/x/time/rate token bucket, most correct for production use
*/
