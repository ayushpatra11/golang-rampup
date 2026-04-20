package concurrency

import (
	"fmt"
	"sync"
)

// ---- WORKER POOL ----
// Fixed number of goroutines processing a job queue.
// Controls parallelism without spinning up a goroutine per task.

func workerPool(numWorkers, numJobs int) {
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// start workers
	for w := 0; w < numWorkers; w++ {
		go func() {
			for j := range jobs {
				results <- j * j // simulate work
			}
		}()
	}

	// enqueue jobs
	for j := 0; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs) // signal no more jobs

	// collect results
	for i := 0; i < numJobs; i++ {
		fmt.Println(<-results)
	}
}

// ---- PIPELINE ----
// Chain stages connected by channels; each stage is a set of goroutines.

func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

func pipelineDemo() {
	c := generate(2, 3, 4, 5)
	out := square(square(c))
	for v := range out {
		fmt.Println(v) // 16, 81, 256, 625
	}
}

// ---- FAN-OUT / FAN-IN ----
// Fan-out: distribute work from one channel to multiple goroutines.
// Fan-in: merge multiple channels into one.

func fanIn(channels ...<-chan int) <-chan int {
	merged := make(chan int)
	var wg sync.WaitGroup

	pipe := func(c <-chan int) {
		defer wg.Done()
		for v := range c {
			merged <- v
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go pipe(c)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

// ---- ERRGROUP (manual implementation) ----
// Run goroutines and collect the first non-nil error.
// In practice use golang.org/x/sync/errgroup.

type ErrGroup struct {
	wg   sync.WaitGroup
	once sync.Once
	err  error
}

func (g *ErrGroup) Go(fn func() error) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		if err := fn(); err != nil {
			g.once.Do(func() { g.err = err }) // keep first error only
		}
	}()
}

func (g *ErrGroup) Wait() error {
	g.wg.Wait()
	return g.err
}

// ---- SEMAPHORE ----
// Limit the number of concurrent goroutines using a buffered channel.

func semaphore(maxConcurrent int, tasks []func()) {
	sem := make(chan struct{}, maxConcurrent)
	var wg sync.WaitGroup

	for _, task := range tasks {
		wg.Add(1)
		task := task
		go func() {
			defer wg.Done()
			sem <- struct{}{}        // acquire
			defer func() { <-sem }() // release
			task()
		}()
	}
	wg.Wait()
}

/*
Pattern selection guide:
- Worker pool:   bounded concurrency on a stream of identical jobs
- Pipeline:      staged data transformation
- Fan-out/in:    parallel work on the same input, merge results
- Semaphore:     rate-limit access to a resource (DB connections, API calls)
- errgroup:      parallel subtasks where any failure should abort
*/
