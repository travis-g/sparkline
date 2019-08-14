package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

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

	fmt.Printf("%s\n", SimpleSpark(data...))
}
