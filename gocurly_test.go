package gocurly

import (
	"testing"
)

const (
	tinyString         = "<{green>Tiny green text<}>"
	bigString          = "<{green>Green<}><{red>Red<}><{blue>Blue<}><{bold>Bold<}><{underline>Underlined<}>"
	nestedString       = "<{green>Green<{red>Red<}><{blue>Blue<}><{bold>Bold<}><{underline>Underlined<}><}>"
	deeplyNestedString = "<{green>Green<{red>Red<{blue>Blue<{bold>Bold<{underline>Underlined<}><}><}><}><}>"
	textlessString     = "<{black>#<{red>#<{green>#<{yellow>#<{blue>#<{magenta>#<{cyan>#<{white>#<}><}><}><}><}><}><}><}>"
)

func TestBasic(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
		return
	}
	testStrings := map[string]string{
		tinyString:         "\x1b[32mTiny green text\x1b[39m",
		bigString:          "\x1b[32mGreen\x1b[39m\x1b[31mRed\x1b[39m\x1b[34mBlue\x1b[39m\x1b[1mBold\x1b[22m\x1b[4mUnderlined\x1b[24m",
		nestedString:       "\x1b[32mGreen\x1b[31mRed\x1b[39m\x1b[32m\x1b[34mBlue\x1b[39m\x1b[32m\x1b[1mBold\x1b[22m\x1b[32m\x1b[4mUnderlined\x1b[24m\x1b[32m\x1b[39m",
		deeplyNestedString: "\x1b[32mGreen\x1b[31mRed\x1b[34mBlue\x1b[1mBold\x1b[4mUnderlined\x1b[24m\x1b[1m\x1b[22m\x1b[34m\x1b[39m\x1b[31m\x1b[39m\x1b[32m\x1b[39m",
		textlessString:     "\x1b[30m#\x1b[31m#\x1b[32m#\x1b[33m#\x1b[34m#\x1b[35m#\x1b[36m#\x1b[37m#\x1b[39m\x1b[36m\x1b[39m\x1b[35m\x1b[39m\x1b[34m\x1b[39m\x1b[33m\x1b[39m\x1b[32m\x1b[39m\x1b[31m\x1b[39m\x1b[30m\x1b[39m",
	}
	testStringsOptimized := map[string]string{
		bigString: "\x1b[32mGreen\x1b[31mRed\x1b[34mBlue\x1b[39m\x1b[1mBold\x1b[22m\x1b[4mUnderlined\x1b[24m",
	}
	for str, expected := range testStrings {
		res := FormatString(str)
		if res != expected {
			t.Errorf("expected (raw): %s", expected)
			t.Errorf("expected (esc): %q", expected)
			t.Errorf("  result (raw): %s", res)
			t.Errorf("  result (esc): %q", res)
		}
	}
	for str, expected := range testStringsOptimized {
		res := OptimizeString(FormatString(str))
		if res != expected {
			t.Errorf("expected (raw): %s", expected)
			t.Errorf("expected (esc): %q", expected)
			t.Errorf("  result (raw): %s", res)
			t.Errorf("  result (esc): %q", res)
		}
	}
}

// Benchmark tests
func BenchmarkTiny1(b *testing.B) { // Tiny
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		FormatString(tinyString)
	}
}

func BenchmarkBig1(b *testing.B) { // Big
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		FormatString(bigString)
	}
}

func BenchmarkNested1(b *testing.B) { // Nested
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		FormatString(nestedString)
	}
}

func BenchmarkDeeplyNested1(b *testing.B) { // Deeply Nested
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		FormatString(deeplyNestedString)
	}
}

func BenchmarkTextless1(b *testing.B) { // Deeply Nested
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		FormatString(textlessString)
	}
}

func BenchmarkTextless2(b *testing.B) { // Deeply Nested
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		OptimizeString(FormatString(textlessString))
	}
}
