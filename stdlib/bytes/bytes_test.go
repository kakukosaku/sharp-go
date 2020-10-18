// bytes api most similar to strings
package bytes_test

import (
	"bytes"
	"fmt"
)

func ExampleBuilder() {
	var b bytes.Buffer
	for i := 0; i < 3; i++ {
		_, err := fmt.Fprintf(&b, "%d...", i)
		if err != nil {
			panic(err)
		}
	}
	s := []byte{'k', 'a', 'k', 'u'}
	for _, v := range s {
		b.WriteByte(v)
	}
	fmt.Println(b.String())
	// Output:
	// 0...1...2...kaku
}
