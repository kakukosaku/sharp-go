# The Go Blog: Organizing Go Code

- [Original Article Summary or Guidance](#original-article-summary-or-guidance)

## Original Article Summary or Guidance

### Introduction

Go code is organized differently to that of other language. This post discusses how to name and package the elements of 
your Go programs to best serve its uses.

### Choose good names

**The names you choose affect how you this about your code, so take care when naming you package and its exported identifiers.**

A package's name provides context for tis contents. For instance, the `bytes` package from the standard library exports
the `Buffer` type. On its own, the name `Buffer` isn't very descriptive, but when combined with its package name its meaning
becomes clear: `bytes.Buffer`.  If the package had a less descriptive name, like util, the buffer would likly acquire the 
longer and clumsier `util.BytesBuffer`.

Don't be shy about renaming thing as you work. As you spend time with your program you will better understand how its pieces
fit together and, therefore, what their names should be. There's no need to lock yourself into early decisions.

A good name is the most important part of a software interface: the name is the first thing every client of the code will
see. A well-chosen name is therefore the starting point for a good documentation. Many of the following practices result
organically from good naming.

### Choose a good import path (make you package "go get"-able)

title is so clean, content omit...

### Minimize the exported interface

**If in doubt, leave it out!**

title is so clean, content omit...

### What to put into a package

title is so clean, content omit...

### Document your code

title is so clean, content omit...
