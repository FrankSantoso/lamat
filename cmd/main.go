package main

import (
	"context"
	"errors"
	"github.com/FrankSantoso/lamat.git/internal/cfg"
	"github.com/FrankSantoso/lamat.git/internal/repo"
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"
)

const (
	globalTimeout = 10 * time.Second
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "lamat [-c] <config_file> [rev|find] [args]",
		Short: "Small utility to find geocode / reverse geocode",
	}
	findCmd = &cobra.Command{
		Use:   "find [address]",
		Short: "Try finding address of inputted coordinates / geocode",
		Long:  `Find address of inputted coordinates / geocode along with some extra details`,
		Args:  cobra.ExactArgs(1),
		Run:   findGeocode,
	}
	reverseCmd = &cobra.Command{
		Use:   "rev -- [longitude] [latitude]",
		Short: "Try finding geocode information from specified input address",
		Long:  `Find geocode information from inputted address`,
		Args:  cobra.ExactArgs(2),
		Run:   findRevGeocode,
	}
)

func main() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "cfg", "c",
		"./config", "config files containing nominatim & google api keys")
	rootCmd.AddCommand(findCmd, reverseCmd)
	// create new context
	ctx, _ := context.WithTimeout(context.Background(), globalTimeout)
	config, err := cfg.ReadConfig(cfgFile)
	if err != nil {
		log.Fatalf("Config file error: \n%v", err)
		os.Exit(1)
	}
	// embeds config to the context
	newCtx := context.WithValue(ctx, "cfg", config)
	if err = rootCmd.ExecuteContext(newCtx); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func getConfigFromContext(ctx context.Context) (*cfg.Config, error) {
	c := ctx.Value("cfg").(*cfg.Config)
	if &c == nil {
		return nil, errors.New("No config in context")
	}
	return c, nil
}

func findGeocode(c *cobra.Command, args []string) {
	config, err := getConfigFromContext(c.Context())
	if err != nil {
		log.Fatalf("Config file not found in context: %v", err)
	}
	g := repo.NewGeo(config, args)
	if err = g.GetGeocode(); err != nil {
		log.Fatalf("Error finding geocode: %v", err)
		os.Exit(1)
	}
}

func findRevGeocode(c *cobra.Command, args []string) {
	config, err := getConfigFromContext(c.Context())
	if err != nil {
		log.Fatalf("Config file not found in context: %v", err)
	}
	g := repo.NewGeo(config, args)
	if err = g.GetReverseGeocode(); err != nil {
		log.Fatalf("Error finding geocode: %v", err)
		os.Exit(1)
	}
}
