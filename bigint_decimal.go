// Package implements a fixed-point BigIntDecimal
package fpd

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"strings"
)

// BigIntDecimal represents a fixed-point BigIntDecimal.
type BigIntDecimal struct {
	value *big.Int
	scale int
}

// New returns a new fixed-point BigIntDecimal
func New(value int64, scale int) *BigIntDecimal {
	return &BigIntDecimal{big.NewInt(value), scale}
}

// NewFromString returns a new fixed-point BigIntDecimal based
// on the given string
func NewFromString(value string, scale int) (*BigIntDecimal, error) {
	dValue := big.NewInt(0)
	_, ok := dValue.SetString(value, 10)
	if !ok {
		return nil, errors.New("can't convert to BigIntDecimal")
	}

	return &BigIntDecimal{dValue, scale}, nil
}

func NewFromFloat(value float64, scale int) *BigIntDecimal {
	scaleMul := math.Pow(10, -float64(scale))
	intValue := int64(value * scaleMul)
	dValue := big.NewInt(intValue)

	return &BigIntDecimal{dValue, scale}
}

// Rescale returns a rescaled version of the BigIntDecimal. Returned
// BigIntDecimal may be less precise if the given scale is bigger
// than the initial scale of the BigIntDecimal
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
func (d BigIntDecimal) rescale(scale int) *BigIntDecimal {
	diff := int(math.Abs(float64(scale - d.scale)))
	value := big.NewInt(0).Set(d.value)
	ten := big.NewInt(10)

	for diff > 0 {
		if scale > d.scale {
			value = value.Quo(value, ten)
		} else if scale < d.scale {
			value = value.Mul(value, ten)
		}

		diff--
	}

	return &BigIntDecimal{value, scale}
}

func (d *BigIntDecimal) Abs() *BigIntDecimal {
	d2Value := big.NewInt(0).Abs(d.value)
	return &BigIntDecimal{d2Value, d.scale}
}

// Add adds d to d2 and return d3
func (d *BigIntDecimal) Add(d2 *BigIntDecimal) *BigIntDecimal {
	d3Value := big.NewInt(0).Add(d.value, d2.rescale(d.scale).value)
	return &BigIntDecimal{d3Value, d.scale}
}

// Sub subtracts d2 from d and returns d3
func (d *BigIntDecimal) Sub(d2 *BigIntDecimal) *BigIntDecimal {
	baseScale := smallestOf(d.scale, d2.scale)
	rd := d.rescale(baseScale)
	rd2 := d2.rescale(baseScale)

	d3Value := big.NewInt(0).Sub(rd.value, rd2.value)
	d3 := &BigIntDecimal{d3Value, baseScale}
	return d3.rescale(d.scale)
}

// Mul multiplies d with d2 and returns d3
func (d *BigIntDecimal) Mul(d2 *BigIntDecimal) *BigIntDecimal {
	baseScale := smallestOf(d.scale, d2.scale)
	rd := d.rescale(baseScale)
	rd2 := d2.rescale(baseScale)

	d3Value := big.NewInt(0).Mul(rd.value, rd2.value)
	d3 := &BigIntDecimal{d3Value, 2 * baseScale}
	return d3.rescale(d.scale)
}

// Mul divides d by d2 and returns d3
func (d *BigIntDecimal) Div(d2 *BigIntDecimal) *BigIntDecimal {
	baseScale := -int(math.Pow(float64(smallestOf(d.scale, d2.scale)), 2))

	rd := d.rescale(baseScale + d.scale)
	rd2 := d2.rescale(baseScale)

	d3Value := big.NewInt(0).Div(rd.value, rd2.value)

	d3 := &BigIntDecimal{d3Value, d.scale}
	return d3.rescale(d.scale)
}

// Cmp compares x and y and returns -1, 0 or 1
//
// Example
//
//-1 if x <  y
// 0 if x == y
//+1 if x >  y
//
func (d *BigIntDecimal) Cmp(d2 *BigIntDecimal) int {
	smallestScale := smallestOf(d.scale, d2.scale)
	rd := d.rescale(smallestScale)
	rd2 := d2.rescale(smallestScale)

	return rd.value.Cmp(rd2.value)
}

func (d *BigIntDecimal) Scale() int {
	return d.scale
}

// String returns the string representatino of the BigIntDecimal
//
// Example:
//
//     d := New(-12345, -3)
//     println(d.String())
//
// Output:
//
//     -12345
//
func (d *BigIntDecimal) String() string {
	return d.value.String()
}

// String returns the string representatino of the BigIntDecimal
// with the fixed point
//
// Example:
//
//     d := New(-12345, -3)
//     println(d.String())
//
// Output:
//
//     -12.345
//
func (d *BigIntDecimal) FormattedString() string {
	if d.scale >= 0 {
		return d.rescale(0).value.String()
	}

	abs := big.NewInt(0).Abs(d.value)
	str := abs.String()

	var a, b string
	if len(str) >= -d.scale {
		a = str[:len(str)+d.scale]
		b = str[len(str)+d.scale:]
	} else {
		a = "0"

		num0s := -d.scale - len(str)
		b = strings.Repeat("0", num0s) + str
	}

	if d.value.Sign() < 0 {
		return fmt.Sprintf("-%v.%v", a, b)
	}

	return fmt.Sprintf("%v.%v", a, b)
}

// StringScaled first scales the BigIntDecimal then calls .String() on it.
func (d *BigIntDecimal) StringScaled(scale int) string {
	return d.rescale(scale).String()
}

func smallestOf(x, y int) int {
	if x >= y {
		return y
	}
	return x
}
