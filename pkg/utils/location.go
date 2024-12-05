package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/jizambrana5/quickfix-back/pkg/domain"

	"github.com/jizambrana5/quickfix-back/pkg/lib/errors"
)

type Locations struct {
	Locations map[string][]string `json:"locations"`
}

var locations Locations

func LoadLocations() (Locations, error) {
	var loc Locations

	// Obtener la ruta absoluta del archivo mendoza_locations.json dentro del contenedor
	filePath := filepath.Join("utils", "mendoza_locations.json")

	// Abrir el archivo JSON
	file, err := os.Open(filePath)
	if err != nil {
		return loc, err
	}
	defer file.Close()

	// Decodificar el archivo JSON en la estructura Locations
	err = json.NewDecoder(file).Decode(&loc)
	if err != nil {
		return loc, err
	}

	locations = loc
	return loc, nil
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

func GetLocations() (Locations, error) {
	if locations.Locations == nil {
		return Locations{}, errors.ErrInvalidLocation
	}
	return locations, nil
}

func ValidateProfession(profession string) bool {
	switch strings.ToUpper(profession) {
	case string(domain.Limpieza):
		return true
	default:
		return false
	}
}
