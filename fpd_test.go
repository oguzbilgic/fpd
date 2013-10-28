package fpd

import "testing"

func TestNewFromString1(t *testing.T) {
	a, err := NewFromString("1234", -3)

	if err != nil {
		t.Errorf("error")
	}

	if a.String() != "1234" {
		t.Errorf("error")
	}
}

func TestNewFromString2(t *testing.T) {
	a, err := NewFromString("-1234", -3)

	if err != nil {
		t.Errorf("error")
	}

	if a.String() != "-1234" {
		t.Errorf("error")
	}
}

func TestNewFromString3(t *testing.T) {
	_, err := NewFromString("qwert", -3)

	if err == nil {
		t.Errorf("error")
	}
}

func TestDecimal_Scale(t *testing.T) {
	a := New(1234, -3)
	if a.Scale() != -3 {
		t.Errorf("error")
	}
}

func TestDecimal_recale1(t *testing.T) {
	a := New(1234, -3).rescale(-5)
	if a.String() != "123400" {
		t.Errorf(a.String() + " != 123400")
	}
}

func TestDecimal_recale2(t *testing.T) {
	a := New(1234, -3).rescale(0)
	if a.String() != "1" {
		t.Errorf("error")
	}
}

func TestDecimal_recale3(t *testing.T) {
	a := New(1234, 3).rescale(0)
	if a.String() != "1234000" {
		t.Errorf("error")
	}
}

func TestDecimal_recale4(t *testing.T) {
	a := New(1234, 3).rescale(5)
	if a.String() != "12" {
		t.Errorf("error")
	}
}

func TestDecimal_recale5(t *testing.T) {
	a := New(1234, 3)
	_ = a.rescale(5)
	if a.String() != "1234" {
		t.Errorf("error")
	}
}

func TestDecimal_Abs1(t *testing.T) {
	a := New(-1234, -4)
	b := New(1234, -4)

	c := a.Abs()
	if c.Cmp(b) != 0 {
		t.Errorf("error")
	}
}

func TestDecimal_Abs2(t *testing.T) {
	a := New(-1234, -4)
	b := New(1234, -4)

	c := b.Abs()
	if c.Cmp(a) == 0 {
		t.Errorf("error")
	}
}

func TestDecimal_Add1(t *testing.T) {
	a := New(1234, -4)
	b := New(9876, 3)

	c := a.Add(b)
	if c.String() != "98760001234" {
		t.Errorf("error")
	}
}

func TestDecimal_Add2(t *testing.T) {
	a := New(1234, 3)
	b := New(9876, -4)

	c := a.Add(b)
	if c.String() != "1234" {
		t.Errorf("error")
	}
}

func TestDecimal_Sub1(t *testing.T) {
	a := New(1234, -4)
	b := New(9876, 3)

	c := a.Sub(b)
	if c.String() != "-98759998766" {
		t.Errorf(c.String())
	}
}

func TestDecimal_Sub2(t *testing.T) {
	a := New(1234, 3)
	b := New(9876, -4)

	c := a.Sub(b)
	if c.String() != "1233" {
		t.Errorf(c.String())
	}
}

func TestDecimal_Mul(t *testing.T) {
	a := New(1398699, -4)
	b := New(6, -3)

	c := a.Mul(b)
	if c.String() != "8392" {
		t.Errorf(c.String())
	}
}

func TestDecimal_Div1(t *testing.T) {
	a := New(1398699, -4)
	b := New(1006, -3)

	c := a.Div(b)
	if c.String() != "1390356" {
		t.Errorf(c.String())
	}
}

func TestDecimal_Div2(t *testing.T) {
	a := New(2345, -3)
	b := New(2, 0)

	c := a.Div(b)
	if c.String() != "1172" {
		t.Errorf(c.String())
	}
}

func TestDecimal_Div3(t *testing.T) {
	a := New(18275499, -6)
	b := New(16275499, -6)

	c := a.Div(b)
	if c.String() != "1122884" {
		t.Errorf(c.String())
	}
}
func TestDecimal_Cmp1(t *testing.T) {
	a := New(123, 3)
	b := New(-1234, 2)

	if a.Cmp(b) != 1 {
		t.Errorf("Error")
	}
}

func TestDecimal_Cmp2(t *testing.T) {
	a := New(123, 3)
	b := New(1234, 2)

	if a.Cmp(b) != -1 {
		t.Errorf("Error")
	}
}

func TestDecimal_StringScaled(t *testing.T) {
	a := New(123, 3)
	if a.StringScaled(-2) != "12300000" {
		t.Errorf("Error")
	}
}

func TestDecimal_StringScaled2(t *testing.T) {
	a := New(1234, -2)
	if a.StringScaled(0) != "12" {
		t.Errorf("Error")
	}
}

func TestDecimal_FormattedString(t *testing.T) {
	a := New(1234, -2)
	if a.FormattedString() != "12.34" {
		t.Errorf(a.FormattedString())
	}
}

func TestDecimal_FormattedString5(t *testing.T) {
	a := New(-1234, -2)
	if a.FormattedString() != "-12.34" {
		t.Errorf(a.FormattedString())
	}
}

func TestDecimal_FormattedString1(t *testing.T) {
	a := New(1234, 2)
	if a.FormattedString() != "123400" {
		t.Errorf(a.FormattedString())
	}
}

func TestDecimal_FormattedString2(t *testing.T) {
	a := New(-1234, 2)
	if a.FormattedString() != "-123400" {
		t.Errorf(a.FormattedString())
	}
}

func TestDecimal_FormattedString3(t *testing.T) {
	a := New(1234, -6)
	if a.FormattedString() != "0.001234" {
		t.Errorf(a.FormattedString())
	}
}

func TestDecimal_FormattedString4(t *testing.T) {
	a := New(1000, -6)
	if a.FormattedString() != "0.001000" {
		t.Errorf(a.FormattedString())
	}
}
