package main

import (
	"fmt"
	cnt2 "foo/counter"
	"sample/counter"
)

func main() {
	counter.Default++
	fmt.Println("main:", counter.Default)

	fmt.Println("main:", cnt2.Default)
}
