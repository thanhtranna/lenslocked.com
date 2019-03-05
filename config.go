package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// PostgresConfig has define config environment
type PostgresConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
// Dialect is define name database
func (c PostgresConfig) Dialect() string {
	return "postgres"
}

// ConnectionInfo get connection info
func (c PostgresConfig) ConnectionInfo() string {
	if c.Password == "" {
		return fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Name)
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Name)
}

// DefaultPostgresConfig set default postgres config
func DefaultPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		Name:     "lenslocked_dev",
	}
}

// Config type config
type Config struct {
	Port     int            `json:"port"`
	Env      string         `json:"env"`
	Pepper   string         `json:"pepper"`
	HMACKey  string         `json:"hmackey"`
	Database PostgresConfig `json:"database"`
}

// IsProd check production
func (c Config) IsProd() bool {
	return c.Env == "prod"
}

// DefaultConfig return config values
func DefaultConfig() Config {
	return Config{
		Port:     3000,
		Env:      "dev",
		Pepper:   "lDS3aue165e3",
		HMACKey:  "secret-hmac-key",
		Database: DefaultPostgresConfig(),
	}
}

// LoadConfig load config into .config file
func LoadConfig(configReq bool) Config {
	f, err := os.Open(".config")
	if err != nil {
		if configReq {
			panic(err)
		}
		fmt.Println("Using the default config...")
		return DefaultConfig()
	}
	var c Config
	dec := json.NewDecoder(f)
	err = dec.Decode(&c)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully loaded .config")
	return c
}