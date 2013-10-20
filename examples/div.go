package main

import (
	"github.com/oguzbilgic/fpd"
	"fmt"
)

func main() {
	// 166.004
	a := fpd.New(10, -1)

	// 1.006
	b := fpd.New(2000, -3)
	c := a.Div(b)

	fmt.Println(c.FormattedString())
}
