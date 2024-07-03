package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jizambrana5/quickfix-back/pkg/lib/errors"
)

type Locations struct {
	Locations map[string][]string `json:"locations"`
}

func LoadLocations() (Locations, error) {
	// Construct the absolute path to mendoza_locations.json
	// Adjust this path as needed based on your project structure

	// Open the JSON file
	file, err := os.Open("./pkg/utils/mendoza_locations.json")
	if err != nil {
		return Locations{}, fmt.Errorf("error opening JSON file: %w", err)
	}
	defer file.Close()

	// Decode JSON into Locations struct
	var locations Locations
	err = json.NewDecoder(file).Decode(&locations)
	if err != nil {
		return Locations{}, fmt.Errorf("error decoding JSON file: %w", err)
	}

	return locations, nil
}

func ValidateLocation(department, district string, locations Locations) error {
	districts, ok := locations.Locations[department]
	if !ok {
		return errors.ErrInvalidDepartment
	}

	for _, d := range districts {
		if d == district {
			return nil
		}
	}

	return errors.ErrInvalidDistrict
}
