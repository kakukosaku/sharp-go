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

func TestLoopPitfall(t *testing.T) {
	// for loop with closure func call
	printFunctions1 := make([]func() int, 0)
	for i := 0; i < 3; i++ {
		printFunctions1 = append(printFunctions1, func() int { return i })
	}

	for _, f := range printFunctions1 {
		assert.Equal(t, 3, f())
	}

	// for loop with non-closure func call
	printFunctions2 := make([]func() int, 0)
	for i := 0; i < 3; i++ {
		t := i
		printFunctions2 = append(printFunctions2, func() int { return t })
	}

	for i, f := range printFunctions2 {
		// Note: this!
		assert.Equal(t, i, f())
	}
}
