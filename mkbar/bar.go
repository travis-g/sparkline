package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/caarlos0/env"
)

// BarConfig objects define all the necessary attributes of how to stringify a
// value. A rendered bar is made up by the following:
//
//            ├──────────{Size}───────────┤
//     {Start}[{Char1}]{Separator}[{Char2}]{End}
//
// A stringified bar can be made up of [0, Size] of Char1, 0 or 1 Separator, and
// [0, Size-1] of Char2. Exactly one instance of Start and End are always added,
// but they can be set to 0-length strings.
type BarConfig struct {
	// Size is the length of the bar's internal characters, not including Start
	// and End's lengths.
	Size int `env:"SIZE"`

	// Start is the bar's prefix
	Start string `env:"START"`

	// Char1 is the character used for the bar's active portion
	Char1 string `env:"CHAR1"`

	// TODO: "Glue"

	// Separator is the first character of the inactive section
	Separator string `env:"SEP"`

	// Char2 is the character used for the bar's inactive portion
	Char2 string `env:"CHAR2"`

	// End is the bar's suffix. It's advised to include a reset to the
	// [default] color
	End string `env:"END"`
}

// A Bar is just an integer value worthy of display.
type Bar int

// DefaultBarConfig is the default set of ASCII bar characters.
func DefaultBarConfig() *BarConfig {
	return &BarConfig{
		Size:      20,
		Start:     "",
		Char1:     "#",
		Separator: "[dark_gray]-",
		Char2:     "-",
		End:       "",
	}
}

// MergeBarConfig returns the result of two Config objects together and returns the
// result. Precedence is set from right to left.
func MergeBarConfig(a, b *BarConfig) *BarConfig {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}

	// Start with the left Config and merge the right overtop
	var merged = *a

	if b.Size != 0 {
		merged.Size = b.Size
	}
	if b.Start != "" {
		merged.Start = b.Start
	}
	if b.Char1 != "" {
		merged.Char1 = b.Char1
	}
	if b.Separator != "" {
		merged.Separator = b.Separator
	}
	if b.Char2 != "" {
		merged.Char2 = b.Char2
	}
	if b.End != "" {
		merged.End = b.End
	}

	return &merged
}

// ConfigFromEnvironment parses any available environment vars into a Config.
func ConfigFromEnvironment() *BarConfig {
	conf := BarConfig{}
	// parses any environment vars into a Config
	err := env.Parse(&conf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
	}

	return &conf
}

// SimpleBar returns a bar string using the default settings.
func SimpleBar(f float64) string {
	return BarString(f, DefaultBarConfig())
}

// BarString
func BarString(value float64, c *BarConfig) string {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "%s", c.Start)

	char1 := c.Char1
	char2 := c.Char2
	sep := c.Separator
	char := &char1

	// if value == 0 {
	// 	char = &char2
	// }

	for i := 0; i < c.Size; i++ {
		// 'length' is the number of bar segments that should be active (Char1)
		// based on the desired Size and value to represent
		length := int(value / 100 * float64(c.Size))
		if i >= length {
			// if previous char printed was Char1, add Separator
			if char == &char1 {
				char = &sep
			} else {
				char = &char2
			}
		}
		fmt.Fprintf(&buf, "%s", *char)
	}
	fmt.Fprintf(&buf, "%s", c.End)

	return buf.String()
}
