# [The Go Blog - Godoc: documenting Go code](https://blog.golang.org/godoc-documenting-go-code)

> The convention is simple: to document a type, variable, constant, function, or even a package, write a regular comment directly preceding its declaration, with no intervening blank line. Godoc will then present that comment as text alongside the item it documents. For example, this is the documentation for the fmt package's Fprint function:

> // Fprint formats using the default formats for its operands and writes to w.  
> // Spaces are added between operands when neither is a string.  
> // It returns the number of bytes written and any write error encountered.  
> func Fprint(w io.Writer, a ...interface{}) (n int, err error) {  

rules:

- Comments is a sentence that begin with the name of the element it describes.

- Comments on package declarations should general package documentation. eg:

> // Package sort provides primitives for sorting slices and user-defined  
> // collections.  
> package sort  