package main

import "testing"

func TestPrimitivePrinter(t *testing.T) {
	intPrinter := NewPrimitivePrinter(1, "test_int", "int")
	inthql := intPrinter.Print()
	expecthql := "  test_int int,\n"
	if inthql != expecthql {
		t.Errorf("result of print should be \n\"%s\", but \n\"%s\"", expecthql, inthql)
	}
}
