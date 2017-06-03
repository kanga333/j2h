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
	members  []Printer
}

// ArrayPrinter is a printer structure corresponding to the array type of hive.
type ArrayPrinter struct {
	depth    int
	colName  string
	typeName string
	sep      string
	member   Printer
}

// NewPrimitivePrinter creates and returns a new PrimitivePrinter.
func NewPrimitivePrinter(depth int, colName, typeName, sep string) *PrimitivePrinter {
	return &PrimitivePrinter{
		depth:    depth,
		colName:  colName,
		typeName: typeName,
		sep:      sep,
	}
}

// NewStructPrinter creates and returns a new StructPrinter.
func NewStructPrinter(depth int, colName, sep string, plist []Printer) *StructPrinter {
	return &StructPrinter{
		depth:    depth,
		colName:  colName,
		typeName: "struct",
		sep:      sep,
		members:  plist,
	}
}

// NewPrimitiveArrayPrinter creates and returns a new ArrayPrinter whose element is a primitive type.
func NewPrimitiveArrayPrinter(depth int, colName, sep, elementType string) *ArrayPrinter {
	p := NewPrimitivePrinter(0, "", elementType, "")

	return &ArrayPrinter{
		depth:    depth,
		colName:  colName,
		typeName: "array",
		sep:      sep,
		member:   p,
	}
}

// Print prints one line of hive ddl corresponding to the primitive type.
func (p PrimitivePrinter) Print() string {
	return fmt.Sprintf("%s%s%s%s", printIndent(p.depth), p.colName, p.sep, p.typeName)
}

// Print prints one line of hive ddl corresponding to the primitive type.
func (p StructPrinter) Print() string {
	structPirntHeader := fmt.Sprintf("%s%s%s%s<\n", printIndent(p.depth), p.colName, p.sep, p.typeName)
	structPirntFooter := fmt.Sprintf("\n%s>", printIndent(p.depth))

	var mPrints []string
	for _, v := range p.members {
		mPrints = append(mPrints, v.Print())
	}
	mPrint := strings.Join(mPrints, ",\n")

	return structPirntHeader + mPrint + structPirntFooter
}

// Print prints one line of hive ddl corresponding to the array type.
func (p ArrayPrinter) Print() string {
	structPirntHeader := fmt.Sprintf("%s%s%s%s<", printIndent(p.depth), p.colName, p.sep, p.typeName)
	structPirntFooter := fmt.Sprint(">")

	return structPirntHeader + p.member.Print() + structPirntFooter
}

func printIndent(depth int) string {
	return strings.Repeat("  ", depth)
}
