package ioutil_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Example() {
	r := strings.NewReader("I am kaku")
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	dir, err := ioutil.TempDir(".", "temp-dir-")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	c := []byte("I am kaku")
	tmpFileName := filepath.Join(dir, "temp-text.txt")
	if err := ioutil.WriteFile(tmpFileName, c, 0666); err != nil {
		log.Fatal(err)
	}

	if files, err := ioutil.ReadDir(dir); err != nil {
		log.Fatal(err)
	} else {
		for _, file := range files {
			fmt.Println(file.Name())
		}
	}

	fc, err := ioutil.ReadFile(tmpFileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(fc))
	// Output:
	// I am kaku
	// temp-text.txt
	// I am kaku
}
