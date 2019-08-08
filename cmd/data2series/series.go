package main

import (
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
)

// A Series is an array of data to discretize.
type Series struct {
	Data []string `mapstructure:"data"`
}

// data that will be part of a series. Data must be numeric in nature, but can
// be weakly typed. See github.com/mitchellh/mapstructure
type series_data []float64

func TimeDurationToIntHook(
	f reflect.Type,
	t reflect.Type,
	data interface{}) (interface{}, error) {
	if f != reflect.TypeOf(time.Duration(5)) {
		return data, nil
	}
	if t.Kind() != reflect.Int64 {
		return data, nil
	}

	return int64(data.(time.Duration)), nil
}

func StringToTimeDurationHook(
	f reflect.Type,
	t reflect.Type,
	data interface{}) (interface{}, error) {
	if f.Kind() != reflect.String {
		return data, nil
	}
	if f != reflect.TypeOf(time.Duration(5)) {
		return data, nil
	}

	d, err := time.ParseDuration(data.(string))
	if err != nil {
		return data, nil
	}
	return d, err
}

var TimeTypesToIntHooks = mapstructure.ComposeDecodeHookFunc(
	StringToTimeDurationHook,
	TimeDurationToIntHook,
)

func Decode(input map[string]interface{}) *Series {
	var series Series
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: TimeTypesToIntHooks,
		// WeaklyTypedInput: true,
		Result: &series,
	})
	if err != nil {
		panic(err)
	}

	err = decoder.Decode(input)
	if err != nil {
		panic(err)
	}

	return &series
}

func String(series Series, c *Config) {
}
