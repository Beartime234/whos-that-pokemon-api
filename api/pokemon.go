package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"log"
	"math/rand"
	"time"
)

const MaxPokemon = 807

type Pokemon struct {
	PokedexID int `dynamo:"PokedexID"`
	Name string `dynamo:"Name"`
	OriginalImageUrl string `dynamo:"OriginalImageUrl"`
	BWImageUrl string `dynamo:"BWImageUrl"`
}

//GetRandomPokemon Gets a random pokemon from the gallery database
func GetRandomPokemon() *Pokemon {
	db := dynamo.New(session.New(), &aws.Config{Region:aws.String("us-east-1", )})
	table := db.Table(GalleryTableName)

	pokedexID := GenerateRandomPokedexID()

	var result *Pokemon
	err := table.Get(GalleryTableHashKey, pokedexID).One(&result)

	if err != nil {
		panic(err) // No point
	}

	log.Printf("Pokemon Data: %+v\n", result)

	return result
}

//GenerateRandomPokedexID Generates a random pokedex id
func GenerateRandomPokedexID() int {
	rand.Seed(time.Now().UnixNano()) // Generate a seed so it's random every time we call this
	randomNumber := rand.Intn(MaxPokemon) + 1
	log.Printf("Pokemon ID: %d", randomNumber)
	return randomNumber
}

func NewPokemon() *Pokemon {
	return GetRandomPokemon()
}