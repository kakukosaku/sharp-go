# sharp-go

This repository's goal is to help you familiar/master with go.

## Packages

- [grammar](#grammar)
- [pitfall](#pitfall)
- [stdlib](#stdlib)
- [project](#project)
- [example](#example)
- [topic](#topic)
- [summary](#summary)

### grammar

- [Data Types](grammar/data_type_test.go)
- [Control Flow](grammar/control_flow_test.go)
- [Goroutine](grammar/goroutine_tour_test.go)

### pitfall

pitfall

### stdlib example codes

- [bufio](stdlib/bufio) implements buffered I/O. Exported function: `bufio.NewReader`, `bufio.NewScanner` etc.
- [bytes](stdlib/bytes) implements function for the manipulation of byte slices. Exported Function: `bytes.Contains`, `bytes.Equal`; Struct: `bytes.Buffer` etc.
- [container/heap](stdlib/container/heap), [contain/list](stdlib/container/list), [contain/ring](stdlib/container/ring): common Interface implementation.
- [context](stdlib/context) define context.Context used in invoke chain or sub-goroutines.
- [database/sql](stdlib/database/sql) database/sql provides a generic interface around SQL (or SQL-like) databases, you need specific driver (import by _ to registered).
- [encoding/json](stdlib/encoding/json) json, csv, xml Marshal, Unmarshal etc.
- [errors](stdlib/errors) implements functions to manipulate errors, such as `errors.New`, `errors.Is`, `errors.As` etc.
- [io](stdlib/io) provides basic interfaces to I/O primitives.
- [io/util](stdlib/io/ioutil) implements some I/O utility functions, such as `ioutil.ReadAll`, `ReadDir, ReadFile, WriteFile` etc.
- [math](stdlib/main) etc.
- [net/http](stdlib/net/http) [net/url](stdlib/net/url) so import for back-end developer!
- [path](stdlib/path) implements utility routines for manipulating slash-separated, such as `Split, Join, Dir`.
- [regexp](stdlib/regexp) RE functions.
- [strconv](stdlib/strconv) types -> string or verse, such as `strconv.ParseInt strconv.FormtInt64`.
- [strings](stdlib/strings) implements simple functions to manipulate UTF-8 encoded strings.
- [time](stdlib/time) time functions, such as `time.Now, time.NewTicker time.After` etc.

### project

project

### example

example

### topic

- [The Go Blog - Godoc: documenting Go code](topic/documenting_go_code.md)
- [Effective Go](topic/effective_go.md)
- [Go Code Review Comments](topic/go_code_review_comments.md)
- [The Go Blog: Go Concurrency Patterns: Context](topic/go_concurrency_patterns.md)
- [The Go Blog: Getting to Go: The Journey of Go's Garbage Collector](topic/go_garbage_collector.md)
- [The Go Blog: Go GC: Prioritizing low latency and simplicity](topic/go_gc_low_latency_and_simplicity.md)
- [The go Memory Model](topic/go_memory_model.md)
- [The Go Blog: Organizing Go Code](topic/organizing_go_code.md)

Just Links:

- [How to Write Go Code](https://golang.org/doc/code.html)
- [The Go Blog: Strings, bytes, runes and characters in Go](https://blog.golang.org/strings)
- [The Go Blog: Defer, Panic and Recover](https://blog.golang.org/defer-panic-and-recover)
- [The Go Blog: Error handling and Go](https://blog.golang.org/error-handling-and-go)
- [The Go Blog: Working with Errors in GO 1.13](https://blog.golang.org/go1.13-errors)
- [The Go Blog: HTTP/2 Server Push](https://blog.golang.org/h2push)
- [The Go Blog: JSON and Go](https://blog.golang.org/json)
- [Channels in Go](https://go101.org/article/channel.html)
- [How to Gracefully Close Channels](https://go101.org/article/channel-closing.html)

Some Video introduced by Go documentation~

- [Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs)
- [Advanced Go Concurrency Patterns](https://www.youtube.com/watch?v=QDDwwePbDtw)
