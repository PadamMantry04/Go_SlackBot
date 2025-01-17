package slack

import (
	"context"

	Gemini "github.com/PadamMantry04/Go_SlackBot/handler"
	"github.com/google/generative-ai-go/genai"
	"github.com/rs/zerolog"
	"github.com/shomali11/slacker"
	// "github.com/shomali11/slacker"
)

// Bot encapsulates the Slack bot and its dependencies
type Bot struct {
	API   *slacker.Slacker
	Log   *zerolog.Logger
	BotID string
}

// NewBot initializes and returns a new Bot instance
func NewBot(bot_id string, bot_token string, app_token string) *Bot {
	api := slacker.NewClient(bot_token, app_token)
	log := new(zerolog.Logger)
	return &Bot{
		API:   api,
		Log:   log,
		BotID: bot_id,
	}
}

// figure out where you can start a new gemini chat session. -> Done in main.go
// RegisterHandlers sets up the command handlers for the bot -> DoNE

func RegisterGeminiCommand(b *Bot, chatSession *genai.ChatSession) {
	command := &slacker.CommandDefinition{
		Description: "Handles any generic query and responds using the Gemini API.",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			// Retrieve the user input
			b.Log.Info().Msg("Command received!")
			text := request.StringParam("message", "")

			if botCtx.Event().UserID == b.BotID {
				// Bot sent the message, so return immediately without responding
				return
			}

			// Get Gemini's response
			geminiResponse, err := Gemini.GeminiHandler(chatSession, text)
			if err != nil {
				b.Log.Error().Err(err).Msg("Failed to get Gemini response")
				response.Reply("Sorry, I encountered an error processing your request.")
				return
			}

			// Send the response back to Slack
			if err := response.Reply(geminiResponse); err != nil {
				b.Log.Error().Err(err).Msg("Failed to send message to Slack")
				return
			}

			b.Log.Info().
				Str("query", text).
				Str("response", geminiResponse).
				Msg("Successfully processed Gemini query")
		},
	}

	// Register the command with a catch-all pattern
	b.API.Command(".*", command)
}

// Start begins listening for Slack events
func (b *Bot) Start() error {
	b.Log.Info().Msg("Starting Slack bot...")
	return b.API.Listen(context.Background())
}
