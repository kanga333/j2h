package main

import (
	"strings"

	"github.com/tidwall/gjson"
)

// LoadJSON reads json as a string and returns a list of printer.
func LoadJSON(json string) []Printer {
	return convertJSON(1, json, false)
}

func convertJSON(depth int, json string, inStruct bool) []Printer {
	var plist []Printer

	var sep string
	if inStruct {
		sep = ":"
	} else {
		sep = " "
	}

	result := gjson.Parse(json)
	result.ForEach(func(k, v gjson.Result) bool {
		if v.Type.String() == "JSON" {
			if strings.HasPrefix(v.Raw, "{") {
				m := convertJSON(depth+1, v.Raw, true)
				p := NewStructPrinter(depth, k.String(), sep, m)
				plist = append(plist, p)
			} else if strings.HasPrefix(v.Raw, "[") {
				children := v.Array()
				var element string
				for _, v := range children {
					if element == "" {
						element = determineTyepOfHive(v.Raw, v.Type.String())
					} else if element != determineTyepOfHive(v.Raw, v.Type.String()) {
						element = "binary"
						break
					}
				}
				p := NewPrimitiveArrayPrinter(depth, k.String(), sep, element)
				plist = append(plist, p)
			}

		} else {
			t := determineTyepOfHive(v.Raw, v.Type.String())
			p := NewPrimitivePrinter(depth, k.String(), t, sep)
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
