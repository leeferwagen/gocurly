package gocurly

import (
	"fmt"
	"testing"
)

const (
	tinyString         = "<{green>Tiny green text<}>"
	bigString          = "<{green>Green<}><{red>Red<}><{blue>Blue<}><{bold>Bold<}><{underline>Underlined<}>"
	nestedString       = "<{green>Green<{red>Red<}><{blue>Blue<}><{bold>Bold<}><{underline>Underlined<}><}>"
	deeplyNestedString = "<{green>Green<{red>Red<{blue>Blue<{bold>Bold<{underline>Underlined<}><}><}><}><}>"
)

func TestTiny(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
		return
	}
	res := FormatString(tinyString)
	fmt.Printf("S: %q\n", res)
	fmt.Printf(">> %s\n", res)
}

func TestBig(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
		return
	}
	res := FormatString(bigString)
	fmt.Printf("S: %q\n", res)
	fmt.Printf(">> %s\n", res)
}

func TestNested(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
		return
	}
	res := FormatString(nestedString)
	fmt.Printf("S: %q\n", res)
	fmt.Printf(">> %s\n", res)
}

func TestDeeplyNested(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
		return
	}
	res := FormatString(deeplyNestedString)
	fmt.Printf("S: %q\n", res)
	fmt.Printf(">> %s\n", res)
}

// Benchmark tests
func BenchmarkTiny(b *testing.B) { // Tiny
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		FormatString(tinyString)
	}
}

func BenchmarkBig(b *testing.B) { // Big
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		FormatString(bigString)
	}
}

func BenchmarkNested(b *testing.B) { // Nested
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		FormatString(nestedString)
	}
}

func BenchmarkDeeplyNested(b *testing.B) { // Deeply Nested
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		FormatString(deeplyNestedString)
	}
}
