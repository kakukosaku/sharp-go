package strconv_test

import (
	"fmt"
	"strconv"
)

func ExampleAppend() {
	fmt.Println(string(strconv.AppendBool([]byte("bool:"), true)))
	fmt.Println(string(strconv.AppendFloat([]byte("float32:"), 3.1415926, 'E', -1, 32)))
	fmt.Println(string(strconv.AppendInt([]byte("int (base 10):"), -42, 10)))
	fmt.Println(string(strconv.AppendQuote([]byte("quote:"), "string")))
	// Output:
	// bool:true
	// float32:3.1415925E+00
	// int (base 10):-42
	// quote:"string"
}

func ExampleAtoiAndItoa() {
	i, err := strconv.Atoi("18")
	if err != nil {
		panic(err)
	}
	s := strconv.Itoa(18)
	fmt.Printf("%d, %[1]T\n", i)
	fmt.Printf("%s, %[1]T\n", s)
	// Output:
	// 18, int
	// 18, string
}

func ExampleFormat() {
	fmt.Printf("%s, %[1]T\n", strconv.FormatBool(true))
	fmt.Printf("%s, %[1]T\n", strconv.FormatInt(18, 10))
	fmt.Printf("%s, %[1]T\n", strconv.FormatFloat(3.14159265, 'f', -1, 32))
	fmt.Printf("%s, %[1]T\n", strconv.FormatFloat(3.14159265, 'g', -1, 64))
	// Output:
	// true, string
	// 18, string
	// 3.1415927, string
	// 3.14159265, string
}

func ExampleParseBool() {
	b, err := strconv.ParseBool("true")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v, %[1]T\n", b)

	i, err := strconv.ParseInt("18", 10, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v, %[1]T\n", i)
	// Output:
	// true, bool
	// 18, int64
}

func ExampleQuote() {
	fmt.Println(strconv.Quote("kaku"))
	fmt.Println(strconv.QuoteRune('中'))
	fmt.Println(strconv.QuoteRuneToASCII('中'))
	fmt.Println(strconv.Unquote(strconv.Quote("kaku")))
	// Output:
	// "kaku"
	// '中'
	// '\u4e2d'
	// kaku <nil>
}
