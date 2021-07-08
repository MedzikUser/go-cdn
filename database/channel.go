package database

import (
	"context"
	"time"

	"gitlab.com/gaming0skar123/go/cdn/common"
	"gitlab.com/gaming0skar123/go/cdn/config"
	"go.mongodb.org/mongo-driver/bson"
)

type Channel struct {
	ID      string `bson:"_id"`
	GuildID string `bson:"gid"`
}

type ChannelAPI struct {
	ID string `bson:"_id"`
}

func InsertChannel(channel *Channel) {
	collection := DB.Collection(config.Mongo_Collection_Channel)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, channel)
	common.CheckErr(err, "inserting to DB")
}

func InsertAPIChannel(channel *ChannelAPI) {
	collection := DB.Collection(config.Mongo_Collection_Channel_API)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, channel)
	common.CheckErr(err, "inserting to DB")
}

func FindChannel(channel *Channel) bool {
	collection := DB.Collection(config.Mongo_Collection_Channel)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, channel).Decode(&channel)
	return err == nil
}

func FindAPIChannel() (*ChannelAPI, error) {
	collection := DB.Collection(config.Mongo_Collection_Channel_API)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var channel *ChannelAPI

	err := collection.FindOne(ctx, bson.D{}).Decode(&channel)
	return channel, err
}
