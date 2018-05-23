package malkist_test

import (
	"github.com/pamungkaski/malkist-go"
	"testing"
	"github.com/pamungkaski/malkist-go/mock"
)

func TestCalculateDistance(t *testing.T) {
	var mocks = mock.DistanceMatrixMocks
	m := malkist.Malkist{}
	for _, malmock := range mocks {
		result, err := m.CalculateDistance(malmock.Origins, malmock.Destinations)
		if err != nil {
			t.Fatal(err)
		}

		expected := malmock.Expected
		if len(expected) != len(result) {
			t.Fatal(err)
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
