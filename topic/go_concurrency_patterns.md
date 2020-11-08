# The Go Blog: Go Concurrency Patterns: Context

Ref: https://blog.golang.org/context

- [Original Article Summary or Guidance](#original-article-summary-or-guidance)

## Original Article Summary or Guidance

### Introduction

In Go  servers, each incoming request is handled in its own goroutine. Request handlers often start additional goroutines
to access backends such as databases and RPC services.

The set of goroutines working on a request typically needs access to request-specific values such as the identity of the
end use, authorization tokens, and the request's deadline. When a request is canceled or times out, all the goroutines 
working on that request should exit quickly, so the system can reclaim any resources they are using.

Use [context](https://golang.org/pkg/context) to pass request-scoped values, cancellation signals, and deadlines across API boundaries to all the goroutines
involved in handling a request.

### Context

The core of the `context` package is the Context type:

```go
package context
import "time"

// A Context carries a deadline, cancellation signal, and request-scoped values
// across API boundaries. Its methods are safe for simultaneous use by multiple
// goroutines.
type Context interface {
    // Done returns a channel that is closed when this Context is canceled
    // or times out.
    Done() <-chan struct{}

    // Err indicates why this context was canceled, after the Done channel
    // is closed.
    Err() error

    // Deadline returns the time when this Context will be canceled, if any.
    Deadline() (deadline time.Time, ok bool)

    // Value returns the value associated with key or nil if none.
    Value(key interface{}) interface{}
}
```

More documentation see godoc...

- The `Done` method returns a channel that acts as a cancellation signal to functions running on behalf of the `Context`:
when the channel is closed, the functions should abandon their work and return.
- The `Err` method returns an error indicating why the `Context` is canceled. The [Pipelines and Cancellation](https://blog.golang.org/pipelines)
Article discusses the Done channel idiom in more detail.

The `WithCancel` function (described below) provides a way to cancel a new `Context` value.

A `Context` is safe for simultaneous use by multiple goroutines. Code can pass a single `Context` to any number of goroutines
and cancel that `Context` to signal all of them.

- The `Deadline` method allows function to determine whether they should work at all; if too little time if left, it may
not be worthwhile. Code may also use a deadline to set timeout for I/O operations.
- `Value` allows a `Context` to carry request-scoped data. That data must be safe for simultaneous use by multiple goroutines.

### Derived contexts

The context package provides functions to derive new Context values from existing ones. These values from a tree: when a
`Context` is canceled, all `Context` derived from it are also canceled.

Background is the root of any Context tree; it is never canceled:

```go
package context

// Context for illustration
type Context struct{}

// Background returns an empty Context. It is never canceled, has no deadline,
// and has no values. Background is typically used in main, init, and tests,
// and as the top-level Context for incoming requests.
func Background() Context
```

example explain omit...
