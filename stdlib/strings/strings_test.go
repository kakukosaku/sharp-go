package strings_test

import (
	"fmt"
	"strings"
	"unicode"
)

func ExampleCompare() {
	fmt.Println(strings.Compare("a", "b"))
	// use the built-in string comparison operators ==, <, > and so on is faster than this func.
	fmt.Println("a" == "a")
	// Output:
	// -1
	// true
}

func ExampleContains() {
	fmt.Println(strings.Contains("kakukosaku", "kaku"))
	// Output:
	// true
}

func ExampleContainsAny() {
	fmt.Println(strings.ContainsAny("kaku", "iu"))
	fmt.Println(strings.ContainsAny("kaku", ""))
	fmt.Println(strings.ContainsAny("", ""))
	// Output:
	// true
	// false
	// false
}

func ExampleContainsRune() {
	fmt.Println(strings.ContainsRune("kaku", 'k'))
	fmt.Println(strings.ContainsRune("kaku", 97))
	// Output:
	// true
	// true
}

func ExampleCount() {
	fmt.Println(strings.Count("kaku", "k"))
	// before & after each rune
	fmt.Println(strings.Count("kaku", ""))
	// Output:
	// 2
	// 5
}

func ExampleEqualFold() {
	fmt.Println(strings.EqualFold("kaku", "KAKU"))
	fmt.Println(strings.EqualFold("kaku", "kuka"))
	// Output:
	// true
	// false
}

func ExampleFields() {
	fmt.Println(strings.Fields(" kaku  gjs kakukosaku  "))
	// Output:
	// [kaku gjs kakukosaku]
}

func ExamplePrefixSuffix() {
	fmt.Println(strings.HasPrefix("kaku", "ka"))
	fmt.Println(strings.HasSuffix("kaku", "ku"))
	// Output:
	// true
	// true
}

func ExampleIndex() {
	fmt.Println(strings.Index("kaku", "ka"))
	fmt.Println(strings.Index("kaku", "unk"))
	fmt.Println(strings.IndexAny("kaku", "unk"))

	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("hello, 世界", f))
	// Output:
	// 0
	// -1
	// 0
	// 7
}

func ExampleJoin() {
	fmt.Println(strings.Join([]string{"kaku", "ko", "saku"}, " "))
	// Output:
	// kaku ko saku
}

func ExampleMap() {
	fmt.Println(strings.Map(func(c rune) rune {
		if unicode.IsLetter(c) {
			return unicode.ToUpper(c)
		} else {
			return ' '
		}
	}, "kaku.ko.saku"))
	// Output:
	// KAKU KO SAKU
}

func ExampleRepeat() {
	fmt.Println(strings.Repeat("k", 2))
	// Output:
	// kk
}

func ExampleReplace() {
	fmt.Println(strings.Replace("kaku", "", " ", -1))
	// Output:
	//  k a k u
}

func ExampleSplit() {
	fmt.Println(strings.Split("a,b,c", ","))
	fmt.Println(strings.Split("abc", ""))
	fmt.Println(strings.Split(" abc", ""))
	fmt.Println(strings.SplitAfter("a,b,c", ","))
	// Output:
	// [a b c]
	// [a b c]
	// [  a b c]
	// [a, b, c]
}

func ExampleToLower() {
	s := "KAKU"
	fmt.Println(strings.ToLower(s))
	// Output:
	// kaku

}

func ExampleTrim() {
	fmt.Println(strings.Trim("-=kaku=-", "-="))
	fmt.Println(strings.TrimSpace(" \r\nkaku\t"))
	// Output:
	// kaku
	// kaku
}

func ExampleBuilder() {
	b := strings.Builder{}
	for i := 0; i < 3; i++ {
		_, err := fmt.Fprintf(&b, "%d...", i)
		if err != nil {
			panic(err)
		}
	}
	b.WriteString("kaku")
	fmt.Println(b.String())
	b.Reset()
	// Output:
	// 0...1...2...kaku
}

func ExampleReader() {
	fmt.Println(strings.NewReader("kaku").ReadRune())
	// Output:
	// 107 1 <nil>
}

func ExampleReplacer() {
	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	fmt.Println(r.Replace("This is <b>HTML</b>!"))
	// Output:
	// This is &lt;b&gt;HTML&lt;/b&gt;!
}