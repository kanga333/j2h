package main

import (
	"strings"

	"fmt"

	"github.com/tidwall/gjson"
)

// ConvertJSONTOHQL reads json as a string and returns a hive ddl string.
func ConvertJSONTOHQL(json string) (string, error) {
	plists, err := LoadJSON(json)
	if err != nil {
		return "", err
	}
	var ddl []string
	for _, v := range plists {
		ddl = append(ddl, v.Print())
	}
	return PrintHeader() + "\n" + strings.Join(ddl, ",\n") + "\n" + PrintFooter(), nil
}

// LoadJSON reads json as a string and returns a list of printer.
func LoadJSON(json string) ([]Printer, error) {
	if !gjson.Valid(json) {
		return nil, fmt.Errorf("input value is not json format")
	}
	return convertJSON(1, json, false)
}

func convertJSON(depth int, json string, inStruct bool) ([]Printer, error) {

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
		hiveType, err := determineTyepOfHive(jsonVal.Raw, jsonVal.Type.String())
		if err != nil {
			return nil, err
		}
		switch hiveType {
		case "array":
			printer, err = newArrayPrinter(depth, jsonKeys[i], delimiter, jsonVal)
			if err != nil {
				return nil, err
			}
		case "struct":
			m, err := convertJSON(depth+1, jsonVal.Raw, true)
			if err != nil {
				return nil, err
			}
			printer = NewStructPrinter(depth, jsonKeys[i], delimiter, m)
		default:
			printer = NewPrimitivePrinter(depth, jsonKeys[i], hiveType, delimiter)
		}
		printerList = append(printerList, printer)

	}

	return printerList, nil
}

func determineArrayType(children []gjson.Result) (string, error) {
	var arrayType string
	for _, v := range children {
		newArrayType, err := determineTyepOfHive(v.Raw, v.Type.String())
		if err != nil {
			return "", err
		}
		if arrayType == "" {
			arrayType = newArrayType
			continue
		}
		arrayType = compareArrayTypes(arrayType, newArrayType)
		if arrayType == "binary" {
			break
		}
	}
	return arrayType, nil
}

func determineTyepOfHive(jsonRaw, jsonType string) (string, error) {
	switch jsonType {
	case "True", "False":
		return "boolean", nil
	case "Null":
		return "binary", nil
	case "String":
		return "string", nil
	case "Number":
		return determineNumberTyepOfHive(jsonRaw), nil
	case "JSON":
		return determineComposite(jsonRaw)
	default:
		return "", fmt.Errorf("unexpected composite type: %v", jsonType)
	}
}

func determineNumberTyepOfHive(numStr string) string {
	if strings.Contains(numStr, ".") {
		return "double"
	}
	return "int"
}

func determineComposite(s string) (string, error) {
	if strings.HasPrefix(s, "{") {
		return "struct", nil
	} else if strings.HasPrefix(s, "[") {
		return "array", nil
	}
	return "", fmt.Errorf("unexpected composite type: %v", s)
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

func newArrayPrinter(depth int, colName, delimiter string, jsonVal gjson.Result) (Printer, error) {
	arrayType, err := determineArrayType(jsonVal.Array())
	if err != nil {
		return nil, err
	}
	switch arrayType {
	case "array":
		p, err := newArrayPrinter(depth+1, "", "", jsonVal.Get("0"))
		if err != nil {
			return nil, err
		}
		return NewMultipleArrayPrinter(depth, colName, delimiter, p), nil
	case "struct":
		m, err := convertJSON(depth+2, jsonVal.Get("0").Raw, true)
		if err != nil {
			return nil, err
		}
		return NewStructArrayPrinter(depth, colName, delimiter, m), nil
	default:
		return NewPrimitiveArrayPrinter(depth, colName, delimiter, arrayType), nil
	}

}
