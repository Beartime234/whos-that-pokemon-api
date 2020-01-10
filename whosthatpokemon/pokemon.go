package whosthatpokemon

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"log"
	"math/rand"
	"time"
)


// Pokemon Details about a Pokemon
type Pokemon struct {
	PokedexID int `dynamo:"PokedexID"`
	Name string `dynamo:"Name"`
	OriginalImageUrl string `dynamo:"OriginalImageUrl"`
	BWImageUrl string `dynamo:"BWImageUrl"`
}

// This type is what we return to users so that they can't see everything
type StrippedPokemon struct {
	BWImageUrl string
}

//GetRandomPokemon Gets a random pokemon from the gallery database
func GetRandomPokemon(notThesePokemon []int) *Pokemon {
	db := dynamo.New(session.New(), &aws.Config{Region:aws.String("us-east-1", )})
	table := db.Table(conf.GalleryTable.TableName)

	pokedexID := GenerateRandomPokedexID()

	// This prevents us randomly picking the pokemon in notThesePokemon
	for intInSlice(pokedexID, notThesePokemon) {
		pokedexID = GenerateRandomPokedexID()
	}

	var result *Pokemon
	err := table.Get(conf.GalleryTable.HashKey, pokedexID).One(&result)

	if err != nil {
		panic(err) // No point
	}

	log.Printf("Pokemon Data: %+v\n", result)

	return result
}

//GenerateRandomPokedexID Generates a random pokedex id
func GenerateRandomPokedexID() int {
	rand.Seed(time.Now().UnixNano()) // Generate a seed so it's random every time we call this
	randomNumber := rand.Intn(conf.MaxPokemon) + 1
	log.Printf("Pokemon PokedexID: %d", randomNumber)
	return randomNumber
}

// newPokemon Gets you a new pokemon picked randomly
func newPokemon(notThesePokemon []int) *Pokemon {
	return GetRandomPokemon(notThesePokemon)
}

// Pokemon_NewStrippedPokemon The information of a pokemon that can be shared to the user
func (poke *Pokemon) NewStrippedPokemon() *StrippedPokemon {
	return &StrippedPokemon{BWImageUrl:poke.BWImageUrl}
}

// intInSlice Helper function for checking if an int is in a slice
func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}