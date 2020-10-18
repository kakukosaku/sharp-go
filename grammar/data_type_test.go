// Author: kaku
// Date: 2020/10/18
//
// GitHub:
//	https://github.com/kakukosaku
//
// How to run:
// 	cd {project}/grammar && go test -v .
//
package grammar_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestBasicDataType(t *testing.T) {
	// declare with zero value
	var i1 int
	assert.Equal(t, 0, i1)
	var s1 string
	assert.Equal(t, "", s1)

	// declare && define
	i2 := 1
	assert.Equal(t, 1, i2)

	// multiple var
	var r1 int32
	r2, s2 := 'a', "a中文"
	assert.IsType(t, r1, r2)
	assert.IsType(t, s1, s2)

	// iterate a string
	for i, r := range s2 {
		// Note:
		// 1. the i is duplicate init.
		// 2. the i is not sequential.
		t.Log(i, r, string(r))
	}

	// int is not alias int64, or int32!
	// type is different, so is not equal
	assert.NotEqual(t, int64(1), i2)

	// basic int operator
	// ++ operator is statement, not return value
	// so you can use i1++ as a function argument!
	// this line will compile error
	// assert.Equal(t, i1++, i2)
	i1++
	assert.Equal(t, i2, i1)

	// basic string operator
	s3 := "a"
	assert.Equal(t, "abc", s3+"bc")

	// go do not support operator overload!
	// so you can use like, "a" * 3 -> "aaa"
	// this line will compile error
	//s4 := "a" * 3

	// more string operation, if you familiar with c/c++, you should know useful function in stdlib.
	rest := strings.Split("a,b,c", ",")
	assert.Equal(t, rest, []string{"a", "b", "c"})
	// another code style
	assert.Equal(
		t,
		strings.Join([]string{"a", "b", "c"}, ","),
		"a,b,c",
	)
}

func TestConstant(t *testing.T) {
	// const with iota, iota auto increase
	const (
		a = iota      // iota begin from 0
		b             // b = iota = 1
		_             // _ = iota = 2
		c = iota + 10 // iota = 3, c = 3 + 10
	)
	assert.Equal(t, 0, a)
	assert.Equal(t, 1, b)
	assert.Equal(t, 13, c)

	const (
		d = iota + 1 // iota reset to 0
		e            // resolve to: iota + 1, iota = 1
	)
	assert.Equal(t, 1, d)
	assert.Equal(t, 2, e)
}

func TestAggregateDataType(t *testing.T) {
	// declare zero value
	// Note:
	// 1. array is different with slice
	// 2. zero value of array is different with zero value of slice
	var arr [3]int
	var sli1 []int
	assert.Equal(t, 0, arr[0])
	assert.Equal(t, []int(nil), sli1)

	// this func access s first element
	// this line will panic
	// fmt.Println(sli[0])
	f := func() {
		t.Log(sli1[0])
	}
	assert.Panics(t, f)

	// init slice by use make() function
	sli2 := make([]int, 3)
	assert.Equal(t, 3, len(sli2))
	assert.Equal(t, 0, sli2[0])

	// you can use slice as *dynamic* array, but its not array.
	sli3 := append(sli2, 1)
	assert.Equal(t, 1, sli3[3])

	// slice can not use "=" operator
	// if sli2 == sli3 {} will compile error
	// but array can
	if arr == [3]int{} {
		t.Log("array can use = operator to check equal or not")
	}

	// same as map
	var m1 map[string]int
	assert.Equal(t, map[string]int(nil), m1)

	// in-line
	m2 := map[string]int{
		"kaku": 18,
	}

	assert.Equal(t, 18, m2["kaku"])
	// access not exist key, return zero value!
	assert.Equal(t, 0, m2["kakukosaku"])

	// return value and exist or not
	if value, ok := m2["kaku"]; ok {
		assert.Equal(t, 18, value)
	}

	// channel & goroutine
	var ch1 chan int
	assert.Equal(t, chan int(nil), ch1)

	ch2 := make(chan int, 1)
	ch2 <- 1
	close(ch2)
	for val := range ch2 {
		assert.Equal(t, 1, val)
	}
}

func TestAbstractDataType(t *testing.T) {
	// custom data type
	type Age int
	var a1 Age = 18
	a2 := Age(27)
	assert.Equal(t, Age(18), a1)
	assert.NotEqual(t, 27, a2)

	// abstract data type
	type User struct {
		Names []string
		Age   int
	}

	// User zero value
	var u User
	assert.NotEqual(t, (*User)(nil), &u)
	assert.Equal(t, 0, u.Age)

	// nil pointer
	var uPtr1 *User
	if uPtr1 == nil {
		t.Log("always be true, declare a ptr by `var`")
	}
	assert.Equal(t, (*User)(nil), uPtr1)

	// new(Type) function return a pointer to a newly allocated zero value of that type.
	uPtr2 := new(User)
	if uPtr2 != nil {
		t.Log("always be true, init a ptr by `new() function`")
	}
	assert.Equal(t, 0, uPtr2.Age)
	assert.NotEqual(t, (*User)(nil), uPtr2) // different with `var ptr *User`
}

type User struct {
	Names []string
	Age   int
}

// SamePerson Judge whether u1, u2 is same person.
// return true if one of them name is same!
func (u *User) SamePerson(u2 *User) bool {
	for _, name1 := range u.Names {
		for _, name2 := range u2.Names {
			if name1 == name2 {
				return true
			}
		}
	}

	return false
}

func TestFuncMethod(t *testing.T) {
	// function just as this test function

	// method: bind to a struct or its pointer!
	u1 := User{Names: []string{"kaku", "kaku ko saku"}, Age: 18}
	u2 := User{Names: []string{"kaku", "kaku.ko.saku"}, Age: 18}

	// pointer method only access by pointer, struct method can be access both~
	assert.Equal(t, true, u1.SamePerson(&u2))
}

func (u User) String() string {
	names := strings.Join(u.Names, " or ")
	return fmt.Sprintf("You can call me %v, I'm %d years old!\n", names, u.Age)
}

// Study notice this method bind to *User!
func (u *User) Study() {
	fmt.Println("I'm a student of society college now!")
}

type Student interface {
	Study()
}

func TestInterface(t *testing.T) {
	// User implement Stringer interface!
	// type Stringer interface {
	// 		String() string
	// }
	u2 := User{Names: []string{"kaku", "kaku.ko.saku"}, Age: 18}
	t.Log(u2)

	// interface used to do IoC
	var student Student
	student = &u2

	// anonymous func with a argument implement Study function!
	doStudy := func(s Student) {
		s.Study()
	}
	doStudy(student)

	// you can do this too, implicit conversion.
	doStudy(&u2)
}
