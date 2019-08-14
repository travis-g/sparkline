package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mitchellh/colorstring"
)

func main() {
	conf := MergeBarConfig(DefaultBarConfig(), ConfigFromEnvironment())

	// Default bar to 0%
	val := 0.0

	// Read arguments or stdin
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Scanf("%f", &val)
	} else {
		var err error
		val, err = strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		}
	}
	if val > 100 {
		val = 100
	}

	fmt.Printf("%s\n", colorstring.Color(BarString(val, conf)))
}
