package main

import (
	"context"
	"errors"
	"github.com/FrankSantoso/lamat.git/internal/cfg"
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
		Use:   "lamat [-c] <config_file> [r|f] [args]",
		Short: "Small utility to find geocode / reverse geocode",
	}
	findCmd = &cobra.Command{
		Use:   "f [longitude] [latitude]",
		Short: "Try finding address of inputted coordinates / geocode",
		Long:  `Find address of inputted coordinates / geocode along with some extra details`,
		Args:  cobra.ExactArgs(2),
		Run:   findGeocode,
	}
	reverseCmd = &cobra.Command{
		Use:   "r [address]",
		Short: "Try finding geocode information from specified input address",
		Long:  `Find geocode information from inputted address`,
		Args:  cobra.ExactArgs(1),
		Run:   findRevGeocode,
	}
)

func main() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "cfg", "", `
	`, "config files containing nominatim & google api keys")
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
	rootCmd.ExecuteContext(newCtx)
}

func getConfigFromContext(ctx context.Context) (*cfg.Config, error) {
	c := ctx.Value("cfg").(cfg.Config)
	if &c == nil {
		return nil, errors.New("No config in context")
	}
	return &c, nil
}

func findGeocode(c *cobra.Command, args []string) {
	config, err := getConfigFromContext(c.Context())
	if err != nil {
		log.Fatalf("Config file not found in context: %v", err)
	}
}

func findRevGeocode(c *cobra.Command, args []string) {
	config, err := getConfigFromContext(c.Context())
	if err != nil {
		log.Fatalf("Config file not found in context: %v", err)
	}
}
