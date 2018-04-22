package main

import (
	"fmt"
)

type A struct {
	X, Y int
}

type Aptr *A
type Aptrptr **A

type B struct {
	App Aptrptr
}

func main() {
	var b B

	fmt.Printf("%v\n", b.App)
}
