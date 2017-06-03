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
