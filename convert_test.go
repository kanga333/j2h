package main

import "testing"

func TestLoadJSON(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{input: `{"null_type": null}`, expect: "  null_type binary,\n"},
		{input: `{"true_type": true}`, expect: "  true_type bool,\n"},
		{input: `{"false_type": false}`, expect: "  false_type bool,\n"},
		{input: `{"int_type": 1}`, expect: "  int_type int,\n"},
		{input: `{"float_type": 1.1}`, expect: "  float_type float,\n"},
		{input: `{"string_type": "string"}`, expect: "  string_type string,\n"},
		{
			input:  `{"struct_type": {"child_a": 1,"child_b": "string"}}`,
			expect: "  struct_type struct<\n    child_a:int,\n    child_b:string,\n  >\n",
		},
		{
			input:  `{"nest_struct": {"child_a": 1,"child_nest_1": {"child_nest_2":{"child_b": 1,"child_c": "string",}}}}`,
			expect: "  nest_struct struct<\n    child_a:int,\n    child_nest_1:struct<\n      child_nest_2:struct<\n        child_b:int,\n        child_c:string,\n      >\n    >\n  >\n",
		},
	}

	for _, test := range tests {
		plist := LoadJSON(test.input)
		if plist == nil {
			t.Fatalf("should not be nil for %s", test.input)
		}
		for _, p := range plist {
			result := p.Print()
			if result != test.expect {
				t.Errorf("result of print should be \n%q\nbut\n%q", test.expect, result)
			}
		}
	}
}
