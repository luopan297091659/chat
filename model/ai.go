package model

import (
	"context"
	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
	"net/http"
	"log"
)

var token string = ""

func ChatGpt(c *gin.Context, prompt string) string {
	client := openai.NewClient(token)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		log.Print("ChatCompletion error: %v\n", err)
		return "request filed"
	}

	log.Print(resp.Choices[0].Message.Content)
	return resp.Choices[0].Message.Content
}

func AjaxTest(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chat.html", nil)
}

func PostAjax(ctx *gin.Context) {
	question := ctx.PostForm("question")
	//age := ctx.PostForm("age")
	log.Print(question)
	//fmt.Println(age)
	var ss string = ChatGpt(ctx, question)
	messgae_map := map[string]interface{}{
		"code":200,
		"msg":ss,
	}
	ctx.JSON(http.StatusOK,messgae_map)
}
