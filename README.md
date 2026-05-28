# golang-rampup
Tracking my progress learning Go for systems engineering.

## References
- https://go.dev/ref/spec — Language specification
- https://go.dev/doc/tutorial — Official tutorial
- https://go.dev/doc/effective_go — Effective Go
- https://www.educative.io/courses/the-way-to-go — Course
- https://gobyexample.com — Examples by topic
- https://pkg.go.dev/std — Standard library docs

## Structure

```
basics/           variables, constants, types, strings, control flow,
                  arrays, slices, maps, pointers, io interfaces,
                  modules, memory/performance, idiomatic patterns
functions/        functions, closures, methods, init, variadic
types/            structs, interfaces, generics, type assertions,
                  embedding, functional options, iota/enums
error-handling/   errors, defer, panic, recover, advanced patterns
concurrency/      goroutines, channels, select, sync primitives,
                  context, patterns, advanced patterns, rate limiting
stdlib/           file I/O, JSON, HTTP, testing, strings/strconv,
                  time, regex, slices/maps packages, os, slog, flags,
                  encoding (csv, base64, hex)
```

## Progress

- [x] Week 1 (Mar 28–Apr 3)  — basics: variables, types, control flow, arrays, slices, maps
- [x] Week 2 (Apr 6–10)      — pointers, functions, closures, methods, structs, interfaces
- [x] Week 3 (Apr 12–16)     — embedding, generics, type assertions, error handling, goroutines
- [x] Week 4 (Apr 20–24)     — channels, concurrency patterns, context, file I/O
- [x] Week 5 (Apr 27–May 1)  — JSON, HTTP, testing, strings, time, regexp
- [x] Week 6 (May 3–7)       — io interfaces, modules, slices/maps packages, functional options
- [x] Week 7 (May 12–14)     — advanced interfaces, advanced errors, slog
- [x] Week 8 (May 18–21)     — os/filepath, CLI flags, memory/GC, idiomatic Go
- [x] Week 9 (May 22–28)     — iota/enums, rate limiting, encoding (csv, base64, hex)

## Topic index

| Topic | File |
|---|---|
| Variables, constants, types | basics/variables-and-constants.go |
| Strings & runes | basics/strings.go |
| Packages & imports | basics/edu-packages.go |
| Control flow | basics/control-flow.go |
| Arrays & slices | basics/arrays-and-slices.go |
| Maps | basics/maps.go |
| Pointers | basics/pointers.go |
| io.Reader / io.Writer | basics/io-interfaces.go |
| Go modules | basics/modules.go |
| Memory & GC | basics/memory-and-performance.go |
| Idiomatic Go | basics/idiomatic-go.go |
| Functions & types | functions/functions-and-types.go |
| Closures & variadic | functions/closures.go |
| Methods | functions/methods.go |
| init, main, named returns | functions/init-and-main.go |
| Structs | types/structs.go |
| Interfaces | types/interfaces.go |
| Interface patterns & DI | types/interfaces-advanced.go |
| Generics | types/generics.go |
| Type assertions & switches | types/type-assertions.go |
| Embedding | types/embedding.go |
| Functional options | types/functional-options.go |
| Error handling | error-handling/errors.go |
| Defer / panic / recover | error-handling/defer-panic-recover.go |
| Advanced error patterns | error-handling/advanced-errors.go |
| Goroutines & sync | concurrency/goroutines.go |
| Channels & select | concurrency/channels.go |
| Concurrency patterns | concurrency/patterns.go |
| Context | concurrency/context.go |
| Advanced sync | concurrency/advanced.go |
| File I/O | stdlib/file-io.go |
| JSON | stdlib/json.go |
| HTTP server & client | stdlib/http.go |
| Testing | stdlib/testing.go |
| strings, strconv, sort, math | stdlib/strings-strconv.go |
| time | stdlib/time.go |
| regexp | stdlib/regex.go |
| slices & maps packages | stdlib/slices-maps-pkg.go |
| os, filepath, exec | stdlib/os-and-env.go |
| slog | stdlib/slog.go |
| flag / CLI | stdlib/flags-and-cli.go |
| iota, enums, stringer | types/stringer-enum.go |
| Rate limiting | concurrency/rate-limiting.go |
| Encoding (csv, base64, hex) | stdlib/encoding.go |
