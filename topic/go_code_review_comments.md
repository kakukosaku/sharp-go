# Go Code Review Comments

Ref: https://github.com/golang/go/wiki/CodeReviewComments#import-dot

This page collects common comments made during reviews of Go code, 
so that a single detailed explanation can be referred to by shorthands. 
This is a laundry list of common mistakes, not a comprehensive style guide.

- Gofmt

use `gofmt` or `goimports`

- Comment Sentences

See https://golang.org/doc/effective_go.html#commentary. 注解应该为完整的句子, 即使有些啰嗦显得. 注释应该以类型/函数的名称起头.

```go
package demo

import "io"

// Request represents a request to run a command.
type Request struct {}

// Encode writes the JSON encoding of req to w.
func Encode(w io.Writer, req *Request) {}
```

- context

使用`context.Context`type作为传递security credentials, tracing information, deadlines, and cancellation signals across 
API and process boundaries. 在go中, 一般显示使用 ctx, 作为整个function call chain中函数的第一个参数. 示例:

```go
package main
import "context"

func F(ctx context.Context, /* other arguments */) {}
```

- Copying

在copy 一个对象时, 需要注意是否"deep copy", 如copy a `bytes.Buffer`, 其内的 []byte 指向同一个底层数组, 会导致一些异想不到bug!

一般而言, 如果某类型`T`, 有绑定在`*T`上的方法, do not copy a value of `T`.

- crypto rand

在生成随机串时, 不要在不加随机量`rand.Seed(seed int64)`的情况下使用`math/rand`Read; 使用`crypto/rand`中Read.

- Declaring Empty Slices

2种方式: `var s []int` 与 `s := []int{}`, 前者为`nil`, 后者非`nil`. 在JSON序列化时, 前者为`null`, 后者为 `[]`.

- Doc Comments

all top-level, exported names should have doc comments, as should non-trivial unexported type or function declarations.

- Don't panic

不要使用panic做为异常处理, 使用`error`, 多个返回值还处理error.

- Error Strings

error strings 不应该有大写, 或符号; 日志不受该约定约束.

- Examples

添加新package时, 应该同时添加runnable example, or a simple test demonstrating a complete call sequence.

- Goroutine Lifetimes

When you spawn goroutines, make it clear when - or whether - they exit.

goroutine在等待channel sends or receives时, 可能导致goroutine泄漏, gc无法结束该goroutine.

Try to keep concurrent code simple enough that goroutine lifetimes are obvious.

- Handle Errors

一般来说, 不要用`_`忽略函数返回的error, handle the error, return it, or, in truly exceptional situation panic.

- Imports & Import Blank

- In-Band Errors & Indent Error Flow

```go
value, ok := Lookup(key)
if !ok {
	return fmt.Errorf("no value for %q", key)
}
return Parse(value)
```

- Initialisms

use ID instead of Id.

- Interface 

- Line Length

- Mixed Caps

- Named Result Parameters

- Naked Return

- Package comments

- Package Name

- Pass Value

Don't pass pointers as function arguments just to save a few bytes.

- Receiver Names

c: client, s: service

- Receiver Type

avoid copy use pointer; mutate the receiver use pointer; large struct or array use pointer is more efficient;
Finally, when in doubt, use a pointer receiver.

- Synchronous Functions

- Useful Test Failures

- Variable Names

Variable names in Go should be short rather than long. This is especially true for local variables with limited scope.
Prefer c to lineCount. Prefer i to sliceIndex.
The basic rule: the further from its declaration that a name is used, the more descriptive the name must be.
