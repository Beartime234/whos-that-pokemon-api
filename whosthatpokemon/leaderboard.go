package whosthatpokemon

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

var partitionValue = "LEADERBOARD"


func GetTopLeaderboard(limit int64) ([]*MaskedGameSession, error) {
	db := dynamo.New(session.New(), &aws.Config{Region:aws.String("us-east-1", )})
	table := db.Table(conf.SessionTable.TableName)

	var result []*GameSession
	// Does the query for the top ten values in the leaderboard Order False is Descending order.
	// This query's the index for a greater then 0 so that we don't get stupid 0 scores
	err := table.Get(conf.SessionTable.GlobalSecondaryIndex.HashKey, partitionValue).Index(conf.SessionTable.GlobalSecondaryIndex.IndexName).Order(false).Range(conf.SessionTable.GlobalSecondaryIndex.RangeKey, dynamo.Greater, 0).Limit(limit).All(&result)

	if err != nil {
		return nil, err
	}

	// We turn all the sessions into masked sessions
	var maskedResult []*MaskedGameSession
	for _, resultSession := range result {
		maskedSession := resultSession.NewMaskedSession()
		maskedResult = append(maskedResult, maskedSession)
	}

	return maskedResult, nil
}