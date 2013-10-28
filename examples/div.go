package main

import (
	"github.com/oguzbilgic/fpd"
	"fmt"
)

func main() {
	// 166.004
	a := fpd.New(18275499, -6)

	// 1.006
	b := fpd.New(1, -1)
	c := a.Div(b)
	d := c.Mul(fpd.New(100, 0))

	fmt.Println(c.FormattedString())
	fmt.Println(d.FormattedString())
}
