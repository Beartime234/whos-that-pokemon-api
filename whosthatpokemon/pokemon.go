package whosthatpokemon

import (
	"github.com/Beartime234/whos-that-pokemon/whosthatpokemon/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"log"
	"math/rand"
	"time"
)

const MaxPokemon = 807

var conf = config.New()

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
func GetRandomPokemon() *Pokemon {
	db := dynamo.New(session.New(), &aws.Config{Region:aws.String("us-east-1", )})
	table := db.Table(conf.GalleryTable.TableName)

	pokedexID := GenerateRandomPokedexID()

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
	randomNumber := rand.Intn(MaxPokemon) + 1
	log.Printf("Pokemon SessionID: %d", randomNumber)
	return randomNumber
}

func NewPokemon() *Pokemon {
	return GetRandomPokemon()
}

func (poke *Pokemon) NewStrippedPokemon() *StrippedPokemon {
	return &StrippedPokemon{BWImageUrl:poke.BWImageUrl}
}