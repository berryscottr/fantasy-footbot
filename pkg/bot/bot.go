package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"path/filepath"
)

// SetDir for setting the directory for the bot
func (bot Data) SetDir() Data {
	bot.Dir, bot.Err = filepath.Abs(".")
	if bot.Err != nil {
		log.Err(bot.Err).Msg("failed to set fantasy-footbot directory")
	}
	return bot
}

// Start the Discord bot listener
func (bot Data) Start() {
	bot.GoBot, bot.Err = discordgo.New("Bot " + bot.Token)
	if bot.Err != nil {
		log.Err(bot.Err).Msg("failed to instantiate fantasy-footbot bot")
		return
	}
	bot.User, bot.Err = bot.GoBot.User(UserID)
	if bot.Err != nil {
		log.Err(bot.Err).Msg("failed to set fantasy-footbot user id")
		return
	}
	bot.GoBot.AddHandler(bot.MessageHandler)
	bot.GoBot.AddHandler(bot.ReactionHandler)
	bot.Err = bot.GoBot.Open()
	if bot.Err != nil {
		log.Err(bot.Err).Msg("failed to start fantasy-footbot listener")
		return
	}
	log.Info().Msg("fantasy-footbot listening")
	return
}
