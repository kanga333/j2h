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
}

// NewPrimitivePrinter creates and returns a new PrimitivePrinter.
func NewPrimitivePrinter(depth int, colName, typeName string) *PrimitivePrinter {

	return &PrimitivePrinter{
		depth:    depth,
		colName:  colName,
		typeName: typeName,
	}
}

// Print prints one line of hive ddl corresponding to the primitive type.
func (p PrimitivePrinter) Print() string {
	return fmt.Sprintf("%s%s %s,\n", printIndent(p.depth), p.colName, p.typeName)
}

func printIndent(depth int) string {
	return strings.Repeat("  ", depth)
}
