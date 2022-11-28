package bot

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// TestData_SetDir confirms ability to set the bot directory
func TestData_SetDir(t *testing.T) {
	assertion := assert.New(t)
	data := Data{Token: os.Getenv("BOT_TOKEN")}
	data = data.SetDir()
	assertion.NoError(data.Err, "failed to set the bot directory")
}

// TestData_Start confirms ability to start the discord listener
func TestData_Start(t *testing.T) {
	assertion := assert.New(t)
	data := Data{Token: os.Getenv("BOT_TOKEN")}
	data = data.SetDir()
	data.Start()
	assertion.NoError(data.Err, "failed to start discord listener")
}

// TestData_MessageHandler confirms ability to handle message contents
func TestData_MessageHandler(t *testing.T) {
	assertion := assert.New(t)
	data := Data{Token: os.Getenv("BOT_TOKEN")}
	data = data.SetDir()
	assertion.NoError(data.Err, "failed to handle message contents")
}

// TestData_ReactionHandler confirms ability to handle reactions
func TestData_ReactionHandler(t *testing.T) {
	assertion := assert.New(t)
	data := Data{Token: os.Getenv("BOT_TOKEN")}
	data = data.SetDir()
	assertion.NoError(data.Err, "failed to handle reaction")
}

// TestData_HandleReaction confirms ability to handle the reaction
func TestData_HandleReaction(t *testing.T) {
	assertion := assert.New(t)
	data := Data{Token: os.Getenv("BOT_TOKEN")}
	data = data.SetDir()
	assertion.NoError(data.Err, "failed to read a game day reaction")
}

// TestData_HandleFoot confirms ability to handle the foot
func TestData_HandleFoot(t *testing.T) {
	assertion := assert.New(t)
	data := Data{Token: os.Getenv("BOT_TOKEN")}
	data = data.SetDir()
	assertion.NoError(data.Err, "failed to read a game day reaction")
}
