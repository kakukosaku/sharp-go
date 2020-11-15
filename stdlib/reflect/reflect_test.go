// Author: kaku
// Date: 2020/11/15
//
// GitHub:
//	https://github.com/kakukosaku
//
package reflect_test

import (
	"fmt"
	"reflect"
)

func ExampleTypeof() {
	var i interface{}
	name := "kaku"
	age := 18

	i = name
	// i = age
	i = &age

	// i's type
	fmt.Println(reflect.TypeOf(age).Name())

	// i is *int pointer
	v := reflect.ValueOf(i)
	fmt.Println(v.CanSet())
	fmt.Println(v.Elem().CanSet())
	// change i's value
	v.Elem().SetInt(19)
	fmt.Println(*(v.Interface().(*int)))
	fmt.Println(age)
	// Output:
	// int
	// false
	// true
	// 19
	// 19
}
