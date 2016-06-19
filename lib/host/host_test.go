package host

import (
	"strings"
	"testing"
)

type capitalizeTPair struct {
	in  string
	out string
}

func TestCapitalize(t *testing.T) {
	testPairs := []capitalizeTPair{
		{"foo", "Foo"},
		{"foo bar", "Foo bar"},
		{"", ""},
		{"Bar", "Bar"},
	}
	for _, testPair := range testPairs {
		out := capitalize(testPair.in)
		t.Logf("input: '%v'; out:'%v'; exp: '%v'", testPair.in, out,
			testPair.out)
		if strings.Compare(out, testPair.out) != 0 {
			t.Error()
		}
	}
}

type garchTPair struct {
	in  string
	out string
}

func TestGuessArch(t *testing.T) {
	testPairs := []garchTPair{
		{"x86_64", "amd64"},
		{"foo", "unknown"},
		{"", "unknown"},
	}
	for _, testPair := range testPairs {
		out := guessArch(testPair.in)
		t.Logf("input: '%v'; out:'%v'; exp: '%v'", testPair.in, out,
			testPair.out)
		if strings.Compare(out, testPair.out) != 0 {
			t.Error()
		}
	}
}
