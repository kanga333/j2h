package main

import "testing"

func TestPrimitivePrinter(t *testing.T) {
	intPrinter := NewPrimitivePrinter(1, "test_int", "int", false)
	inthql := intPrinter.Print()
	expecthql := "  test_int int,\n"
	if inthql != expecthql {
		t.Errorf("result of print should be \n\"%s\", but \n\"%s\"", expecthql, inthql)
	}
}

func TestStructPrinter(t *testing.T) {

	childPrinterA := NewPrimitivePrinter(2, "childa", "int", true)
	childPrinterB := NewPrimitivePrinter(2, "childb", "string", true)
	plist := []Printer{childPrinterA, childPrinterB}

	structPrinter := NewStructPrinter(1, "structprinter", false, plist)

	structhql := structPrinter.Print()
	expecthql := "  structprinter struct<\n    childa:int,\n    childb:string,\n  >\n"
	if structhql != expecthql {
		t.Errorf("result of print should be \n%q\nbut\n%q", expecthql, structhql)
	}
}
