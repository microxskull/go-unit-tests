package main

import (
	"log"

	"github.com/microxskull/go-unit-tests/internal/platform/strrev"
)

func main() {
	in := "hello"
	if out, err := strrev.Reverse(in); err != nil {
		log.Println("Reversed string failed:", err)
	} else {
		log.Println("Reversed string:", out)
	}
}
