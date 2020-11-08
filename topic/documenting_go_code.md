# The Go Blog - Godoc: documenting Go code
 
Ref: https://blog.golang.org/godoc-documenting-go-code

- [Original Article Summary or Guidance](#original-article-summary-or-guidance)
- [Introduction to Commend godoc](#introduction-to-commend-godoc)

## Original Article Summary or Guidance

Godoc are not language constructs(as Python's [Docstring](https://www.python.org/dev/peps/pep-0257/), Its part of 
function or class) nor must they have their own machine-readable syntax(as Java's [Javadoc](https://www.oracle.com/java/technologies/javase/javadoc-tool.html)).
Godoc comments are just good comments, the sort you would want to read even if godoc didn't exist.

> Godoc are just plain&good comments.

The convention is simple: to document a type, variable, constant, function, or even a package, write a regular comment 
directly preceding its declaration, with no intervening blank line. Godoc will then present that comment as text alongside
the item it documents. For example, this is the documentation for the fmt package's [Fprint](https://golang.org/pkg/fmt/#Fprint)

```go
package fmt
import "io"

// Fprint formats using the default formats for its operands and writes to w.  
// Spaces are added between operands when neither is a string.  
// It returns the number of bytes written and any write error encountered.  
func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
    // for markdown language render sake, Add by kaku.
    return 0, nil
}
```

Notice this comment is a complete sentence that begins with the name of the element it describes.

Comments on the package declarations should provide general package documentation. These comments can be short, like the 
[sort](https://golang.org/pkg/sort/) package's brief description:

```go
// Package sort provides primitives for sorting slices and user-defined
// collections.
package sort
```

They can also be detailed like the [gob package](https://golang.org/pkg/encoding/gob/)'s overview. That package uses another
convention for packages that need large amounts of introductory documentation: the package comment is placed in its own 
file, **doc.go**, which contains only those comments and a package clause.

Sometimes a struct field, function, type, or even a whole package becomes redundant or unnecessary, but must be kept for 
compatibility with existing programs. To signal that an identifier should not be used, add a *paragraph to its doc comment*
*that begins with "Deprecated:"* followed by some information about the deprecation. example:

```go
package driver
// Execer is an optional interface that may be implemented by a Conn.
//
// If a Conn implements neither ExecerContext nor Execer,
// the sql package's DB.Exec will first prepare a query, execute the statement,
// and then close the statement.
//
// Exec may return ErrSkip.
//
// Deprecated: Drivers should implement ExecerContext instead.
type Execer interface {
    // for illustration problem, not complete code!
}
```

There are a few formatting rules that Godoc uses where converting comments to HTML:

- Subsequent lines of text are considered part of the same paragraph; you must leave a blank to separate paragraphs.
- Pre-formatted text must be indented relative to the surrounding comment text.
- URLs will be converted to HTML links; no special markup is necessary.

Note that none of these rules requires you to do anything out of the ordinary.

## Introduction to Commend godoc

`godoc` extracts and generates documentation for Go programs.

It runs as a web server and presents the documentation as a web page. 

```shell script
godoc -http=:6060
```

Usage:

```shell script
godoc [flage]
```

some important flag:

- -v: verbose mode
- -http=addr: HTTP service address (e.g. '127.0.0.1:6060' or just ':6060'), use ?m=all to see all declarations, not just the exported ones.

-m=methods, -m=src, -m=flat

don't have anything else important. :)
