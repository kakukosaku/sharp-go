package errors_test

import (
	"errors"
	"fmt"
	"os"
)

type APIError struct {
	url       string
	timestamp int64
}

func (e APIError) Error() string {
	return fmt.Sprintf("api error: %s, %d", e.url, e.timestamp)
}

func WrapOSError() error {
	return fmt.Errorf("wrapped error: %w", os.ErrExist)
}

func Example() {
	f := func() error {
		return APIError{"api/v1/test", 1587515139}
	}

	if err := f(); err != nil {
		fmt.Println(err)
	}
	err := WrapOSError()
	if errors.Is(err, os.ErrExist) {
		fmt.Println(err)
	}
	// Output:
	// api error: api/v1/test, 1587515139
	// wrapped error: file already exists
}
