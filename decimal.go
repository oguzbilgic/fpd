// Package implements a fixed-point Decimal
package fpd

// Decimal represents a fixed-point Decimal.
type Decimal interface {
	Abs() *Decimal

	// Add adds d to d2 and return d3
	Add(d2 *Decimal) *Decimal

	// Sub subtracts d2 from d and returns d3
	Sub(d2 *Decimal) *Decimal

	// Mul multiplies d with d2 and returns d3
	Mul(d2 *Decimal) *Decimal

	// Mul divides d by d2 and returns d3
	Div(d2 *Decimal) *Decimal

	// Cmp compares x and y and returns -1, 0 or 1
	//
	// Example
	//
	//-1 if x <  y
	// 0 if x == y
	//+1 if x >  y
	//
	Cmp(d2 *Decimal) int

	Scale() int

	// String returns the string representatino of the Decimal
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
	String() string

	// String returns the string representatino of the Decimal
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
	FormattedString() string

	// StringScaled first scales the Decimal then calls .String() on it.
	StringScaled(scale int) string
}
