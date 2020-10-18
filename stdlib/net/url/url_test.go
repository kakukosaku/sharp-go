package url_test

import (
	"fmt"
	"net/url"
)

func Example() {
	v := url.Values{"name": {"kaku", "gjs"}}
	v.Add("age", "18")
	fmt.Println(v)
	fmt.Println(v.Encode())
	fmt.Println(v.Get("age"))
	fmt.Printf("%q", v.Get("fake"))
	// Output:
	// map[age:[18] name:[kaku gjs]]
	// age=18&name=kaku&name=gjs
	// 18
	// ""
}
