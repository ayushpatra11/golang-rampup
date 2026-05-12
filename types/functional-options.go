package types

import (
	"fmt"
	"time"
)

// Functional options pattern: a clean way to handle optional config for structs.
// Avoids large constructor parameter lists and boolean flags.

type ServerConfig struct {
	host        string
	port        int
	timeout     time.Duration
	maxRetries  int
	debug       bool
}

// Option is a function that modifies a ServerConfig
type Option func(*ServerConfig)

// Option constructors
func WithHost(host string) Option {
	return func(c *ServerConfig) {
		c.host = host
	}
}

func WithPort(port int) Option {
	return func(c *ServerConfig) {
		c.port = port
	}
}

func WithTimeout(d time.Duration) Option {
	return func(c *ServerConfig) {
		c.timeout = d
	}
}

func WithMaxRetries(n int) Option {
	return func(c *ServerConfig) {
		c.maxRetries = n
	}
}

func WithDebug(debug bool) Option {
	return func(c *ServerConfig) {
		c.debug = debug
	}
}

// Constructor applies options on top of defaults
func NewServerConfig(opts ...Option) *ServerConfig {
	cfg := &ServerConfig{
		host:       "localhost",
		port:       8080,
		timeout:    30 * time.Second,
		maxRetries: 3,
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

func functionalOptionsDemo() {
	// use only defaults
	cfg1 := NewServerConfig()
	fmt.Printf("%+v\n", cfg1)

	// override some
	cfg2 := NewServerConfig(
		WithHost("0.0.0.0"),
		WithPort(9090),
		WithDebug(true),
	)
	fmt.Printf("%+v\n", cfg2)
}

// ---- Builder pattern alternative ----
// Less idiomatic in Go; functional options are preferred for libraries.

type QueryBuilder struct {
	table  string
	where  []string
	limit  int
	offset int
}

func NewQuery(table string) *QueryBuilder {
	return &QueryBuilder{table: table, limit: 100}
}

func (q *QueryBuilder) Where(cond string) *QueryBuilder {
	q.where = append(q.where, cond)
	return q
}

func (q *QueryBuilder) Limit(n int) *QueryBuilder {
	q.limit = n
	return q
}

func (q *QueryBuilder) Offset(n int) *QueryBuilder {
	q.offset = n
	return q
}

func (q *QueryBuilder) Build() string {
	sql := fmt.Sprintf("SELECT * FROM %s", q.table)
	for i, w := range q.where {
		if i == 0 {
			sql += " WHERE " + w
		} else {
			sql += " AND " + w
		}
	}
	if q.limit > 0 {
		sql += fmt.Sprintf(" LIMIT %d", q.limit)
	}
	if q.offset > 0 {
		sql += fmt.Sprintf(" OFFSET %d", q.offset)
	}
	return sql
}

func builderDemo() {
	q := NewQuery("users").
		Where("age > 18").
		Where("active = true").
		Limit(10).
		Offset(20).
		Build()
	fmt.Println(q)
	// SELECT * FROM users WHERE age > 18 AND active = true LIMIT 10 OFFSET 20
}

/*
Functional options vs config struct vs many parameters:

Many parameters:      NewServer("host", 8080, 30*time.Second, 3, false)  <- unreadable
Config struct:        NewServer(Config{Host: "host", Port: 8080})         <- verbose, no defaults
Functional options:   NewServer(WithHost("host"), WithPort(8080))         <- clean, extensible, backwards-compatible

Functional options are preferred in Go library code because:
- Adding new options doesn't break existing callers
- Options self-document what they do
- Defaults are clearly defined in the constructor
*/
