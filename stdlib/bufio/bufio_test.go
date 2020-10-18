package bufio_test

import (
	"bufio"
	"fmt"
	"strings"
)

func ExampleScan() {
	t := "first line\nsecond line\n"
	scanner := bufio.NewScanner(strings.NewReader(t))
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	// Output:
	// first line
	// second line
}

func ExampleScanCustom() {
	t := "1,2,3,4,"
	scanner := bufio.NewScanner(strings.NewReader(t))
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// this code is bug prone when t with chinese char...
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		// There is one final token to be delivered, which may be th empty string!
		if !atEOF {
			return 0, nil, nil
		}
		return 0, data, bufio.ErrFinalToken
	}
	scanner.Split(onComma)
	for scanner.Scan() {
		fmt.Printf("%q ", scanner.Text())
	}
	// Output:
	// "1" "2" "3" "4" ""
}
