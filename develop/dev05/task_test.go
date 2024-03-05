package main

import (
	"github.com/Wmuga/wildberries-l2/develop/dev05/flags"
	"slices"
	"testing"
)

type TestCase struct {
	flags   flags.Flags
	pattern string
	input   []string
	output  []string
}

var pattern = "r.s"

var allCases = []string{
	"ctxBefore",
	"res",
	"r.s",
	"rEs",
	"ctxAfter",
}

var noCtx = []string{
	"res",
	"r.s",
	"rEs",
}

var numbered = []string{
	"2:res",
	"3:r.s",
	"4:rEs",
}

func runTests(t *testing.T, prefix string, cases []TestCase) {
	for i, tc := range cases {
		out, err := Grep(tc.pattern, tc.flags, tc.input)
		if err != nil {
			t.Errorf("%s: Test %d %s", prefix, i, err)
			continue
		}
		if slices.Compare(out, tc.output) != 0 {
			t.Errorf("%s: Test %d: Expected %#v Got %#v", prefix, i, tc.output, out)
		}
	}
}

func TestRegex(t *testing.T) {
	tc := TestCase{
		flags.Flags{},
		pattern,
		allCases,
		noCtx,
	}
	runTests(t, "Regex", []TestCase{tc})
}

func TestFixed(t *testing.T) {
	tc := TestCase{
		flags.Flags{Fixed: true},
		pattern,
		allCases,
		[]string{"r.s"},
	}
	runTests(t, "Fixed", []TestCase{tc})
}

func TestIgnoreCase(t *testing.T) {
	tc := TestCase{
		flags.Flags{IgnoreCase: true},
		"res",
		allCases,
		[]string{"res", "rEs"},
	}
	runTests(t, "Ignore case not fixed", []TestCase{tc})
	tc.flags.Fixed = true
	runTests(t, "Ignore case fixed", []TestCase{tc})
}

func TestContext(t *testing.T) {
	tc := TestCase{
		flags.Flags{Context: 1},
		pattern,
		allCases,
		allCases,
	}
	runTests(t, "Context", []TestCase{tc})
	tc.flags.Context = 0
	tc.flags.After = 1
	tc.output = allCases[1:]
	runTests(t, "After", []TestCase{tc})
	tc.flags.After = 0
	tc.flags.Before = 1
	tc.output = allCases[:len(allCases)-1]
	runTests(t, "After", []TestCase{tc})
}

func TestInvert(t *testing.T) {
	tc := TestCase{
		flags.Flags{Invert: true},
		pattern,
		allCases,
		[]string{"ctxBefore", "ctxAfter"},
	}
	runTests(t, "Invert", []TestCase{tc})
}

func TestLineNum(t *testing.T) {
	tc := TestCase{
		flags.Flags{LineNum: true},
		pattern,
		allCases,
		numbered,
	}
	runTests(t, "Line number", []TestCase{tc})
}

func TestCount(t *testing.T) {
	tc := TestCase{
		flags.Flags{Count: true},
		pattern,
		allCases,
		[]string{"3"},
	}
	runTests(t, "Count", []TestCase{tc})
}
