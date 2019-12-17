package config

import (
	"log"
	"os"
)

type Config struct {
	GalleryTable DynamoTableConfig
	SessionTable DynamoTableConfig
}

type DynamoTableConfig struct {
	TableName string
	HashKey string
}

// New returns a new Config struct
func New() *Config {
	return &Config{
		GalleryTable: DynamoTableConfig{
			TableName: getEnv("GALLERY_TABLE_NAME"),
			HashKey: "PokedexID",
		},
		SessionTable:DynamoTableConfig{
			TableName: getEnv("SESSION_TABLE_NAME"),
			HashKey:   "SessionID",
		},
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	log.Panicf("Missing Environment Variable %s", key)
	return ""
}
