package repo

import (
	"fmt"
	"github.com/FrankSantoso/gosm"
	"github.com/FrankSantoso/lamat/internal/cfg"
	"github.com/codingsince1985/geo-golang/chained"
	"github.com/codingsince1985/geo-golang/google"
	"github.com/codingsince1985/geo-golang/mapquest/nominatim"
	"io"
	"os"
	"reflect"
	"strconv"
	"text/tabwriter"

	"github.com/codingsince1985/geo-golang"
	"github.com/fatih/color"
)

// Geo struct represents geocoder struct
type Geo struct {
	geocoder geo.Geocoder
	args     []string
}

// gosmData
type gosmData struct {
	Address string
	OsmType string
	OsmID   int64
	Long    string
	Lat     string
}

func NewGeo(conf *cfg.Config, args []string) *Geo {
	geocoder := chained.Geocoder(
		nominatim.Geocoder(conf.APIKeys.Nominatim),
		google.Geocoder(conf.APIKeys.GoogleGeocode),
	)
	return &Geo{
		geocoder, args,
	}
}

var (
	errColor    = color.New(color.FgRed).SprintFunc()
	infoColor   = color.New(color.FgBlue).SprintFunc()
	stringColor = color.New(color.FgGreen).SprintFunc()
)

func getTabWriterOutput() *tabwriter.Writer {
	return tabwriter.NewWriter(os.Stdout, 12, 2, 4, ' ', tabwriter.TabIndent)
}

// GetGeocode location
func (g *Geo) GetGeocode() error {
	out := getTabWriterOutput()
	location, err := gosm.GetGeo(g.args[0])
	if err != nil {
		return err
	}
	if location != nil && len(location) > 0 {
		for _, v := range location {
			locref := reflect.ValueOf(gosmData{
				Address: v.Name,
				OsmType: v.OsmType,
				OsmID:   v.OsmID,
				Long:    v.Lon,
				Lat:     v.Lat,
			})
			printRows(out, locref)
		}
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
	if gcode != nil {
		gref := reflect.ValueOf(*gcode)
		printRows(out, gref)
	}
	out.Flush()
	return nil
}

func printRows(out *tabwriter.Writer, rows reflect.Value) {
	for i := 0; i < rows.NumField(); i++ {
		if rows.Field(i).String() != "" {
			printRow(
				out, "\n%s\t%s\t%v",
				rows.Type().Field(i).Name,
				rows.Field(i).Interface(),
			)
		}
	}
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
