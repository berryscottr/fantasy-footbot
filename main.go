package main

import (
	"fantasy-footbot/pkg/bot"
	"os"
)

func main() {
	gobot := bot.Data{Token: os.Getenv("BOT_TOKEN")}
	if gobot.Token == "" {
		panic("BOT_TOKEN not set")
	}
	gobot = gobot.SetDir()
	gobot.Start()
	<-make(chan struct{})
	return
}
