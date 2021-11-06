package main

import (
	"log"

	"github.com/d-fal/bverify/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Cannot execute the program due to this error : %v\n", err)
	}
}
