//go:generate wkg wit fetch
//go:generate wit-bindgen-go generate -o ./internal ./wit

package main

import (
	"context"
	"fmt"
	"github.com/devigned/wasm-packs/compose/internal/example/domain/adder"
	"github.com/devigned/wasm-packs/compose/internal/wasi/cli/run"
	"github.com/openai/openai-go"
	_ "github.com/ydnar/wasi-http-go/wasihttp" // enable wasi-http
	"go.bytecodealliance.org/cm"
)

func init() {
	adder.Exports.Add = func(x int32, y int32) int32 {
		// This is where you would implement the logic for the add function.
		// For example, you could return the sum of x and y.
		return x + y
	}

	run.Exports.Run = func() cm.BoolResult {
		fmt.Println("Running the program...")
		fmt.Println(generateImage("gpt-4", "a cat"))
		fmt.Println("Program finished.")
		return true
	}
}

// main is required for the `wasi` target, even if it isn't used.
func main() {
}

func generateImage(model, prompt string) (string, error) {
	client := openai.NewClient(
	// defaults to os.LookupEnv("OPENAI_API_KEY")
	)
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("Say this is a test"),
		},
		Model: openai.ChatModelGPT4o,
	})
	if err != nil {
		return "", err
	}
	return chatCompletion.Choices[0].Message.Content, nil
}
