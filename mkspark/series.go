package main

import (
	"github.com/mitchellh/mapstructure"
)

type Series struct {
	Data []float64 `mapstructure:"data"`
}

func Decode(input map[string]interface{}, data *Series) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &data,
	})
	if err != nil {
		panic(err)
	}

	return decoder.Decode(input)
}

// NewSeries creates a new data series.
func NewSeries(data ...float64) *Series {
	s := &Series{
		Data: data,
	}
	return s
}
