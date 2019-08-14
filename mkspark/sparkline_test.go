package main

import (
	"testing"
)

func quote(s string) string {
	return "\"" + s + "\""
}

func TestSpark(t *testing.T) {
	testSeriesCases := []struct {
		data   []float64
		output string
	}{
		{[]float64{}, ""},
		{[]float64{0}, "▅"},
		{[]float64{0, 0}, "▅▅"},
		{[]float64{0, 1}, "▁█"},
		{[]float64{-1, 0}, "▁█"},
		{[]float64{1, 1}, "▅▅"},
		{[]float64{1, 2}, "▁█"},
		{[]float64{0, 1, 2}, "▁▄█"},
		{[]float64{0, 1, 2, 3}, "▁▃▅█"},
		{[]float64{0, 1, 2, 3, 5, 8, 13}, "▁▁▂▂▃▅█"},
	}
	for _, tc := range testSeriesCases {
		s := SimpleSpark(tc.data...)
		if s != tc.output {
			t.Errorf("received %v; want %s, got %s", tc.data, quote(tc.output), quote(s))
		}
	}
}

func BenchmarkNewSparkline(b *testing.B) {
	var s *Sparkline
	data := []float64{0, 1, 2, 3, 4}
	for n := 0; n < b.N; n++ {
		s = NewSparkline(data...)
	}
	result = &s
}

func benchmarkSpark(i []float64, b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		s = SimpleSpark(i...)
	}
	result = s
}

func BenchmarkSparkConstant(b *testing.B) { benchmarkSpark([]float64{0, 0, 0}, b) }

func BenchmarkSparkVaried(b *testing.B) { benchmarkSpark([]float64{0, 1, 0}, b) }
