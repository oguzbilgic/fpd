// Package implements a fixed-point decimal
package fpd

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Decimal represents a fixed-point decimal.
type Decimal struct {
	value int
	scale int
}

// New returns a new fixed-point decimal
func New(value int, scale int) *Decimal {
	return &Decimal{value, scale}
}

// Rescale returns a rescaled version of the decimal. Returned
// decimal may be less precise if the given scale is bigger
// than the initial scale of the Decimal
//
// Example:
//
// 	d := New(12345, -4)
//	d2 := d.rescale(-1)
//	d3 := d2.rescale(-4)
//	println(d1)
//	println(d2)
//	println(d3)
//
// Output:
//
//	1.2345
//	1.2
//	1.2000
//
func (d *Decimal) rescale(scale int) *Decimal {
	var value int
	if scale > d.scale {
		diff := scale - d.scale
		div := int(math.Pow10(diff))
		value = d.value / div
	} else if scale == d.scale {
		value = d.value
	} else {
		diff := d.scale - scale
		mult := int(math.Pow10(diff))
		value = d.value * mult
	}

	return New(value, scale)
}

func (d *Decimal) Add(d2 *Decimal) *Decimal {
	d3Value := d.value + d2.rescale(d.scale).value
	return New(d3Value, d.scale)
}

func (d *Decimal) Sub(d2 *Decimal) *Decimal {
	smallestScale := smallestOf(d.scale, d2.scale)
	dR := d.rescale(smallestScale)
	d2R := d2.rescale(smallestScale)

	d3Value := dR.value - d2R.value
	d3 := New(d3Value, smallestScale)
	return d3.rescale(d.scale)
}

func (d *Decimal) Mul(d2 *Decimal) *Decimal {
	scale := smallestOf(d.scale, d2.scale)
	d3Value := d.rescale(scale).value * d2.rescale(scale).value
	d3 := New(d3Value, scale*2)
	return d3.rescale(d.scale)
}

func (d *Decimal) Div(d2 *Decimal) *Decimal {
	return d.DivScale(d2, d.scale)
}

// DivScale makes the division on the given scale
func (d *Decimal) DivScale(d2 *Decimal, scale int) *Decimal {
	smallestScale := smallestOf(d.scale, d2.scale)
	d3Scale := -int(math.Pow(float64(smallestScale), 2))

	d3Value := float64(d.rescale(d3Scale).value) / float64(d2.rescale(d3Scale).value)
	d3Value = d3Value * math.Pow10(-scale)

	return New(int(d3Value), scale)
}

// Cmp compares x and y and returns -1, 0 or 1
//
// Example
//
//	-1 if x <  y
//	 0 if x == y
//	+1 if x >  y
//
func (d *Decimal) Cmp(d2 *Decimal) int {
	smallestScale := smallestOf(d.scale, d2.scale)
	dR := d.rescale(smallestScale)
	d2R := d2.rescale(smallestScale)

	if dR.value > d2R.value {
		return 1
	} else if dR.value == d2R.value {
		return 0
	} else {
		return -1
	}
}

func (d *Decimal) Int() int {
	return d.value
}

func (d *Decimal) Scale() int {
	return d.scale
}

// String returns the string representatino of the decimal
// with the trailing zeros based on the scale
//
// FIXME: negative value breaks everything
// FIXME: positive scale breaks everything
func (d *Decimal) String() string {
	strValue := strconv.Itoa(d.value)

	var a, b string
	if -len(strValue)+1 <= d.scale {
		a = strValue[:len(strValue)+d.scale]
		b = strValue[len(strValue)+d.scale:]
	} else {
		a = "0"
		num0s := -len(strValue) - d.scale
		b = strings.Repeat("0", num0s) + strValue
	}
	return fmt.Sprintf("%v.%v", a, b)
}

// StringScaled first scales the decimal then calls .String() on it.
func (d *Decimal) StringScaled(scale int) string {
	return d.rescale(scale).String()
}

// ShortString returns the string representation of the decimal
// with out the trailing zeros
//
// FIXME: Implement
func (d *Decimal) StringShort() string {
	return d.String()
}

func smallestOf(x, y int) int {
	if x >= y {
		return y
	}
	return x
}
