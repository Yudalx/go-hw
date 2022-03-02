package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	reverseResult := stringutil.Reverse("Hello, OTUS!")
	fmt.Println(reverseResult)
}
