package main

import (
	"bytes"
	"fmt"
	"math"
)

// A Sparkline is a glorified array of integers
type Sparkline struct {
	Data []float64
}

// Ticks are the ASCII symbols (runes) that make up the sparklines. Runes are
// used as these characters have a length greater than 1 when treated as
// strings.
var Ticks = []rune{
	'▁', // lowest relative value
	'▂',
	'▃',
	'▄',
	'▅',
	'▆',
	'▇',
	'█', // highest relative value
}

// NewSparkline creates a new Sparkline object
func NewSparkline(data ...float64) *Sparkline {
	s := &Sparkline{
		Data: data,
	}
	return s
}

// SimpleSpark will create a sparkline out of a variable number of integer
// datapoints.
func SimpleSpark(data ...float64) string {
	if len(data) == 0 {
		data = []float64{}
	}
	spark := Sparkline{
		Data: data,
	}
	return spark.String()
}

// String returns the string representation of a sparkline.
func (s *Sparkline) String() string {

	// Short-circuit if there's no data to render
	if len(s.Data) == 0 {
		return ""
	}

	var (
		line bytes.Buffer
		max  = -math.MaxFloat64
		min  = math.MaxFloat64
	)

	// Get the min and max values of the dataset by traversing it
	for _, i := range s.Data {
		if i > max {
			max = i
		}
		if i < min {
			min = i
		}
	}

	// check for constant data after all data has been traversed
	if max == min {
		t := Ticks[len(Ticks)/2]
		for n := 0; n < len(s.Data); n++ {
			fmt.Fprintf(&line, "%s", string(t))
		}
		return line.String()
	}

	// Determine interval size of the ranges to use when assigning tick heights
	f := ((max - min) / float64(len(Ticks)-1))

	// Range over the data, determine the height of the point in comparison to
	// the other values, then add the character to the buffer
	for _, x := range s.Data {
		i := int((x - min) / f)
		fmt.Fprintf(&line, "%s", string(Ticks[i]))
	}

	// return the buffer as a string
	return line.String()
}
