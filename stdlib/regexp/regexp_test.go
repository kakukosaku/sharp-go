package regexp

import (
	"fmt"
	"regexp"
)

func Example() {
	r1 := `([a-z]+\[[0-9]+\])`
	var validID = regexp.MustCompile(r1)

	fmt.Println(validID.MatchString("kaku[1]"))
	fmt.Println(regexp.MatchString(r1, "kaku[]"))
	fmt.Println(validID.ReplaceAllString("kaku[1]", "-${1}-"))
	fmt.Println(validID.Split("1kaku[1]2gjs[2]3", -1))
	// Output:
	// true
	// false <nil>
	// -kaku[1]-
	// [1 2 3]
}
