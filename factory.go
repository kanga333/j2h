package main

import "fmt"

type printerFactory struct {
}

func (pf *printerFactory) getPrinter(key, valType, value string) (*Printer, error) {
	switch valType {
	case "True", "False":
	case "Null":
	case "String":
	case "Number":
	case "JSON":
	default:
		return nil, fmt.Errorf("unexpected type: %s", valType)
	}

	return nil, nil
}

func (Pf *printerFactory) getPrimitivePrinter(key, valType, value string) (*Printer, error) {
	return nil, nil
}

func (Pf *printerFactory) getCompostitePrinter(key, valType, value string) (*Printer, error) {
	return nil, nil
}
