// author: kaku
// date: 2020/月/日
//
// GitHub:
//	https://github.com/kakukosaku
//
// Description:
//
//	"The function and arguments expressions are evaluated when the statement is executed,
//	but the actual call is deferred until that contains the defer statement has finished,
//	whenever normally, by executing a return statement or falling off the end, or abnormally, by panicking."
//		-- The Go Programming Language
//
//	multi-defer implement by stack, so defers is executed FIFO
package main

import "fmt"

func main() {
	fmt.Println(test1())
	fmt.Println(test2())
	fmt.Println(test3())
	fmt.Println(test4())
}

func test1() int {
	var i = 1
	defer func() {
		i++
	}()
	return 1
}

func test2() int {
	var i = 1
	defer func() {
		if p := recover(); p != nil {
			// no return
			fmt.Println("Hi panic")
		}
	}()

	panic("Hi")
	return i
}

func test3() (i int) {
	i = 1
	defer func() {
		if p := recover(); p != nil {
			// no return
			fmt.Println("Hi panic")
		}
		// i++
	}()

	panic("Hi")
	return i
}

func test4() int {
	var i int
	defer func() {
		fmt.Println("defer 1")
	}()

	defer func() {
		fmt.Println("defer 2")
	}()

	return i
}
