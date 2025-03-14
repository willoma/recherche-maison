package ville

import "errors"

// Custom errors for the city service
var (
	// ErrCityInUse is returned when attempting to delete a city that is used by houses
	ErrCityInUse = errors.New("la ville est utilisée par des maisons et ne peut pas être supprimée")
)
