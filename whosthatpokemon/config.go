package whosthatpokemon

import (
	"log"
	"os"
)

var conf = New()

type Config struct {
	MaxPokemon          int
	CorrectnessRequired int
	GalleryTable        DynamoTableConfig
	SessionTable        DynamoTableConfig
}

type DynamoTableConfig struct {
	TableName string
	HashKey string
	GlobalSecondaryIndex GlobalSecondaryIndexConfig
}

type GlobalSecondaryIndexConfig struct {
	IndexName string
	HashKey string
	RangeKey string
} 

// New returns a new Config struct
func New() *Config {
	return &Config{
		MaxPokemon:          807,
		CorrectnessRequired: 2,
		GalleryTable: DynamoTableConfig{
			TableName: getEnv("GALLERY_TABLE_NAME"),
			HashKey:   "PokedexID",
		},
		SessionTable: DynamoTableConfig{
			TableName: getEnv("SESSION_TABLE_NAME"),
			HashKey:   "SessionID",
			GlobalSecondaryIndex:GlobalSecondaryIndexConfig{
				IndexName: "Leaderboard-Index",
				HashKey:  "LeaderboardPartition",
				RangeKey: "Score",
			},
		},
	}
}

// Simple helper function to read an environment or panic with a good error message
func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	log.Panicf("Missing Environment Variable %s", key)
	return ""
}
