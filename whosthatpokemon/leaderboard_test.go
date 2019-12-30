package whosthatpokemon

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"log"
	"testing"
)

func Test_getTopLeaderboard(t *testing.T) {
	gameSession, err := NewGameSession()

	if err != nil {
		log.Printf("Wasnt a mistake with the sessison was a mistake creating the session.")
		t.Fail()
		return
	}

	err = gameSession.save()

	if err != nil {
		log.Printf("Wasnt a mistake with the session was a error saving the session.")
		t.Fail()
		return
	}

	const testLimit = 1
	values, err := GetTopLeaderboard(testLimit)

		if err != nil {
		log.Fatal(err)
	}

	if len(values) > testLimit {
		log.Printf("Too many values for the limit %d", len(values))
		t.Fail()
		return
	}

	gameSession.Score = 99999999
	err = gameSession.save()

	if err != nil {
		log.Printf("Wasnt a mistake with the session was a mistake saving the session.")
		t.Fail()
		return
	}

	values, err = GetTopLeaderboard(10)

	if values[0].Score != 99999999 {
		log.Print("The top score was not the top score")
		t.Fail()
	}

	db := dynamo.New(session.New(), &aws.Config{Region:aws.String("us-east-1", )})
	table := db.Table(conf.SessionTable.TableName)
	err = table.Delete(conf.SessionTable.HashKey, gameSession.SessionID).Run()  // Delete it so that we don't have to worry about it failing next time

	if err != nil {
		log.Printf("Error deleting the session.")
		t.Fail()
		return
	}
}