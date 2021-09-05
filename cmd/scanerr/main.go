package main

import (
	"fmt"
	"os"

	"github.com/libmonsoon-dev/scanerr/app"
	"github.com/libmonsoon-dev/scanerr/config"
)

func main() {
	conf := config.DefaultConfig()
	s := app.NewScanerr(conf)
	result, err := s.Scan(os.Args[1])
	if err != nil {
		fmt.Println("runtime error:", err)
		os.Exit(1)
	}

	fmt.Println(result)
}
