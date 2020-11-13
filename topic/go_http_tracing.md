# The Go Blog: Introducing HTTP Tracing

- [Original Article Summary or Guidance](#original-article-summary-or-guidance)

## Original Article Summary or Guidance

### Introduction

Support for HTTP tracing is provided by the [net/http/httptrace](https://golang.org/pkg/net/http/httptrace/) package.
The collected information can be used for debugging latency issues, service monitoring, writing adaptive systems and more.

### HTTP events

The httptrace package provides a number of hooks to gater information during an HTTP roud trip about a variety of events.
These events include: 

- Connection creation
- Connection reuse
- DNS lookups
- Writing the request to the wire
- Reading the response

### Tracing events

You can enable HTTP tracing by putting an `*httptrace.ClientTrace` containing hook functions into a request's `context.Context`.
Various `http.RoundTripper` implementations report the internal events by looking for context's `*httptrace.ClientTrace`
and calling the relevant hook functions.

[reference code](../stdlib/net/http/httptrace/httptrace_test.go)

### Tracing with http.Client

The tracing mechanism is designed to trace the events in the lifecycle of a single http.Transport.RoundTrip. However, 
a client may make multiple round trips to complete an HTTP request. For example, in the case of a URL redirection, the 
registered hooks will be called as many times as the client follows HTTP redirects, making multiple requests. Users are 
responsible for recognizing such events at the http.Client level. The program below identifies the current request by 
using an http.RoundTripper wrapper.

The Transport in the net/http package supports tracing of both HTTP/1 and HTTP/2 requests.

example libs: [httpstat](https://github.com/davecheney/httpstat)
