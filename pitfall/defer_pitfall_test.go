// Author: kaku
// Date: 2020/10/19
//
// GitHub:
//	https://github.com/kakukosaku
//
package pitfall_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Demo1 defer can change function non-named return value.
func Demo1(t *testing.T) int {
	i := 1
	defer func() {
		i++
		assert.Equal(t, 2, i)
	}()
	assert.Equal(t, 1, i)
	return i
}

// Demo2 but defer can change function named return value.
func Demo2(t *testing.T) (i int) {
	i = 1
	defer func() {
		i++
		assert.Equal(t, 2, i)
	}()
	assert.Equal(t, 1, i)
	return i
}

// Demo3 defer execute when 1.function return; 2.function end, return named value; 3. panic cause defer func execute with order FILO
func Demo3(t *testing.T) int {
	i := 1
	defer func() {
		if err := recover(); err != nil {
			if value, ok := err.(string); ok {
				t.Log(value)
				// cannot change non-named return value!
				i = 3
			}
		}
	}()
	panic("Crush!")
	// unreachable code
	i++
	return i
}

func TestDeferPitfall(t *testing.T) {
	rest1 := Demo1(t)
	assert.Equal(t, 1, rest1)

	rest2 := Demo2(t)
	assert.Equal(t, 2, rest2)

	rest3 := Demo3(t)
	assert.Equal(t, 0, rest3)
}
