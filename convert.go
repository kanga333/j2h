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

	/*As gjson.ForEach can not propagate errors, put the child elements of json in the array.
	  Use two arrays of key and value instead of map because we want to guarantee the order of input.*/
	jsonKeys := []string{}
	jsonVals := []gjson.Result{}
	result := gjson.Parse(json)
	result.ForEach(func(k, v gjson.Result) bool {
		jsonKeys = append(jsonKeys, k.String())
		jsonVals = append(jsonVals, v)
		return true
	})

	var printerList []Printer
	delimiter := determineDelimiter(inStruct)

	for i, jsonVal := range jsonVals {
		var printer Printer
		hiveType := determineTyepOfHive(jsonVal.Raw, jsonVal.Type.String())
		switch hiveType {
		case "array":
			printer = newArrayPrinter(depth, jsonKeys[i], delimiter, jsonVal)
		case "struct":
			m := convertJSON(depth+1, jsonVal.Raw, true)
			printer = NewStructPrinter(depth, jsonKeys[i], delimiter, m)
		default:
			printer = NewPrimitivePrinter(depth, jsonKeys[i], hiveType, delimiter)
		}
		printerList = append(printerList, printer)

	}

	return printerList
}

func determineArrayType(children []gjson.Result) string {
	var arrayType string
	for _, v := range children {
		newArrayType := determineTyepOfHive(v.Raw, v.Type.String())
		if arrayType == "" {
			arrayType = newArrayType
			continue
		}
		arrayType = compareArrayTypes(arrayType, newArrayType)
		if arrayType == "binary" {
			break
		}
	}
	return arrayType
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
	case "JSON":
		return determineComposite(jsonRaw)
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

func determineComposite(s string) string {
	if strings.HasPrefix(s, "{") {
		return "struct"
	} else if strings.HasPrefix(s, "[") {
		return "array"
	}
	return ""
}

func determineDelimiter(inStruct bool) string {
	if inStruct {
		return ":"
	}
	return " "
}

func compareArrayTypes(oldType, newType string) string {
	if oldType == newType {
		return oldType
	}
	return "binary"

}

func newArrayPrinter(depth int, colName, delimiter string, jsonVal gjson.Result) Printer {
	arrayType := determineArrayType(jsonVal.Array())
	switch arrayType {
	case "array":
		p := newArrayPrinter(depth+1, "", "", jsonVal.Get("0"))
		return NewMultipleArrayPrinter(depth, colName, delimiter, p)
	case "struct":
		m := convertJSON(depth+2, jsonVal.Get("0").Raw, true)
		return NewStructArrayPrinter(depth, colName, delimiter, m)
	default:
		return NewPrimitiveArrayPrinter(depth, colName, delimiter, arrayType)
	}

}
