package main

import (
	"testing"

	"github.com/Wmuga/wildberries-l2/develop/dev03/flags"
	"golang.org/x/exp/slices"
)

type TestCase struct {
	flags  flags.Flags
	input  []string
	output []string
}

func runTests(t *testing.T, prefix string, cases []TestCase) {
	for i, tc := range cases {
		out, err := Sort(tc.flags, tc.input)
		if err != nil {
			t.Error("Test", i, err)
			continue
		}
		if slices.Compare(out, tc.output) != 0 {
			t.Errorf("%s: Test %d: Expected %#v Got %#v", prefix, i, tc.output, out)
		}
	}
}

func TestBase(t *testing.T) {
	emptyFlags := flags.Flags{}
	cases := []TestCase{{emptyFlags, []string{"a", "b", "c"}, []string{"a", "b", "c"}}, {emptyFlags, []string{"cab", "abcc", "abc", "bac"}, []string{"abc", "abcc", "bac", "cab"}}}
	runTests(t, "Basic", cases)
}

func TestColon(t *testing.T) {
	colonFlags := flags.Flags{Colon: 1}
	cases := []TestCase{{colonFlags, []string{"a", "b", "c"}, []string{"a", "b", "c"}}, {colonFlags, []string{"abc cab", "bac abcc", "cabb abc", "aaa bac"}, []string{"cabb abc", "bac abcc", "aaa bac", "abc cab"}}}
	runTests(t, "Colon", cases)
}

func TestInt(t *testing.T) {
	intFlags := flags.Flags{ByInt: true}
	cases := []TestCase{{intFlags, []string{"3", "1", "2"}, []string{"1", "2", "3"}}, {intFlags, []string{"abc", "1123", "768", "44"}, []string{"44", "768", "1123", "abc"}}}
	runTests(t, "Int", cases)
}

func TestReverse(t *testing.T) {
	strFlags := flags.Flags{Reverse: true}
	cases := []TestCase{{strFlags, []string{"a", "b", "c"}, []string{"c", "b", "a"}}, {strFlags, []string{"cab", "abcc", "abc", "bac"}, []string{"cab", "bac", "abcc", "abc"}}}
	runTests(t, "Str Reverse", cases)
	intFlags := flags.Flags{ByInt: true, Reverse: true}
	cases = []TestCase{{intFlags, []string{"3", "1", "2"}, []string{"3", "2", "1"}}, {intFlags, []string{"abc", "1123", "768", "44"}, []string{"abc", "1123", "768", "44"}}}
	runTests(t, "Int Reverse", cases)
}

func TestUniques(t *testing.T) {
	uniqFlags := flags.Flags{Unique: true}
	cases := []TestCase{{uniqFlags, []string{"a", "b", "c"}, []string{"a", "b", "c"}}, {uniqFlags, []string{"cab", "abcc", "abc", "abc", "bac", "bac"}, []string{"abc", "abcc", "bac", "cab"}}}
	runTests(t, "Int", cases)
}
