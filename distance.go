package malkist

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// DistanceMatrix represent the value that returned from CalculateDistance.
type DistanceMatrix struct {
	Origin      string  `json:"origin"`
	Destination string  `json:"destination"`
	Distance    float64 `json:"distance"`
	Duration    float64 `json:"duration"`
}

// DistanceMatrixAPIResponse is the response of the request from Google Maps Web Service API.
type distanceMatrixAPIResponse struct {
	DestinationAddresses []string                    `json:"destination_addresses"`
	OriginAddresses      []string                    `json:"origin_addresses"`
	Rows                 []distanceMatrixAPIElements `json:"rows"`
	Status               string                      `json:"status"`
}

// distanceMatrixAPIElements represent the elements array on the API response.
type distanceMatrixAPIElements struct {
	Elements []distanceMatrixAPIElement `json:"elements"`
}

// distanceMatrixAPIElement represent each element based on the cross product of OriginsXDestinations
// on the API response.
type distanceMatrixAPIElement struct {
	Distance distanceMatrixAPIDistance `json:"distance"`
	Duration distanceMatrixAPIDuration `json:"duration"`
	Status   string                    `json:"status"`
}

// distanceMatrixAPIDistance is the Distance value from the Element.
type distanceMatrixAPIDistance struct {
	Text  string  `json:"text"`
	Value float64 `json:"value"`
}

// distanceMatrixAPIDuration is the Duration value from the Element.
type distanceMatrixAPIDuration struct {
	Text  string  `json:"text"`
	Value float64 `json:"value"`
}

// CalculateDistance will hit Google API and the Wrap it into DistanceMatrix Array.
// The size of the array is the size of cross product OriginsXDestinations.
func (m Malkist) CalculateDistance(origins, destinations []string) ([]DistanceMatrix, error) {
	var result []DistanceMatrix

	query := fmt.Sprint("https://maps.googleapis.com/maps/api/distancematrix/json?origins=")
	for key, origin := range origins {
		query += url.QueryEscape(fmt.Sprintf("%v", origin))
		if key != len(origins)-1 {
			query += "%7C"
		}
	}

	query += fmt.Sprint("&destinations=")

	for key, destination := range destinations {
		query += url.QueryEscape(fmt.Sprintf("%v", destination))
		if key != len(destinations)-1 {
			query += "%7C"
		}
	}

	if m.Key != "" {
		query += fmt.Sprintf("&key=%v", m.Key)
	}

	//fmt.Println(query)
	res, err := http.Get(query)
	if err != nil {
		return nil, fmt.Errorf("distance calculation error: %v", err.Error())
	}

	var body distanceMatrixAPIResponse
	json.NewDecoder(res.Body).Decode(&body)

	if body.Status != "OK" {
		return nil, fmt.Errorf("distance calculation error: %v", body.Status)
	}

	for originKey, row := range body.Rows {
		for destKey, element := range row.Elements {
			var matrix DistanceMatrix
			matrix.Origin = body.OriginAddresses[originKey]
			matrix.Destination = body.DestinationAddresses[destKey]
			matrix.Distance = element.Distance.Value
			matrix.Duration = element.Duration.Value
			result = append(result, matrix)
		}
	}

	return result, nil
}
