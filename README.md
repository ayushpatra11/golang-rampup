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
                  embedding, functional options
error-handling/   errors, defer, panic, recover, advanced patterns
concurrency/      goroutines, channels, select, sync primitives,
                  context, patterns, advanced patterns
stdlib/           file I/O, JSON, HTTP, testing, strings/strconv,
                  time, regex, slices/maps packages, os, slog, flags
```

## Progress

- [x] Week 1 (Mar 28–Apr 3) — basics: variables, types, control flow, arrays, slices, maps, pointers
- [x] Week 2 (Apr 4–10) — functions: closures, methods, generics, type assertions, embedding
- [x] Week 3 (Apr 11–17) — error handling, goroutines, channels, concurrency patterns
- [x] Week 4 (Apr 18–24) — context, file I/O, JSON, HTTP, testing, strings
- [x] Week 5 (Apr 25–May 1) — io interfaces, init/main, regexp, modules
- [x] Week 6 (May 2–10) — slices/maps packages, advanced concurrency, os/filepath, interfaces
- [x] Week 7 (May 11–17) — functional options, advanced errors, slog
- [x] Week 8 (May 18–21) — memory/GC, CLI flags, idiomatic Go patterns

## Key concepts covered

| Topic | File |
|---|---|
| Variables, constants, types | basics/variables-and-constants.go |
| Strings & runes | basics/strings.go |
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
