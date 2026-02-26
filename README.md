# URL Shortener (Go starter)

If you're new to Go, this repository now includes a minimal **Hello World** setup using Go modules.

## 1) Install Go

Download and install Go from the official website:

- https://go.dev/dl/

Then verify it is installed:

```bash
go version
```

## 2) Project structure

This is the smallest runnable Go project:

- `go.mod` → defines your module name and Go version.
- `main.go` → your entrypoint program.

## 3) Module and package basics (quick intro)

- **Module**: your project (managed by `go.mod`).
- **Package**: a folder of `.go` files. Every file in a folder uses the same `package` name.
- `package main` + `func main()` means executable app.
- `import "fmt"` lets you use the standard formatting package.

## 4) Run Hello World

```bash
go run .
```

Expected output:

```text
Hello, World!
```

## 5) Build a binary

```bash
go build -o app .
./app
```

## 6) Add dependencies later

When you import an external package, Go can add it automatically with:

```bash
go mod tidy
```

You can inspect dependencies with:

```bash
go list -m all
```

---

If you want, I can help you next with a beginner-friendly folder layout for a real URL shortener API (`cmd/`, `internal/`, routing, handlers, and storage).
