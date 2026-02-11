package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"google.golang.org/genai"
)

func main() {
	_, _ = askGemini("who are you?")
}

func askGemini(question string) (res string, err error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: os.Getenv("GEMINI_API_KEY"),
	})
	if err != nil {
		return "", err
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash-lite",
		genai.Text(question),
		nil,
	)
	if err != nil {
		return "", err
	}

	return result.Text(), nil
}

type APIResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
			Role string `json:"role"`
		} `json:"content"`
		FinishReason string `json:"finishReason"`
		Index        int    `json:"index"`
	} `json:"candidates"`
}

func askGeminiWithNativeAPI() {
	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.5-flash-lite:generateContent"
	method := "POST"

	payload := strings.NewReader(`{
    "contents": [
        {
            "parts": [
                {
                    "text": "jawab komplain perbankan berikut dengan singkat mengenai = mengapa rekening saya menjadi rekening dormant"
                }
            ]
        }
    ]
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("x-goog-api-key", os.Getenv("GEMINI_API_KEY"))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
