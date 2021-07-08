package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var Bot_Token = os.Getenv("BOT_TOKEN")
var Webhook_URL = os.Getenv("WEBHOOK_URL")

var Embed_Color = 1127128
