package main

// A SeriesConfig is a configuration for displaying a series as a sparkline.
// Configs should specify either the Bins OR the Span argument. If both Bins and
// Span are provided Span will be used.
type SeriesConfig struct {
	// Start and End values of the series. Values that lay outside of [start,
	// end] are discarded. If unspecified, Start and End will be extrapolated
	// from the input data prior to discretizing.
	Start *float64
	End   *float64

	// Number of bins to discretize data into. Bins must be non-negative. If
	// Span is nil, Bins will effectively be the length of the sparklines
	// created using the config.
	Bins int

	// Span is the size of each bin. If non-nil, Span will take precident over
	// Bins. If nil, Span will be calculated for each series it is used to
	// render.
	Span *float64
}

// DefaultSeriesConfig returns a default SeriesConfig.
func DefaultSeriesConfig() *SeriesConfig {
	return &SeriesConfig{
		Start: nil,
		End:   nil,
		Bins:  20,
		Span:  nil,
	}
}

// MergeSeriesConfig returns the result of two SeriesConfig objects merged
// together. Precedence is set from right to left. MergeSeriesConfig should not
// check for config validity.
func MergeSeriesConfig(a, b *SeriesConfig) (c *SeriesConfig) {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}

	// duplicate the data of config a into c, to then merge config b onto
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
