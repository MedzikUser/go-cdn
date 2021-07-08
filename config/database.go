package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var Mongo_URI = os.Getenv("MONGODB")

var Mongo_Database = "cdn"

var Mongo_Collection_Imgur = "imgur"
var Mongo_Collection_Channel = "channel"
var Mongo_Collection_Channel_API = "api_channel"
