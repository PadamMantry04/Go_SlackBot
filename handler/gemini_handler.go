// we have to have a structured workflow and look how
package Gemini

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func SetUpGemini(api_key string) *genai.ChatSession {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(api_key))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	cs := model.StartChat()
	return cs
}

func GeminiHandler(cs *genai.ChatSession, input string) (string, error) {
	// if we have any history of chat we can feed this to our model, and then, add successive parts of our chat to this History struct itself,
	// basically use two functions, one to initialize and one repetitive function to add to the new chats
	ctx := context.Background()
	res, err := cs.SendMessage(ctx, genai.Text(input))
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(res);
	var response string
	for _, cand := range res.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				// how to convert this custom defined datatype Part to string
				fmt.Println(part)
				jsonData, err := json.Marshal(part)
				if err != nil {
					// log.Printf("Failed to convert part to JSON: %v", err)
					panic(err)
				}
				response += string(jsonData)
				// most basic way is to convert to a interface and then somehow just pass that interface to print the respons.
			}
		}
	}
	newHistory := []*genai.Content{
		{
			Parts: []genai.Part{
				genai.Text(input),
			},
			Role: "user",
		},
		{
			Parts: []genai.Part{
				genai.Text(response),
			},
			Role: "model",
		},
	}
	cs.History = append(cs.History, newHistory...) // basically adds the chat history for the model.
	return response, err
}
