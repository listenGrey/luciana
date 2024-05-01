package logic

import (
	"context"
	"github.com/tmc/langchaingo/llms/ollama"
)

func Generate(prompt string) (res string, err error) {
	llm, err := ollama.New(ollama.WithModel("llama3"))
	if err != nil {
		return
	}
	res, err = llm.Call(context.Background(), prompt)
	if err != nil {
		return
	}

	return
}
