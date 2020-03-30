package main

import (
	"flag"
	"log"
)

var (
	cfgFile = flag.String("c", "config",
		"config file containing spec in cfg package")
)

func main() {
	flag.Parse()
	if cfgFile == nil {
		log.Fatalf("configuration file must be set")
	}

}
