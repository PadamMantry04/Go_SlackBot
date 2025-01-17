package config

import "github.com/joho/godotenv"

// create a new struct with all required envs and then create and return a new instance of the struct via a wrapper function.

type Env struct {
	// slack app,bot,gemini token
	SLACK_APP_TOKEN string
	SLACK_BOT_TOKEN string
	GEMINI_API_KEY  string
	BOT_ID          string
}

// L is capital since this function shall be used in other files to import the env files.
func LoadEnv() *Env {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()

	if err != nil {
		panic(err)
	} // very basic way of handling an error, may introduce code for graceful shutdown.

	return &Env{
		SLACK_APP_TOKEN: myEnv["SLACK_APP_TOKEN"],
		SLACK_BOT_TOKEN: myEnv["SLACK_BOT_TOKEN"],
		GEMINI_API_KEY:  myEnv["GEMINI_API_KEY"],
		BOT_ID:          myEnv["BOT_ID"],
	}
}
