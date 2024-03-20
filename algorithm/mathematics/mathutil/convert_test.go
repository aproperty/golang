package mathutil

import "testing"

func TestDecimalToAny(t *testing.T) {
	got := DecimalToAny(10, 8)
	want := "12"
	if got != want {
		t.Errorf("excepted:%v, got:%v", want, got)
	}
}
