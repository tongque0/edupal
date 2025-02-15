package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

func main() {
	config := openai.DefaultConfig("sk-MlUCO1oxmEHjx0D6UDeqfQzK7kVd1KAccoc0N41pbvJw0lx4")
	config.BaseURL = "https://api.zetatechs.com/v1"
	client := openai.NewClientWithConfig(config)
	ctx := context.Background()

	type Result struct {
		Questions []struct {
			Question string `json:"question"`
			Answer   string `json:"answer"`
		} `json:"questions"`
	}
	var result Result
	schema, err := jsonschema.GenerateSchemaForType(result)
	if err != nil {
		log.Fatalf("GenerateSchemaForType error: %v", err)
	}
	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "你是一名专业的初高中老师，可以按照要求输出一定的初高中科目题目，生成内容包含题目，答案",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "初二物理牛顿第二定律类的选择题 3道",
			},
		},
		ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeJSONSchema,
			JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
				Name:        "create_question",
				Description: "将生成的题目进行格式化整理,将答案，题目，提示分开",
				Schema:      schema,
				Strict:      true,
			},
		},
	})
	if err != nil {
		log.Fatalf("CreateChatCompletion error: %v", err)
	}
	err = schema.Unmarshal(resp.Choices[0].Message.Content, &result)
	if err != nil {
		log.Fatalf("Unmarshal schema error: %v", err)
	}
	fmt.Println("question:", result.Questions)
	// fmt.Println("answer:", result.Answer)
}
