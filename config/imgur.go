package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var Proxy_Domain = os.Getenv("PROXY_DOMAIN")
var Imgur_Client_ID = os.Getenv("IMGUR_CLIENT_ID")
var Imgur_Album_ID = ""
