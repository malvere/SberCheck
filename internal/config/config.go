package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env      string
	Database DatabaseConfig `yaml:"database" env-required:"true"`
	Parser   ParserConfig   `yaml:"parser" env-required:"true"`
}

type DatabaseConfig struct {
	Driver string
	URL    string
}

type ParserConfig struct {
	Experiments  string
	CookieFolder string `yaml:"cookie-folder"`
}

func LoadConfig(cfgPath string) (*Config, error) {
	path := fetchConfigPath(cfgPath)
	if path == "" {
		panic("Config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("Config file does not exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("Failed to read config: " + err.Error())
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL != "" {
		cfg.Database.URL = dbURL
	}
	return &cfg, nil

}

func fetchConfigPath(cfgPath string) string {
	var res string

	// --config="path/to/cfg.yaml"
	flag.StringVar(&res, "config", "", "path")
	flag.Parse()

	if cfgPath == "" {
		res = os.Getenv("CONFIG_PATH")
	} else {
		res = cfgPath
	}
	return res
}
