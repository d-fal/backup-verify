package main

import (
	"log"
	"runtime"

	"github.com/d-fal/bverify/cmd"
)

func main() {

	runtime.GOMAXPROCS(1)

	if err := cmd.Execute(); err != nil {
		log.Fatalf("Cannot execute the program due to this error : %v\n", err)
	}
}
