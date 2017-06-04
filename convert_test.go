package main

import "testing"

func TestLoadJSON(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{input: `{"null_type": null}`, expect: "  null_type binary"},
		{input: `{"true_type": true}`, expect: "  true_type boolean"},
		{input: `{"false_type": false}`, expect: "  false_type boolean"},
		{input: `{"int_type": 1}`, expect: "  int_type int"},
		{input: `{"double_type": 1.1}`, expect: "  double_type double"},
		{input: `{"string_type": "string"}`, expect: "  string_type string"},
		{
			input:  `{"struct_type": {"child_a": 1,"child_b": "string"}}`,
			expect: "  struct_type struct<\n    child_a:int,\n    child_b:string\n  >",
		},
		{
			input:  `{"nest_struct": {"child_a": 1,"child_nest_1": {"child_nest_2":{"child_b": 1,"child_c": "string",}}}}`,
			expect: "  nest_struct struct<\n    child_a:int,\n    child_nest_1:struct<\n      child_nest_2:struct<\n        child_b:int,\n        child_c:string\n      >\n    >\n  >",
		},
		{input: `{"array_type": [10,21,20]}`, expect: "  array_type array<int>"},
		{input: `{"mixed_array_type": ["hoge",21,20]}`, expect: "  mixed_array_type array<binary>"},
		{
			input:  `{"struct_type": {"array_type": [10,21,20],"child_b": "string"}}`,
			expect: "  struct_type struct<\n    array_type:array<int>,\n    child_b:string\n  >",
		},
		{
			input:  `{"array_type": [{"struct_type": {"child_a": 1,"child_b": "string"}},{"struct_type": {"child_a": 1,"child_b": "string"}}]}`,
			expect: "  array_type array<\n    struct<\n      struct_type:struct<\n        child_a:int,\n        child_b:string\n      >\n    >\n  >",
		},
		{
			input:  `{"array_type": [{"child_a": 1,"child_b": "string"},{"child_a": 1,"child_b": "string"}]}`,
			expect: "  array_type array<\n    struct<\n      child_a:int,\n      child_b:string\n    >\n  >",
		},
		{
			input:  `{"multi_array_type": [[10,21,20],[10,21,22]]}`,
			expect: "  multi_array_type array<\n    array<int>\n  >",
		},
		{
			input:  `{"multi_array_type": [[{"struct_type": {"child_a": 1,"child_array": ["a","b"]}},{"struct_type": {"child_a": 1,"child_array": ["a","b"]}}],[{"struct_type": {"child_a": 1,"child_array": ["a","b"]}}]]}`,
			expect: "  multi_array_type array<\n    array<\n      struct<\n        struct_type:struct<\n          child_a:int,\n          child_array:array<string>\n        >\n      >\n    >\n  >",
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
