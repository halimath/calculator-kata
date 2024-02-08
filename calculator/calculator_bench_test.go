package calculator

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func BenchmarkSuperSimple(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := Eval(strings.NewReader("1 + 2 + 3"))
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSimple(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := Eval(strings.NewReader("38034 - 172.432 * 16864 / 45030 - 162 / (663.45532 * 535)"))
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark1k(b *testing.B) {
	benchmarkFile("../testdata/1k", b)
}

func Benchmark10k(b *testing.B) {
	benchmarkFile("../testdata/10k", b)
}

func Benchmark1m(b *testing.B) {
	benchmarkFile("../testdata/1m", b)
}

func Benchmark10m(b *testing.B) {
	benchmarkFile("../testdata/10m", b)
}

func benchmarkFile(filename string, b *testing.B) {
	content, err := os.ReadFile(filename)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := Eval(bytes.NewReader(content))
		if err != nil {
			b.Fatal(err)
		}
	}

}
