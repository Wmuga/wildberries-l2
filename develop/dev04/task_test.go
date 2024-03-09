package main

import (
	"reflect"
	"testing"
)

var (
	words    = []string{"пятак", "тяпка", "пятка", "листок", "слиток", "столик", "моск"}
	expected = map[string][]string{
		"пятак":  {"пятак", "пятка", "тяпка"},
		"листок": {"листок", "слиток", "столик"},
	}
)

func TestAnagrams(t *testing.T) {
	res := FindAnagrams(words)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Wrong results\n\rGot : %#v\n\rWant: %#v\n", res, expected)
	}
}
