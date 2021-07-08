package database

import (
	"context"
	"time"

	"gitlab.com/gaming0skar123/go/cdn/common"
	"gitlab.com/gaming0skar123/go/cdn/config"
)

type Imgur struct {
	ID      string `bson:"_id"`
	DelHash string `bson:"delHash"`
}

func InsertImgur(imgur *Imgur) {
	collection := DB.Collection(config.Mongo_Collection_Imgur)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, imgur)
	common.CheckErr(err, "inserting to DB")
}
