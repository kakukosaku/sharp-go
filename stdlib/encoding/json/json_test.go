package json_test

import (
	"encoding/json"
	"fmt"
	"log"
)

func ExampleMarshal() {
	var s = struct {
		Name string `json:"name,"`
		Age  int    `json:"-,"`
	}{
		Name: "kaku",
		Age:  18,
	}
	c := `
	{
		"name": "kaku",
		"-": 19
	}
	`
	m := map[string][]byte{"<kaku>": {'a', 'b'}, "gjs": {'1', '2'}}
	r, _ := json.Marshal(m)
	fmt.Println(string(r))
	r2, _ := json.MarshalIndent(s, "", "\t")
	fmt.Println(string(r2))
	err := json.Unmarshal([]byte(c), &s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s.Age)
	// Output:
	// {"\u003ckaku\u003e":"YWI=","gjs":"MTI="}
	// {
	// 	"name": "kaku",
	// 	"-": 18
	// }
	// 19
}
