package main

import (
	"fmt"

	"github.com/mwielbut/pointy"
)

func main() {
	foo := pointy.Int64(2018)
	fmt.Println("foo contains value:", *foo)

	bar := pointy.String("point to me")
	fmt.Println("bar contains value:", *bar)
}
