package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
)

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

// A Sparkline is a glorified array of integers
type Sparkline struct {
	Data []float64
}

// NewSparkline creates a new Sparkline object
func NewSparkline(data ...float64) *Sparkline {
	s := &Sparkline{
		Data: data,
	}
	return s
}

// Length returns the length of []Data, i.e. the number of data points in the
// sparkline.
func (s *Sparkline) Length() int {
	return len(s.Data)
}

// Pop removes a data point from the begining of the sparkline, shorting the
// total length by 1.
func (s *Sparkline) Pop() {
	if s.Length() == 0 {
		return
	}
	s.Data = s.Data[1:]
	return
}

// Append adds data to the sparkline
func (s *Sparkline) Append(data ...float64) {
	s.Data = append(s.Data, data...)
}

// Spark will create a Sparkline out of a variable number of integer datapoints.
func Spark(data ...float64) string {
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

	// Get the min and max values of the dataset
	for _, i := range s.Data {
		if i > max {
			max = i
		}
		if i < min {
			min = i
		}
	}

	// If history is constant use a mid-sized tick
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

func main() {
	// _args will be the faux arg list that we either read CLI arguments into or
	// craft based on what's piped through os.Stdin.
	var _args []string

	// TODO: flags?

	pipe, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	// If we received bytes through stdin, use those.
	if pipe.Mode()&os.ModeCharDevice == 0 {
		// Save what's in the stdin buffer since buffers are consumed when read.
		reader := bufio.NewReader(os.Stdin)
		var input bytes.Buffer
		for {
			char, _, err := reader.ReadRune()
			if err != nil && err == io.EOF {
				break
			}
			fmt.Fprintf(&input, "%c", char)
		}

		// split stdin based on what's not whitespace or commas
		_args = regexp.MustCompile("[^\\s,]+").FindAllString(input.String(), -1)
	} else {
		_args = os.Args[1:]
	}

	// Convert arguments from strings to floats and save them. If there's a
	// problem converting the strings we'll print an error but continue.
	var data []float64
	for _, i := range _args {
		j, err := strconv.ParseFloat(i, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			continue
		}
		data = append(data, j)
	}

	fmt.Printf("%s\n", Spark(data...))
}
