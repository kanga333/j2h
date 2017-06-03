package main

import (
	"strings"

	"github.com/tidwall/gjson"
)

// LoadJSON reads json as a string and returns a list of printer.
func LoadJSON(json string) []Printer {
	return convertJSON(1, json)
}

func convertJSON(depth int, json string) []Printer {
	var plist []Printer

	result := gjson.Parse(json)
	result.ForEach(func(k, v gjson.Result) bool {
		if v.Type.String() != "JSON" {
			t := determineTyepOfHive(v.Raw, v.Type.String())
			p := NewPrimitivePrinter(depth, k.String(), t)
			plist = append(plist, p)
		}
		return true

	})
	return plist
}

func determineTyepOfHive(jsonRaw, jsonType string) string {
	switch jsonType {
	case "True", "False":
		return "bool"
	case "Null":
		return "binary"
	case "String":
		return "string"
	case "Number":
		return determineNumberTyepOfHive(jsonRaw)
	default:
		return "binary"
	}
}

func determineNumberTyepOfHive(numStr string) string {
	if strings.Contains(numStr, ".") {
		return "float"
	}
	return "int"
}
