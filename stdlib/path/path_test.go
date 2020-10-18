package path_test

import (
	"fmt"
	"path"
)

func Example() {
	fmt.Println(path.Base("/home/kaku/t.go"))
	fmt.Println(path.Clean("/home/kaku/dir/.."))
	fmt.Println(path.Dir("/home/kaku/t.go"))
	fmt.Println(path.Ext("/home/kaku/t.go"))
	fmt.Println(path.Join("/", "home", "/kaku"))
	fmt.Println(path.Match("/home/k*", "/home/kaku"))
	fmt.Println(path.Split("/home/kaku/t.go"))
	// Output:
	// t.go
	// /home/kaku
	// /home/kaku
	// .go
	// /home/kaku
	// true <nil>
	// /home/kaku/ t.go
}
