package main

import "testing"

// package-level var to eliminate compiler optimizations
var result interface{}

func quote(s string) string {
	return "\"" + s + "\""
}

func TestString(t *testing.T) {
	var r string

	// Make a testing bar that's easy to see which characters
	// are which parts.
	//
	// Colored output is also disabled here to remove potential
	// interference.
	conf := BarConfig{
		Size:      20,
		Char1:     "1",
		Char2:     "2",
		Separator: "s",
		Start:     "[",
		End:       "]",
	}

	testCases := []struct {
		str   string
		value float64
	}{
		{"[s2222222222222222222]", 0},
		{"[s2222222222222222222]", 1},
		{"[1111111111s222222222]", 50},
		{"[1111111111111111111s]", 99},
		{"[11111111111111111111]", 100},
	}

	for _, tc := range testCases {
		r = BarString(tc.value, &conf)
		if r != tc.str {
			t.Errorf("bar from %f was incorrect, got: %s, want: %s", tc.value, r, tc.str)
		}
	}
	result = r
}

// TestLength0 is to test if a 0-length bar with any value will render properly.
func TestLength0(t *testing.T) {
	var r string

	conf := BarConfig{
		Size:      0,
		Char1:     "1",
		Char2:     "2",
		Separator: "s",
		Start:     "[",
		End:       "]",
	}

	testCases := []struct {
		str   string
		value float64
	}{
		{"[]", 0},
		{"[]", 50},
		{"[]", 100},
	}
	for _, tc := range testCases {
		r = BarString(tc.value, &conf)
		if r != tc.str {
			t.Errorf("0-length bar for %f was incorrect, got: %s, want: %s", tc.value, r, tc.str)
		}
	}
	result = r
}

// TestLength1 tests that bar rending complies with how a bar should render
// based on what each bar component represents. If this test fails something is
// amiss with how the bar is stringified.
func TestLength1(t *testing.T) {
	var r string

	conf := BarConfig{
		Size:      1,
		Char1:     "1",
		Char2:     "2",
		Separator: "s",
		Start:     "[",
		End:       "]",
	}

	testCases := []struct {
		str   string
		value float64
	}{
		{"[s]", 0},
		{"[s]", 1},
		{"[s]", 50},
		{"[s]", 99},
		{"[1]", 100},
	}
	for _, tc := range testCases {
		r = BarString(tc.value, &conf)
		if r != tc.str {
			t.Errorf("1-length bar for %f was incorrect, got: %s, want: %s", tc.value, r, tc.str)
		}
	}
	result = r
}

func BenchmarkBar0(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = SimpleBar(0)
	}
	result = r
}

func BenchmarkBar50(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = SimpleBar(50)
	}
	result = r
}

func BenchmarkBar100(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = SimpleBar(100)
	}
	result = r
}

func BenchmarkStringify0(b *testing.B) {
	var r string
	// create bar
	conf := BarConfig{
		Char1:     "1",
		Char2:     "2",
		Separator: "s",
		Start:     "",
		End:       "",
	}
	for n := 0; n < b.N; n++ {
		// record the output to ensure compiler doesn't
		// eliminate the function call
		r = BarString(0, &conf)
	}
	// record to package-level variable
	result = r
}

func BenchmarkStringify50(b *testing.B) {
	var r string
	conf := BarConfig{
		Char1:     "1",
		Char2:     "2",
		Separator: "s",
		Start:     "",
		End:       "",
	}
	for n := 0; n < b.N; n++ {
		r = BarString(50, &conf)
	}
	result = r
}

func BenchmarkStringify100(b *testing.B) {
	var r string
	conf := BarConfig{
		Char1:     "1",
		Char2:     "2",
		Separator: "s",
		Start:     "",
		End:       "",
	}
	for n := 0; n < b.N; n++ {
		r = BarString(100, &conf)
	}
	result = r
}
