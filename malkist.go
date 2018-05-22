package malkist_go

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DistanceMatrix struct {
	Origin      string  `json:"origin"`
	Destination string  `json:"destination"`
	Distance    float64 `json:"distance"`
	Duration    float64 `json:"duration"`
}
type DistanceMatrixAPIResponse struct {
	DestinationAddresses []string                    `json:"destination_addresses"`
	OriginAddresses      []string                    `json:"origin_addresses"`
	Rows                 []DistanceMatrixAPIElements `json:"rows"`
	Status               string                      `json:"status"`
}
type DistanceMatrixAPIElements struct {
	Elements []DistanceMatrixAPIElement `json:"elements"`
}

type DistanceMatrixAPIElement struct {
	Distance DistanceMatrixAPIDistance `json:"distance"`
	Duration DistanceMatrixAPIDuration `json:"duration"`
	Status               string                      `json:"status"`
}
type DistanceMatrixAPIDistance struct {
	Text  string  `json:"text"`
	Value float64 `json:"value"`
}
type DistanceMatrixAPIDuration struct {
	Text  string  `json:"text"`
	Value float64 `json:"value"`
}

func CalculateDistance(origins, destinations []string) ([]DistanceMatrix, error) {
	var result []DistanceMatrix

	query := fmt.Sprint("https://maps.googleapis.com/maps/api/distancematrix/json?origins=")
	for key, origin := range origins {
		query += fmt.Sprintf("%v", origin)
		if key != len(origins) - 1 {
			query += "|"
		}
	}

	query += fmt.Sprint("&destinations=")

	for key, destination := range destinations {
		query += fmt.Sprintf("%v", destination)
		if key != len(destinations) - 1 {
			query += "|"
		}
	}
	res, err := http.Get(query)
	if err != nil {
		return nil, fmt.Errorf("Distance calculation error: %v", err.Error())
	}
	var body DistanceMatrixAPIResponse
	json.NewDecoder(res.Body).Decode(&body)
	if body.Status != "OK" {
		return nil, fmt.Errorf("Distance calculation error: %v", body.Status)
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
