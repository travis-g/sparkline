package main

// A Config is a configuration for displaying a series as a sparkline. Configs
// should specify either the Bins OR the Span argument. If both Bins
// and Span are provided Span will be used.
type Config struct {
	// Start and End values of the series. Values that lay outside of [start,
	// end] are discarded. If unspecified, Start and End will be extrapolated
	// from the input data prior to discretizing.
	Start *int
	End   *int

	// Number of bins to discretize data into. Bins must be non-negative. If
	// Span is nil, Bins will effectively be the length of the sparklines
	// created using the Config.
	Bins int

	// Span is the size of each bin. If non-nil, Span will take precident over
	// Bins. If nil, Span will be calculated for each series it is used to
	// render.
	Span *float64
}

// DefaultConfig returns a default Config.
func DefaultConfig() *Config {
	return &Config{
		Start: nil,
		End:   nil,
		Bins:  20,
		Span:  nil,
	}
}

// MergeConfig returns the result of two Config objects merged together.
// Precedence is set from right to left. MergeConfig should not check for Config
// validity.
func MergeConfig(a, b *Config) (c *Config) {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}

	// duplicate the data of Config a into c, to then merge Config b onto
	*c = *a

	if b.Bins > 0 {
		c.Bins = b.Bins
	}

	if b.Start != nil {
		c.Start = b.Start
	}

	if b.End != nil {
		c.End = b.End
	}

	if b.Span != nil {
		c.Span = b.Span
	}

	return
}
