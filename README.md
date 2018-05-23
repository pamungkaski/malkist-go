# Malkist

[![Build Status](https://travis-ci.com/pamungkaski/malkist-go.svg?branch=master)](https://travis-ci.com/pamungkaski/malkist-go)

Simple Google Maps Web Service API Wrapper on Go [Google Maps Web Service API](https://developers.google.com/maps/documentation/)

## Installation

    $ go get github.com/pamungkaski/malkist-go

## Functionality

* [x] Distance Matrix API [link](https://developers.google.com/maps/documentation/distance-matrix/intro)
* [ ] Elevation API [link](https://developers.google.com/maps/documentation/elevation/intro)
* [ ] Geolocation API [link](https://developers.google.com/maps/documentation/geolocation/intro)
* [ ] Time Zone API [link](https://developers.google.com/maps/documentation/timezone/intro)
* [ ] Directions API [link](https://developers.google.com/maps/documentation/directions/intro)
* [ ] Roads API [link](https://developers.google.com/maps/documentation/roads/intro)
* [ ] Geocoding [link](https://developers.google.com/maps/documentation/geocoding/intro)
* [ ] Place API for Web [link](https://developers.google.com/maps/documentation/web-service/intro)

## Usage
You'll need Google API key to use malkist, so checkout [get api key](https://developers.google.com/places/web-service/get-api-key)

### Initialize

```go
package main

import (
	"github.com/pamungkaski/malkist-go"
	"fmt"
)

func main() {
	m := malkist.Malkist{
		Key: "YOUR API KEY",
	}
	// to changeAPI key
	m.ChangeAPIKey("YOUR API KEY")
}

```
### Distance Matrix API
**Main Function :**

````
// distance.go
func (m Malkist) CalculateDistance(origins, destinations []string) ([]DistanceMatrix, error)
````
   
**Return Struct :**
```go
// distance.go
package malkist

// DistanceMatrix represent the value that returned from CalculateDistance.
type DistanceMatrix struct {
	Origin      string  `json:"origin"`
	Destination string  `json:"destination"`
	Distance    float64 `json:"distance"`
	Duration    float64 `json:"duration"`
}
```
`CalculateDistance` will return array of DistanceMatrix even when it is only between two coordinate. 

**Calculate distance between two coordinate :**
```go
package main

import (
	"github.com/pamungkaski/malkist-go"
	"fmt"
)

func main() {
	m := malkist.Malkist{}

	distance, err := m.CalculateDistance([]string{"40.6655101, -73.89188969999998"}, []string{"40.6905615, -73.9976592",})
	if err != nil {
		fmt.Errorf(err.Error())
	}

	fmt.Println(distance)
}
```

**Calculate distance between multiple coordinate :**
```go
package main

import (
	"github.com/pamungkaski/malkist-go"
	"fmt"
)

func main() {
	m := malkist.Malkist{}

	distance, err := m.CalculateDistance(
		[]string{"40.6655101, -73.89188969999998", "40.6905615, -73.9976592",},
		[]string{"40.6905615, -73.9976592", "40.6655101, -73.89188969999998",},
		)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	fmt.Println(distance)
}
```

## Development

After checking out the repo, run

    $ dep ensure

To install the dependencies of this library

### Testing

Run:

    $ go test -v

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/pamungkaski/malkist-go. This project is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [Contributor Covenant](http://contributor-covenant.org) code of conduct.

## License

The gem is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).

## Code of Conduct

Everyone interacting in the Malkist project’s codebases, issue trackers, chat rooms and mailing lists is expected to follow the [code of conduct](https://github.com/pamungkaski/Malkist-Ruby/blob/master/CODE_OF_CONDUCT.md).
