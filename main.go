package main

import (
	"encoding/json"
	"fmt"
	gpt "github.com/DanPlayer/chatgpt-sdk/v1"
	"io/ioutil"
	"os"
)

const SecretKey = "your gpt api key"

// ChatGpt init gpt client, if you are in China, you should use proxy
var ChatGpt = gpt.Client(gpt.ChatGptOption{SecretKey: SecretKey, HasProxy: true, ProxyUrl: "http://localhost:7890"})

func main() {
	// open json file
	jsonFile, err := os.Open("gpt/ProgrammerLiMing.json")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()

	GptCreate(jsonFile)
}

// GptCreate 生成Gpt对话
func GptCreate(jsonFile *os.File) {
	msg := make([]gpt.ChatMessage, 0)
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err := json.Unmarshal(byteValue, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	msg = append(msg, gpt.ChatMessage{
		Role:    "user",
		Content: "你的性格怎么样",
	})

	response, err := ChatGpt.CreateChatCompletion(gpt.CreateChatCompletionRequest{
		Model:    gpt.GPT3Dot5Turbo,
		Messages: msg,
	})
	if err != nil {
		fmt.Printf("completions error: %s", err.Error())
		return
	}
	fmt.Println(response)
}
