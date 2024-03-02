package main

import "testing"

type TestCase struct {
	input  string
	output string
	nf     bool
}

func runTests(t *testing.T, cases []TestCase) {
	for i, tc := range cases {
		out, err := UnpackString(tc.input)
		if err != nil {
			if tc.nf {
				continue
			}
			t.Error("Test", i, err)
			continue
		}
		if out != tc.output {
			t.Errorf("Test %d: Expected %s Got %s", i, tc.output, out)
		}
	}
}

func TestNormal(t *testing.T) {
	cases := []TestCase{{"a4bc2d5e", "aaaabccddddde", false}, {"abcd", "abcd", false}}
	runTests(t, cases)
}

func TestEscape(t *testing.T) {
	cases := []TestCase{{`qwe\4\5`, "qwe45", false}, {`qwe\45`, "qwe44444", false}, {`qwe\\5`, `qwe\\\\\`, false}}
	runTests(t, cases)
}

func TestErrors(t *testing.T) {
	cases := []TestCase{{`45`, "", true}, {`\`, "", true}, {`\\\`, ``, true}}
	runTests(t, cases)
}
