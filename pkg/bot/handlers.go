package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"golang.org/x/exp/slices"
	"strings"
)

// MessageHandler for interpreting which function to launch from message contents
func (bot Data) MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == bot.User.ID {
		return
	}
	if strings.Contains(strings.ToLower(m.Content), "!foot") {
		bot.HandleFoot(s, m)
	}
	return
}

// ReactionHandler for interpreting how to respond to reactions
func (bot Data) ReactionHandler(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if r.Member.User.ID == bot.User.ID {
		return
	}
	if slices.Contains(Reactions, r.MessageReaction.Emoji.Name) &&
		r.MessageReaction.ChannelID == FantasyFootbotChannelID {
		bot.HandleReaction(s, r)
	}
	return
}

// HandleReaction for handling the football reaction
func (bot Data) HandleReaction(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	log.Info().Msg("handling football reaction")
	var reaction string
	switch r.MessageReaction.Emoji.Name {
	case "🏈":
		reaction = "football"
	default:
		log.Info().Msg("unknown reaction")
		return
	}
	name := r.Member.Nick
	if name == "" {
		name = r.Member.User.Username
	}
	message := discordgo.MessageSend{
		Content: fmt.Sprintf(
			"%s found the %s.", name, reaction,
		),
	}
	if r.MessageReaction.ChannelID == FantasyFootbotChannelID {
		_, bot.Err = s.ChannelMessageSendComplex(FantasyFootbotChannelID, &message)
	}
	if bot.Err != nil {
		log.Err(bot.Err).Msg("failed to post message")
		return
	}
	log.Info().Msgf("%s reaction from %s to football posted to Discord channel %s",
		r.MessageReaction.Emoji.Name, name, r.ChannelID)
	return
}

// HandleFoot for posting foot
func (bot Data) HandleFoot(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Info().Msg("handling football message")
	name := m.Member.Nick
	if name == "" {
		name = m.Member.User.Username
	}
	message := discordgo.MessageSend{
		Content: fmt.Sprintf(
			"Hi, %s. I'm a work in progress fantasy football bot.\n" +
				name,
		),
	}
	if m.ChannelID == FantasyFootbotChannelID {
		_, bot.Err = s.ChannelMessageSendComplex(FantasyFootbotChannelID, &message)
	}
	if bot.Err != nil {
		log.Err(bot.Err).Msg("failed to post message")
		return
	}
	log.Info().Msgf("football message posted to Discord channel %s", m.ChannelID)
	return
}
