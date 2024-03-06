package main

import (
	"slices"
	"testing"
)

type TestCase struct {
	flags  Flags
	input  string
	output []string
}

func runTests(t *testing.T, prefix string, cases []TestCase) {
	for i, tc := range cases {
		out := Cut(tc.flags, tc.input)
		if slices.Compare(out, tc.output) != 0 {
			t.Errorf("%s: Test %d: Expected %#v Got %#v", prefix, i, tc.output, out)
		}
	}
}

func TestCut(t *testing.T) {
	cases := []TestCase{
		// Простой тест на обрезание полей
		{Flags{FieldNums: []int{1, 3}, Delimiter: "\t"}, "123\t456\t789", []string{"123", "789"}},
		// Тот же тест, но с запятыми
		{Flags{FieldNums: []int{1, 3}, Delimiter: ","}, "123,456,789", []string{"123", "789"}},
		// Тест с полями за пределами массива
		{Flags{FieldNums: []int{1, 3, 6, 7}, Delimiter: "\t"}, "123\t456\t789", []string{"123", "789"}},
		// Тест с поднятыми флагом -s
		{Flags{FieldNums: []int{1, 3, 6, 7}, Delimiter: "\t", Separated: true}, "123,456,789", nil},
	}

	runTests(t, "AIO", cases)
}
