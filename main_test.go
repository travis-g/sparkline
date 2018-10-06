package main

import (
	"log"
	"os/exec"
	"reflect"
	"testing"
)

func quote(s string) string {
	return "\"" + s + "\""
}

// generic interface to prevent compiler optimizations
var result interface{}

func TestString(t *testing.T) {
	testCases := []struct {
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
	for _, tc := range testCases {
		spark := NewSparkline(tc.data...)
		s := spark.String()
		if s != tc.output {
			t.Errorf("received %v; want %s, got %s", tc.data, quote(tc.output), quote(s))
		}
	}
}

func TestAppend(t *testing.T) {
	testCases := []struct {
		data   []float64
		add    float64
		result []float64
	}{
		{[]float64{}, 1, []float64{1}},
		{[]float64{1}, 1, []float64{1, 1}},
	}
	for _, tc := range testCases {
		spark := NewSparkline(tc.data...)
		spark.Append(tc.add)
		if !reflect.DeepEqual(spark.Data, tc.result) {
			t.Errorf("added %v; want %v, got %v", tc.add, tc.result, spark.Data)
		}
	}
}

func TestPop(t *testing.T) {
	testCases := []struct {
		data   []float64
		result []float64
	}{
		{[]float64{}, []float64{}},
		{[]float64{1}, []float64{}},
		{[]float64{1, 2}, []float64{2}},
	}
	for _, tc := range testCases {
		spark := NewSparkline(tc.data...)
		spark.Pop()
		if !reflect.DeepEqual(spark.Data, tc.result) {
			t.Errorf("popped %v; want %v, got %v", tc.data, tc.result, spark.Data)
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
		s = Spark(i...)
	}
	result = s
}

func BenchmarkSparkConstant(b *testing.B) { benchmarkSpark([]float64{0, 0, 0}, b) }

func BenchmarkSparkVaried(b *testing.B) { benchmarkSpark([]float64{0, 1, 0}, b) }

func benchmarkExec(cmd string, b *testing.B) {
	var (
		out  []byte
		err  error
		data = []string{"0", "1", "2", "8", "9", "30", "3", "18", "1", "3", "10"}
	)
	for n := 0; n < b.N; n++ {
		out, err = exec.Command(cmd, data...).Output()
		if err != nil {
			log.Fatalf("cmd failed: %s\n", err)
		}
	}
	result = out
}

func BenchmarkShell(b *testing.B) {
	benchmarkExec("./sparkline", b)
}
