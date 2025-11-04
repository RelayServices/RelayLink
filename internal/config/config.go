package config

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env-default:"prod" env-oneof:"local dev prod"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

// MustLoad reads and parses the config file.
// Terminates the app if CONFIG_PATH is not set
func MustLoad() *Config {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	if *configPath == "" {
		*configPath = os.Getenv("CONFIG_PATH")
	}

	if *configPath == "" {
		*configPath = "./config/local.yaml"
		log.Printf("CONFIG_PATH is not set, using default path: %s", *configPath)
	}

	if _, err := os.Stat(*configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exits: %s", *configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(*configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}