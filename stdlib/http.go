package stdlib

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// ---- HTTP SERVER ----
// net/http: standard library server
// http.Handler interface: ServeHTTP(ResponseWriter, *Request)

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// handler function: must match func(ResponseWriter, *Request) signature
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

// Go 1.22+: method + path pattern in ServeMux
// "GET /users/{id}" - method-based routing with path parameters
func userHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id") // Go 1.22+
	writeJSON(w, http.StatusOK, map[string]string{"id": id})
}

func startServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", helloHandler)
	mux.HandleFunc("GET /users/{id}", userHandler)
	mux.HandleFunc("POST /users", createUserHandler)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      withLogging(mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Println("listening on :8080")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	writeJSON(w, http.StatusCreated, map[string]string{"name": body.Name})
}

// middleware: wrap a Handler with pre/post logic
func withLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("%s %s %v\n", r.Method, r.URL.Path, time.Since(start))
	})
}

// ---- HTTP CLIENT ----

var httpClient = &http.Client{
	Timeout: 10 * time.Second, // always set a timeout
}

// GET request
func getJSON(ctx context.Context, url string, target any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("build request: %w", err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close() // always close body

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status %d: %s", resp.StatusCode, body)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

/*
Server tips:
- Always set read/write timeouts on http.Server
- Use http.Error() for simple text error responses
- Use Go 1.22 method routing: "GET /path" instead of matching in handler
- Middleware chains via wrapper functions (no framework needed for simple cases)

Client tips:
- Always set a Timeout on http.Client (default has none)
- Always close resp.Body even on error status
- Use http.NewRequestWithContext to pass cancellation
- Reuse http.Client (has connection pooling)
- Never use http.DefaultClient in production (no timeout)
*/
