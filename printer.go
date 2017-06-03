package main

import (
	"fmt"
	"strings"
)

// Printer is an interface that prints ddl of hive.
type Printer interface {
	Print() string
}

// PrimitivePrinter is a printer structure corresponding to the primitive type of hive.
type PrimitivePrinter struct {
	depth    int
	colName  string
	typeName string
	sep      string
}

// StructPrinter is a printer structure corresponding to the struct type of hive.
type StructPrinter struct {
	depth    int
	colName  string
	typeName string
	sep      string
	member   []Printer
}

// NewPrimitivePrinter creates and returns a new PrimitivePrinter.
func NewPrimitivePrinter(depth int, colName, typeName string, inStruct bool) *PrimitivePrinter {
	var sep string
	if inStruct {
		sep = ":"
	} else {
		sep = " "
	}

	return &PrimitivePrinter{
		depth:    depth,
		colName:  colName,
		typeName: typeName,
		sep:      sep,
	}
}

// NewStructPrinter creates and returns a new StructPrinter.
func NewStructPrinter(depth int, colName string, inStruct bool, plist []Printer) *StructPrinter {
	var sep string
	if inStruct {
		sep = ":"
	} else {
		sep = " "
	}

	return &StructPrinter{
		depth:    depth,
		colName:  colName,
		typeName: "struct",
		sep:      sep,
		member:   plist,
	}
}

// Print prints one line of hive ddl corresponding to the primitive type.
func (p PrimitivePrinter) Print() string {
	return fmt.Sprintf("%s%s%s%s,", printIndent(p.depth), p.colName, p.sep, p.typeName)
}

// Print prints one line of hive ddl corresponding to the primitive type.
func (p StructPrinter) Print() string {
	structPirntHeader := fmt.Sprintf("%s%s%s%s<\n", printIndent(p.depth), p.colName, p.sep, p.typeName)
	structPirntFooter := fmt.Sprintf("\n%s>", printIndent(p.depth))

	var mPrints []string
	for _, v := range p.member {
		mPrints = append(mPrints, v.Print())
	}
	mPrint := strings.Join(mPrints, "\n")

	return structPirntHeader + mPrint + structPirntFooter
}

func printIndent(depth int) string {
	return strings.Repeat("  ", depth)
}
