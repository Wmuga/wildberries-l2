package main

import "testing"

type TestCase struct {
	input  string
	output string
	nf     bool
}

type unpacker func(string) (string, error)

func runTests(t *testing.T, prefix string, testFunc unpacker, cases []TestCase) {
	for i, tc := range cases {
		out, err := testFunc(tc.input)
		if err != nil {
			if tc.nf {
				continue
			}
			t.Error("Test", i, err)
			continue
		}
		if out != tc.output {
			t.Errorf("%s: Test %d: Expected %s Got %s", prefix, i, tc.output, out)
		}
	}
}

func TestNormal(t *testing.T) {
	cases := []TestCase{{"a4bc2d5e", "aaaabccddddde", false}, {"abcd", "abcd", false}, {"", "", false}}
	runTests(t, "Unpack normal naive", UnpackString, cases)
	runTests(t, "Unpack normal states", UnpackStringState, cases)
}

func TestEscape(t *testing.T) {
	cases := []TestCase{{`qwe\4\5`, "qwe45", false}, {`qwe\45`, "qwe44444", false}, {`qwe\\5`, `qwe\\\\\`, false}}
	runTests(t, "Unpack escape naive", UnpackString, cases)
	runTests(t, "Unpack escape states", UnpackStringState, cases)
}

func TestErrors(t *testing.T) {
	cases := []TestCase{{`45`, "", true}, {`\`, "", true}, {`\\\`, ``, true}}
	runTests(t, "Unpack errors naive", UnpackString, cases)
	runTests(t, "Unpack errors states", UnpackStringState, cases)
}

var longStringToUnpack = `asdfghj123\6klxcvbnmo6`

func BenchmarkNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := UnpackString(longStringToUnpack)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkState(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := UnpackStringState(longStringToUnpack)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkNaive2(b *testing.B) {
	rle := &RLEDecoder{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rle.NewString(longStringToUnpack)
		for {
			step, err := rle.Step()
			if err != nil {
				b.Error(err)
			}
			if !step {
				break
			}
		}
	}
}

func BenchmarkState2(b *testing.B) {
	fsm := newFSM()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := fsm.Unpack(longStringToUnpack)
		if err != nil {
			b.Error(err)
		}
	}
}
