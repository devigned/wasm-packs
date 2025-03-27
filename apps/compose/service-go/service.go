//go:generate wkg wit fetch
//go:generate wit-bindgen-go generate -o ./internal ./wit

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/devigned/wasm-packs/compose/internal/example/domain/adder"
	"github.com/devigned/wasm-packs/compose/internal/example/domain/chat"
	domainTypes "github.com/devigned/wasm-packs/compose/internal/example/domain/types"
	"github.com/devigned/wasm-packs/compose/internal/wasi/cli/run"
	wasihttp "github.com/ydnar/wasi-http-go/wasihttp"
	"go.bytecodealliance.org/cm"
)

func init() {
	adder.Exports.Add = func(x int32, y int32) int32 {
		// This is where you would implement the logic for the add function.
		// For example, you could return the sum of x and y.
		return x + y
	}

	run.Exports.Run = func() cm.BoolResult {
		res, err := chatCompletion(chat.ChatRequest{
			Model: "gpt-4o",
			Messages: cm.ToList([]domainTypes.Message{
				{
					Role:    "user",
					Content: "Hello!",
				},
			}),
		})

		if chatErr, ok := err.(chatError); ok {
			fmt.Println("Chat error: ", chatErr)
			return false
		}

		fmt.Println("Response: ", res.Choices.Slice()[0].Message.Content)
		return true
	}

	chat.Exports.Chat = func(request chat.ChatRequest) cm.Result[chat.ChatResponseShape, chat.ChatResponse, chat.Error] {
		res, err := chatCompletion(request)
		if chatErr, ok := err.(chatError); ok {
			return cm.Err[cm.Result[chat.ChatResponseShape, chat.ChatResponse, chat.Error]](chat.Error{
				Code:    chatErr.Code,
				Message: chatErr.Message,
			})
		}
		return cm.OK[cm.Result[chat.ChatResponseShape, chat.ChatResponse, chat.Error]](res)
	}

}

// main is required for the `wasi` target, even if it isn't used.
func main() {
}

type chatError struct {
	Code    int32
	Message string
}

func (e chatError) Error() string {
	return fmt.Sprintf("Chat error %d: %s", e.Code, e.Message)
}

func chatCompletion(chatRequest chat.ChatRequest) (chat.ChatResponse, error) {
	client := &http.Client{
		Transport: &wasihttp.Transport{},
	}

	reqBody, err := json.Marshal(chatRequest)
	if err != nil {
		return chat.ChatResponse{}, chatError{Code: 0, Message: "failed to marshal request"}
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewReader(reqBody))
	if err != nil {
		return chat.ChatResponse{}, chatError{Code: 0, Message: "failed to create request"}
	}
	if req == nil {
		return chat.ChatResponse{}, chatError{Code: 0, Message: "request is nil"}
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))

	res, err := client.Do(req)
	if err != nil {
		return chat.ChatResponse{}, chatError{Code: 0, Message: "failed to send request"}
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return chat.ChatResponse{}, chatError{Code: int32(res.StatusCode), Message: "failed to get valid response"}
	}

	chatRes := chat.ChatResponse{}
	bits, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return chat.ChatResponse{}, chatError{Code: int32(res.StatusCode), Message: "failed to read response body"}
	}
	if err := json.Unmarshal(bits, &chatRes); err != nil {
		return chat.ChatResponse{}, chatError{Code: int32(res.StatusCode), Message: "failed to unmarshal response"}
	}
	return chatRes, nil
}

/*
Request:
curl https://api.openai.com/v1/chat/completions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $OPENAI_API_KEY" \
  -d '{
    "model": "gpt-4o",
    "messages": [
      {
        "role": "developer",
        "content": "You are a helpful assistant."
      },
      {
        "role": "user",
        "content": "Hello!"
      }
    ]
  }

*/

/*
Response:
{
  "id": "chatcmpl-B9MBs8CjcvOU2jLn4n570S5qMJKcT",
  "object": "chat.completion",
  "created": 1741569952,
  "model": "gpt-4o-2024-08-06",
  "choices": [
    {
      "index": 0,
      "message": {
        "role": "assistant",
        "content": "Hello! How can I assist you today?",
        "refusal": null,
        "annotations": []
      },
      "logprobs": null,
      "finish_reason": "stop"
    }
  ],
  "usage": {
    "prompt_tokens": 19,
    "completion_tokens": 10,
    "total_tokens": 29,
    "prompt_tokens_details": {
      "cached_tokens": 0,
      "audio_tokens": 0
    },
    "completion_tokens_details": {
      "reasoning_tokens": 0,
      "audio_tokens": 0,
      "accepted_prediction_tokens": 0,
      "rejected_prediction_tokens": 0
    }
  },
  "service_tier": "default"
}
*/
