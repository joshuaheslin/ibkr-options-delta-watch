package myfunction

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"

	oauth1 "github.com/dghubble/oauth1"
	openai "github.com/sashabaranov/go-openai"
)

func postTweet(text string) {
	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	postBody, _ := json.Marshal(map[string]string{
		"text": text,
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := httpClient.Post("https://api.twitter.com/2/tweets", "application/json", responseBody)

	fmt.Println(resp)
	fmt.Println(err)
}

func GenerateTweet(name string) string {
	prompt := fmt.Sprintf("Generate a tweet for a software app that %s. Ensure to include trending hashtags.", name)

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       openai.GPT3Dot5Turbo,
			Temperature: 0.8,
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

	postTweet(resp.Choices[0].Message.Content)
	return resp.Choices[0].Message.Content
}

func MakeAudioTourPost() {
	tweet := GenerateTweet("is a website that creates online audio tours using QRcodes")

	fmt.Println(tweet)
}

func MakeFormTrackersPost() {
	tweet := GenerateTweet("makes it easy to see what customers are typing into your website form before they submit.' #webflow, #wordpress #form #website")

	// easystory_me

	fmt.Println(tweet)
}

func MakeEasyStoryPost() {
	tweet := GenerateTweet("generates bedtime stories for children. Put your kids to sleep. Use Emojis. Don't use \"Say goodbye\" words. #bedtime #stories #parents")
	fmt.Println(tweet)
}

func MakePost() {
	fmt.Println("Making posts...")
	// MakeTweet()

	// MakeAudioTourPost()
	MakeEasyStoryPost()
	// MakeFormTrackersPost()

	fmt.Println("Done!")
}
