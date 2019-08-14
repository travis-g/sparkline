package main

// func TestSeriesAppend(t *testing.T) {
// 	testCases := []struct {
// 		data   []float64
// 		add    float64
// 		result []float64
// 	}{
// 		{[]float64{}, 1, []float64{1}},
// 		{[]float64{1}, 1, []float64{1, 1}},
// 	}
// 	for _, tc := range testCases {
// 		series := sparkline.NewSeries(tc.data...)
// 		series.Append(tc.add)
// 		if !reflect.DeepEqual(spark.Data, tc.result) {
// 			t.Errorf("added %v; want %v, got %v", tc.add, tc.result, spark.Data)
// 		}
// 	}
// }

// func TestPop(t *testing.T) {
// 	testCases := []struct {
// 		data   []float64
// 		result []float64
// 	}{
// 		{[]float64{}, []float64{}},
// 		{[]float64{1}, []float64{}},
// 		{[]float64{1, 2}, []float64{2}},
// 	}
// 	for _, tc := range testCases {
// 		spark := sparkline.NewSparkline(tc.data...)
// 		spark.Pop()
// 		if !reflect.DeepEqual(spark.Data, tc.result) {
// 			t.Errorf("popped %v; want %v, got %v", tc.data, tc.result, spark.Data)
// 		}
// 	}
// }
