package repo

import (
	"github.com/FrankSantoso/geo-golang"
)

type GeoFetcher interface{
	GetGeocode(geo.Geocoder)
}
