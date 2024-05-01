package main

import (
	"flag"
	"log"
	"sbercheck/internal/config"
	sberparser "sbercheck/internal/parser"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/config.yaml", "Path to config file")
}

func main() {
	flag.Parse()
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	if err := sberparser.Start(cfg); err != nil {
		log.Fatal(err)
	}
}
