package io_test

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ExampleCopy() {
	r1 := strings.NewReader("kaku")
	if _, err := io.Copy(os.Stdout, r1); err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	r2 := strings.NewReader("gjs")
	buf := make([]byte, 8)
	// buf is used here...
	r1.Seek(0, 0)
	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	// ... reused here also, No need to allocate an extra buffer.
	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	r1.Seek(0, 0)
	if _, err := io.CopyN(os.Stdout, r1, 2); err != nil {
		log.Fatal(err)
	}
	// Output:
	// kaku
	// kaku
	// gjs
	// ka
}

func ExamplePipe() {
	r, w := io.Pipe()
	go func() {
		_, _ = fmt.Fprintln(w, "some text to be read")
		_ = w.Close()
	}()

	buf := bytes.Buffer{}
	_, _ = buf.ReadFrom(r)
	fmt.Println(buf.String())
	// Output:
	// some text to be read
}

func ExampleRead() {
	r := strings.NewReader("some text to be read")
	buf := make([]byte, 10)

	if _, err := io.ReadAtLeast(r, buf, 4); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)

	// Output:
	// some text
}

func ExampleWrite() {
	_, _ = io.WriteString(os.Stdout, "kaku")
	// Output:
	// kaku
}
