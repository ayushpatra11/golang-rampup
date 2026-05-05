package basics

// ---- Go Modules ----
// Module = collection of related Go packages with a version
// go.mod defines the module path and its dependencies

/*
go.mod anatomy:
  module github.com/username/myproject  <- module path (used in import paths)
  go 1.22                               <- minimum Go version

  require (
      github.com/some/dep v1.2.3
      github.com/other/dep v0.4.1
  )

  replace github.com/old/dep => ./local/dep  <- local override
  exclude github.com/bad/dep v1.0.0          <- exclude broken version
*/

// ---- Common go commands ----

// go mod init github.com/user/module   initialize a new module
// go mod tidy                          add missing, remove unused dependencies
// go mod download                      download deps to local cache
// go mod vendor                        copy deps to ./vendor
// go mod verify                        verify checksums match go.sum

// go get github.com/pkg@v1.2.3         add/upgrade a dependency
// go get github.com/pkg@latest         get latest version
// go get github.com/pkg@none           remove a dependency

// go list -m all                       list all modules in the build
// go list -m -versions github.com/pkg  list available versions

// ---- Versioning ----
// Semantic versioning: vMAJOR.MINOR.PATCH
// v1.2.3  - stable release
// v0.x.y  - unstable/pre-1.0
// v2.x.y+ - major version MUST change the module path:
//           module github.com/user/pkg/v2

// ---- Workspaces (Go 1.18+) ----
// go.work: develop multiple modules together without editing go.mod

/*
go.work:
  go 1.22
  use (
      ./module-a
      ./module-b
  )
*/

// go work init ./module-a ./module-b   create workspace
// go work use ./new-module             add module to workspace
// go work sync                         sync workspace deps

// ---- Private modules ----
// GONOSUMCHECK, GONOSUMDB: skip checksum for private repos
// GOFLAGS, GOENV, GOPRIVATE, GONOSUMCHECK environment vars

// ---- Build tags ----
// Restrict which files are included in a build

// //go:build linux && amd64     <- Go 1.17+ syntax (top of file, before package)
// // +build linux,amd64         <- old syntax (still supported)

// go build -tags integration     build with a specific tag
// go test -tags integration ./...

// ---- Useful tools ----
// go vet ./...          static analysis
// go fmt ./...          format code (or: gofmt -w .)
// go doc pkg            show documentation
// godoc -http=:6060     local doc server
// gopls                 language server (for IDE integration)
// golangci-lint         popular linter aggregator

// ---- Environment variables ----
// GOPATH   workspace root (legacy, less relevant with modules)
// GOROOT   Go installation directory
// GOOS     target OS for cross-compilation (linux, darwin, windows)
// GOARCH   target architecture (amd64, arm64, 386)
// CGO_ENABLED  enable/disable cgo (0 for pure Go static binary)

// Cross-compilation example:
// GOOS=linux GOARCH=amd64 go build -o myapp-linux ./cmd/myapp
