package repo

import (
	"fmt"
	"github.com/FrankSantoso/geo-golang"
	"github.com/fatih/color"
	"github.com/prometheus/common/log"
	"os"
	"text/tabwriter"
)

// GeoGetter interface outputs formatted
type GeoGetter interface {
	GetGeocode() error
	GetReverseGeocode() error
}

type Geo struct {
	Geocoder geo.Geocoder
	Args []string
}

var (
	red = color.New(color.Red)
)

func (g *Geo) GetGeocode() {
	out := tabwriter.NewWriter(os.Stdout, 8, 4, 4, ' ', tabwriter.TabIndent)
	location, err := g.Geocoder.Geocode(g.Args[0])
	if err != nil {
		log.Errorf("Error getting address: \t%v\n", err)
	}
	fmt.Fprintf(out, )
}

