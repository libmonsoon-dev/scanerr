package main

import (
	"fmt"
	"io"

	"github.com/libmonsoon-dev/scanerr/testdata/file-not-found/lib"
)

func run() error {
	const path = "/not-exist"
	f, err := lib.FS.Open(path)
	if err != nil {
		return fmt.Errorf("open %v: %w", path, err)
	}
	defer f.Close()

	_, err = io.ReadAll(f)
	if err != nil {
		return fmt.Errorf("read %v: %w", path, err)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		//Output:
		//runtime error: open /not-exist: open /not-exist: file does not exist
		fmt.Printf("runtime error: %v", err)
	}
}
