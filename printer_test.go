package main

import "testing"

func TestPrimitivePrinter(t *testing.T) {
	intPrinter := NewPrimitivePrinter(1, "test_int", "int", " ")
	inthql := intPrinter.Print()
	expecthql := "  test_int int"
	if inthql != expecthql {
		t.Errorf("result of print should be \n%q\nbut\n%q", expecthql, inthql)
	}
}

func TestStructPrinter(t *testing.T) {
	childPrinterA := NewPrimitivePrinter(2, "childa", "int", ":")
	childPrinterB := NewPrimitivePrinter(2, "childb", "string", ":")
	plist := []Printer{childPrinterA, childPrinterB}

	structPrinter := NewStructPrinter(1, "structprinter", " ", plist)

	structhql := structPrinter.Print()
	expecthql := "  structprinter struct<\n    childa:int,\n    childb:string\n  >"
	if structhql != expecthql {
		t.Errorf("result of print should be \n%q\nbut\n%q", expecthql, structhql)
	}
}

func TestPrimitiveArrayPrinter(t *testing.T) {
	arrayPrinter := NewPrimitiveArrayPrinter(1, "arrayprinter", " ", "int")
	arrayhql := arrayPrinter.Print()
	expecthql := "  arrayprinter array<int>"
	if arrayhql != expecthql {
		t.Errorf("result of print should be \n%q\nbut\n%q", expecthql, arrayhql)
	}
}

func TestStructArrayPrinter(t *testing.T) {
	childPrinterA := NewPrimitivePrinter(3, "childa", "int", ":")
	childPrinterB := NewPrimitivePrinter(3, "childb", "string", ":")
	plist := []Printer{childPrinterA, childPrinterB}
	arrayPrinter := NewStructArrayPrinter(1, "structarrayprinter", " ", plist)
	arrayhql := arrayPrinter.Print()
	expecthql := "  structarrayprinter array<\n    struct<\n      childa:int,\n      childb:string\n    >\n  >"
	if arrayhql != expecthql {
		t.Errorf("result of print should be \n%q\nbut\n%q", expecthql, arrayhql)
	}
}

func TestMultipleArrayPrinter(t *testing.T) {
	descendant := NewPrimitiveArrayPrinter(3, "", "", "int")
	child := NewMultipleArrayPrinter(2, "", "", descendant)
	arrayPrinter := NewMultipleArrayPrinter(1, "multiarrayprinter", " ", child)
	arrayhql := arrayPrinter.Print()
	expecthql := "  multiarrayprinter array<\n    array<\n      array<int>\n    >\n  >"
	if arrayhql != expecthql {
		t.Errorf("result of print should be \n%q\nbut\n%q", expecthql, arrayhql)
	}
}
