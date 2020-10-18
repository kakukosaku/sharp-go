// author: kaku
// date: 2020/月/日
//
// GitHub:
//	https://github.com/kakukosaku
//
// Description:
//	loop 中初始化的变量, 在loop期间为"同一个", 意味着共享内存地址, 注意闭包函数使用loop init 中的变量
//
package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	var pFunc []func()
	arr := [5]int{1, 2, 3, 4, 5}

	for _, i := range arr {
		fmt.Printf("%#v\n", &i)
		pFunc = append(pFunc, func() { fmt.Println(i) })
	}

	for _, i := range arr {
		t := i
		pFunc = append(pFunc, func() { fmt.Println(t) })
	}

	for _, f := range pFunc {
		f()
	}
}
