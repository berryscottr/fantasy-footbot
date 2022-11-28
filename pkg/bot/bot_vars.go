package bot

import (
	"github.com/bwmarrin/discordgo"
)

const (
	// UserID assigns ID to the bot
	UserID = "@me"
	// FantasyFootbotChannelID is the ID of channel #fantasy-footbot
	FantasyFootbotChannelID = "955291440643207228"
)

var (
	Reactions = []string{"🏈"}
)

// Data for the bot to track along a request
type Data struct {
	// Err for error tracking
	Err error
	// User for the bot to track
	User *discordgo.User
	// GoBot for the bot to track
	GoBot *discordgo.Session
	// Token for the bot to track
	Token string
	// Dir for the bot to track
	Dir string
}

// Methods for the bot to use
type Methods interface {
	// SetDir for the bot to use
	SetDir() Data
	// Start the Discord bot listener
	Start()
	// MessageHandler for interpreting which function to launch from message contents
	MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate)
	// ReactionHandler for interpreting how to respond to reactions
	ReactionHandler(s *discordgo.Session, r *discordgo.MessageReactionAdd)
	// HandleReaction for responding to a reaction
	HandleReaction(s *discordgo.Session, r *discordgo.MessageReactionAdd)
	// HandleFoot for posting foot
	HandleFoot(s *discordgo.Session, m *discordgo.MessageCreate)
}