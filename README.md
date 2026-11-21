# Hello Go

A small Go starter project for GitHub Codespaces. The example in [`main.go`](./main.go) starts two goroutines and waits for both to finish.

## Run in Codespaces

Open `main.go` and click **Run Code** (the play icon) in the editor title bar. Code Runner runs `go run .` in the integrated terminal. The devcontainer installs Code Runner and the Go extension via [`.devcontainer/devcontainer.json`](./.devcontainer/devcontainer.json). If the container is already open, rebuild it to install Code Runner. You can also run `go run .` in the terminal, or select **Launch Go program** in Run and Debug and press F5 to debug; its configuration is in [`.vscode/launch.json`](./.vscode/launch.json).

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
