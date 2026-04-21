package stdlib

import (
	"fmt"
	"time"
)

// ---- time.Time ----
// time.Time is immutable value type; methods return new values

func timeBasics() {
	now := time.Now()
	fmt.Println(now)                    // 2026-04-20 10:00:00.123456789 +0000 UTC
	fmt.Println(now.Year())             // 2026
	fmt.Println(now.Month())            // April
	fmt.Println(now.Day())              // 20
	fmt.Println(now.Weekday())          // Monday
	fmt.Println(now.Hour(), now.Minute(), now.Second())
	fmt.Println(now.Unix())             // unix seconds since epoch
	fmt.Println(now.UnixMilli())        // milliseconds
	fmt.Println(now.UnixNano())         // nanoseconds
}

// ---- Formatting and parsing ----
// Go uses a reference time: Mon Jan 2 15:04:05 MST 2006
// That specific moment is the "layout" - you write your format using those values

func formatTime() {
	t := time.Now()

	// format using reference time
	fmt.Println(t.Format("2006-01-02"))                          // 2026-04-20
	fmt.Println(t.Format("2006-01-02 15:04:05"))                 // ISO datetime
	fmt.Println(t.Format("Jan 2, 2006"))                         // Apr 20, 2026
	fmt.Println(t.Format(time.RFC3339))                          // 2026-04-20T10:00:00Z
	fmt.Println(t.Format(time.RFC3339Nano))                      // with nanoseconds
	fmt.Println(t.Format(time.Kitchen))                          // 10:00AM
}

func parseTime() {
	t, err := time.Parse("2006-01-02", "2026-04-20")
	if err != nil {
		panic(err)
	}
	fmt.Println(t) // 2026-04-20 00:00:00 +0000 UTC

	// parse in a specific timezone
	loc, _ := time.LoadLocation("America/New_York")
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2026-04-20 10:00:00", loc)
	fmt.Println(t2)
}

// ---- Duration ----
// time.Duration is int64 nanoseconds

func durationDemo() {
	d := 2*time.Hour + 30*time.Minute + 15*time.Second
	fmt.Println(d)                // 2h30m15s
	fmt.Println(d.Hours())        // 2.504166...
	fmt.Println(d.Minutes())      // 150.25
	fmt.Println(d.Seconds())      // 9015
	fmt.Println(d.Milliseconds()) // 9015000

	// parse from string
	d2, _ := time.ParseDuration("1h30m")
	fmt.Println(d2) // 1h30m0s

	// arithmetic
	t := time.Now()
	future := t.Add(24 * time.Hour)
	past := t.Add(-7 * 24 * time.Hour)
	fmt.Println(future.Sub(past)) // 168h0m0s (one week)
}

// ---- Timers and tickers ----

func timerDemo() {
	// one-shot timer
	timer := time.NewTimer(100 * time.Millisecond)
	<-timer.C // blocks until timer fires
	fmt.Println("timer fired")

	// time.After: shorthand (can't be stopped/reset, may leak for short durations)
	<-time.After(50 * time.Millisecond)

	// stop a timer to prevent it firing (returns false if already fired)
	t2 := time.NewTimer(1 * time.Second)
	t2.Stop()
}

func tickerDemo() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop() // ALWAYS stop ticker to avoid goroutine leak

	count := 0
	for range ticker.C {
		count++
		fmt.Println("tick", count)
		if count >= 3 {
			break
		}
	}
}

// ---- Timezone ----

func timezoneDemo() {
	loc, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		panic(err)
	}
	t := time.Now().In(loc)
	fmt.Println(t.Format("2006-01-02 15:04:05 MST"))

	utc := t.UTC()
	fmt.Println(utc)
}

// measuring elapsed time
func measure(fn func()) time.Duration {
	start := time.Now()
	fn()
	return time.Since(start)
}

/*
Reference time: Mon Jan 2 15:04:05 MST 2006
  Year:  2006
  Month: 01 (or Jan)
  Day:   02
  Hour:  15 (24h) or 3 (12h)
  Min:   04
  Sec:   05
  Zone:  MST or -0700

time.Sleep vs ticker: use ticker for repeated work, Sleep for one-shot delays.
Always defer ticker.Stop() to release the goroutine and timer resources.
*/
