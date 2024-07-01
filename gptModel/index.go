package gptmodel

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"quizGo/constants"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var client *genai.Client

func init() {
	GENAISECRET := constants.EnvConstant("GENAISECRET")

	ctx := context.Background()

	var err error

	client, err = genai.NewClient(ctx, option.WithAPIKey(GENAISECRET))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("GenAI Client is created")
}

func GPTResults(prompt string, results interface{}) error {
	ctx := context.Background()

	model := client.GenerativeModel("gemini-pro")

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))

	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.Marshal(resp.Candidates[0].Content)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal([]byte(jsonData), results)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))

	return nil
}
