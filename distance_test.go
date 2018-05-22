package malkist_test

import (
	"github.com/pamungkaski/malkist-go"
	"testing"
)

func TestCalculateDistance(t *testing.T) {
	var mocks = malkist.DistanceMatrixMocks
	m := malkist.Malkist{}
	for _, mock := range mocks {
		result, err := m.CalculateDistance(mock.Origins, mock.Destinations)
		if err != nil {
			t.Error(err)
			return
		}

		expected := mock.Expected
		if len(expected) != len(result) {
			t.Error(err)
			return
		}
		for key, res := range result {
			if expected[key].Duration != res.Duration {
				t.Errorf("%v to %v duration expected %v got %v", res.Origin, res.Destination, expected[key].Duration, res.Duration)
				return
			}
			if expected[key].Distance != res.Distance {
				t.Errorf("%v to %v distance expected %v got %v", res.Origin, res.Destination, expected[key].Distance, res.Distance)
				return
			}
		}
	}
}
