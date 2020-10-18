// Author: kaku
// Date: 2020/10/18
//
// GitHub:
//	https://github.com/kakukosaku
//
package grammar_test

import (
	"strings"
	"testing"
	"time"
)

func TestControlFlow(t *testing.T) {
	// for each
	s := []string{"a", "b", "c"}
	for idx, ele := range s {
		t.Log(idx, ele)
	}

	m := map[string]int{
		"kaku":         18,
		"kaku.ko.saku": 27,
	}
	for k, v := range m {
		t.Log(k, v)
	}

	// for loop
	for i := 0; i < 3; i++ {
		t.Log(i)
	}

	// for..ever :) loop
	//for {
	//	t.Logf(".")
	//}

	choice := "A"
	switch choice {
	case "A", "B", "C":
		t.Log("child choice")
	case "D", "E", "F":
		t.Log("my love")
	}

	// switch to true express
	choice = "D"
	switch {
	case choice == "A":
		t.Log("child choice")
	case strings.Contains("DEF", choice):
		t.Log("my love")
	}

	// select used for channel
	select {
	case <- time.After( time.Duration(3) * time.Second):
		t.Log("after three seconds")
	}
}
