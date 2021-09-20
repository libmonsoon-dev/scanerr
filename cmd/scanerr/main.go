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

	appConf := config.DefaultAppConfig()
	s := scanerr.NewScanerr(appConf)

	scannerConf := config.DefaultScannerConfig()
	for _, arg := range args {
		result, err := s.Scan(arg, scannerConf)
		if err != nil {
			fmt.Println("runtime error:", err)
			os.Exit(1)
		}
		fmt.Println(result)
	}
}
