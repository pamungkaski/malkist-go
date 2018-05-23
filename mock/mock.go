package mock

import "github.com/pamungkaski/malkist-go"

type distanceMatrixMock struct {
	Origins      []string
	Destinations []string
	Expected     []malkist.DistanceMatrix
}

// DistanceMatrixMocks contain the scenario for testing CalculateDistance function.
var DistanceMatrixMocks = []distanceMatrixMock{
	{
		Expected: []malkist.DistanceMatrix{
			{
				Origin:      "566 Vermont St, Brooklyn, NY 11207, USA",
				Destination: "67-89 Pacific St, Brooklyn, NY 11201, USA",
				Distance:    float64(10423),
				Duration:    float64(2062),
			},
		},
		Origins: []string{
			"40.6655101, -73.89188969999998",
		},
		Destinations: []string{
			"40.6905615, -73.9976592",
		},
	},
}
