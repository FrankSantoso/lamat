package repo

import (
	"fmt"
	"github.com/FrankSantoso/lamat.git/internal/cfg"
	"github.com/codingsince1985/geo-golang/chained"
	"github.com/codingsince1985/geo-golang/google"
	"github.com/codingsince1985/geo-golang/mapquest/nominatim"
	"github.com/codingsince1985/geo-golang/openstreetmap"
	"io"
	"os"
	"reflect"
	"strconv"
	"text/tabwriter"

	"github.com/codingsince1985/geo-golang"
	"github.com/fatih/color"
	"github.com/fatih/structs"
)

// GeoGetter interface outputs formatted
type GeoGetter interface {
	GetGeocode() error
	GetReverseGeocode() error
}

// Geo struct represents geocoder struct
type Geo struct {
	geocoder geo.Geocoder
	args     []string
}

func NewGeo(conf *cfg.Config, args []string) *Geo {
	geocoder := chained.Geocoder(
		google.Geocoder(conf.APIKeys.GoogleGeocode),
		nominatim.Geocoder(conf.APIKeys.Nominatim),
		openstreetmap.Geocoder(),
	)
	return &Geo{
		geocoder, args,
	}
}

var (
	errColor    = color.New(color.FgRed).SprintFunc()
	infoColor   = color.New(color.FgBlue).SprintFunc()
	stringColor = color.New(color.FgGreen).SprintFunc()
	warnColor   = color.New(color.FgYellow).SprintFunc()
)

func getTabWriterOutput() *tabwriter.Writer {
	return tabwriter.NewWriter(os.Stdout, 12, 2, 4, ' ', tabwriter.TabIndent)
}

// GetGeocode location
func (g *Geo) GetGeocode() error {
	out := getTabWriterOutput()
	location, err := g.geocoder.Geocode(g.args[0])
	if err != nil {
		fmt.Fprintf(out, "\n%s\t%s\n", errColor("Error:"), err)
		return err
	}
	printRow(out, "\n%s\t%s\t%s\n", infoColor("Address"), g.args[0])
	locmap := structs.Map(location)
	if err = printRows(out, locmap); err != nil {
		return err
	}
	out.Flush()
	return nil
}

// GetReverseGeocode
func (g *Geo) GetReverseGeocode() error {
	out := getTabWriterOutput()
	args, err := strsToFloats(g.args)
	if err != nil {
		return err
	}
	gcode, err := g.geocoder.ReverseGeocode(args[0], args[1])
	if err != nil {
		fmt.Fprintf(out, "\n%s\t%s\n", errColor("Error:"), err)
		return err
	}
	jmap := structs.Map(gcode)
	if err = printRows(out, jmap); err != nil {
		return err
	}
	out.Flush()
	return nil
}

func printRows(out *tabwriter.Writer, rows map[string]interface{}) error {
	for k, v := range rows {
		t := reflect.TypeOf(v).Name()
		if t == "string" && v != "" {
			printRow(out, "\n%s\t%s\t%s", k, v)
		} else if t == "float64" {
			printRow(out, "\n%s\t%s\t%.6f", k, v)
		}
		printRow(out, "\n%s\t%s\t%v", k, v)
	}
	return nil
}

func printRow(out io.Writer, format, key string, value interface{}) {
	fmt.Fprintf(out, format, stringColor(key), stringColor(":"), value)
}

func strsToFloats(strs []string) ([]float64, error) {
	var f64s []float64
	for _, val := range strs {
		f, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return nil, err
		}
		f64s = append(f64s, f)
	}
	return f64s, nil
}
