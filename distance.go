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

// createDistanceMatrixURL to construct the API endpoint with its params.
func createDistanceMatrixURL(origins, destinations []string, key string) (*url.URL, error) {
	endpoint, err := url.Parse("https://maps.googleapis.com/maps/api/distancematrix/json")
	if err != nil {
		return nil, err
	}
	endpoint.Scheme = "https"
	endpoint.Host = "google.com"

	query := endpoint.Query()

	for _, origin := range origins {
		query.Add("origins", origin)
	}

	for _, destination := range destinations {
		query.Add("destinations", destination)
	}

	if key != "" {
		query.Add("key", key)
	}

	endpoint.RawQuery = query.Encode()

	return endpoint, nil
}

// distanceWrapper will wrap raw response into DistanceMatrix struct.
func distanceWrapper(response distanceMatrixAPIResponse) ([]DistanceMatrix) {
	var result []DistanceMatrix
	for originKey, row := range response.Rows {
		for destKey, element := range row.Elements {
			var matrix DistanceMatrix
			matrix.Origin = response.OriginAddresses[originKey]
			matrix.Destination = response.DestinationAddresses[destKey]
			matrix.Distance = element.Distance.Value
			matrix.Duration = element.Duration.Value
			result = append(result, matrix)
		}
	}
	return result
}

// CalculateDistance will hit Google API and the Wrap it into DistanceMatrix Array.
// The size of the array is the size of cross product OriginsXDestinations.
func (m Malkist) CalculateDistance(origins, destinations []string) ([]DistanceMatrix, error) {
	endpoint, err := createDistanceMatrixURL(origins, destinations, m.Key)
	if err != nil {
		return nil, fmt.Errorf("URL error: %v", err.Error())
	}

	res, err := http.Get(endpoint.String())
	if err != nil {
		return nil, fmt.Errorf("distance calculation error: %v", err.Error())
	}

	var body distanceMatrixAPIResponse
	json.NewDecoder(res.Body).Decode(&body)

	defer res.Body.Close()

	if body.Status != "OK" {
		return nil, fmt.Errorf("distance calculation error: %v", body.Status)
	}

	return distanceWrapper(body), nil
}
