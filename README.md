# Hello Go

A small Go starter project for GitHub Codespaces. The example in [`main.go`](./main.go) starts two goroutines and waits for both to finish.

## Run in Codespaces

Open `main.go`. The Go extension adds a **Run/Debug** play button in the editor's upper-right corner. A workspace launch configuration is provided in [`.vscode/launch.json`](./.vscode/launch.json); you can also select **Launch Go program** from Run and Debug. The first Codespaces setup uses the Go devcontainer configuration in [`.devcontainer/devcontainer.json`](./.devcontainer/devcontainer.json), which installs Go and the Go extension. If prompted to rebuild the container after opening the project, do so.

You can also run the program in the integrated terminal:

```sh
go run .
```

Or build and run an executable:

```sh
go build -o hello-go .
./hello-go
```

The two greeting lines may appear in either order because each runs concurrently.

## Goroutines

A **goroutine** is a lightweight concurrent function, started with the `go` keyword. In this example, `main` starts two `greet` goroutines. A `sync.WaitGroup` keeps the program running until both have finished; without the wait, `main` could exit before they complete.

## Go references

- [A Tour of Go](https://go.dev/tour/) — interactive introduction to Go.
- [Effective Go](https://go.dev/doc/effective_go) — Go style and idioms.
- [Go by Example](https://gobyexample.com/) — concise, runnable examples, including [goroutines](https://gobyexample.com/goroutines) and [channels](https://gobyexample.com/channels).
- [Go documentation on concurrency](https://go.dev/doc/effective_go#concurrency) — goroutines and channels in the Go docs.
