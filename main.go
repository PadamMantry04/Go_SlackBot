// Your Go server will:

// Handle incoming Slack events. -> Slacker based implemnentation
// Forward user messages to the Gemini API. ->
// Send Gemini's responses back to Slack.

// Get env variables from a env file
// A new route for slack absed events

// a new function to send request to gemini and then get the response and then

// a new function to display back to slack

package main

import (
	"github.com/PadamMantry04/Go_SlackBot/config"
	Gemini "github.com/PadamMantry04/Go_SlackBot/handler"
	"github.com/PadamMantry04/Go_SlackBot/logger"
	"github.com/PadamMantry04/Go_SlackBot/slack"
)

func main() {
	// Initialize logger
	log := logger.InitLogger()
	// Load environment variables
	env := config.LoadEnv()
	cs := Gemini.SetUpGemini(env.GEMINI_API_KEY) // returns a chat session
	// Start the Slack bot
	bot := slack.NewBot(env.BOT_ID, env.SLACK_BOT_TOKEN, env.SLACK_APP_TOKEN)
	bot.Log = &log
	slack.RegisterGeminiCommand(bot, cs)

	// Start listening for Slack events
	err := bot.Start()
	if err != nil {
		bot.Log.Fatal().Err(err).Msg("Failed to start Slack bot")
		bot.Log.Fatal().AnErr("Error:", err)
	}
}
