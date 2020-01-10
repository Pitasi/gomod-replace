package counter

import "fmt"

var Default int

func init() {
	fmt.Println("init")
	Default = 42
}
