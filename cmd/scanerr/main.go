package main

import (
	"fmt"
	"os"

	"github.com/libmonsoon-dev/scanerr"
	"github.com/libmonsoon-dev/scanerr/config"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Printf("Not enough arguments in call %v", os.Args[0])
		os.Exit(1)
	}

	conf := config.DefaultConfig()
	s := scanerr.NewScanerr(conf)

	for _, arg := range args {
		result, err := s.Scan(arg)
		if err != nil {
			fmt.Println("runtime error:", err)
			os.Exit(1)
		}
		fmt.Println(result)
	}
}
