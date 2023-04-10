package myfunction

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func GenerateTweet(name string) string {
	prompt := fmt.Sprintf("Generate a tweet for a software app that %s. Ensure to include trending hashtags.", name)

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       openai.GPT3Dot5Turbo,
			Temperature: 1,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(resp.Choices[0].Message.Content)
	return resp.Choices[0].Message.Content
}

func MakeAudioTourPost() {
	tweet := GenerateTweet("is a website that creates online audio tours using QRcodes")

	fmt.Println(tweet)
}

func MakeEasyStoryPost() {
	tweet := GenerateTweet("makes it easy to see what customers are typing into your website form before they submit.' #webflow, #wordpress #form #website")

	// easystory_me

	fmt.Println(tweet)
}

func MakeFormTrackersPost() {
	tweet := GenerateTweet("generates bedtime stories for children. It helps parents put their kids to sleep. #bedtime #stories #parents")
	fmt.Println(tweet)
}

func MakePost() {
	fmt.Println("Making posts...")

	// MakeAudioTourPost()
	MakeEasyStoryPost()
	// MakeFormTrackersPost()

	fmt.Println("Done!")
}
