// Code generated by wit-bindgen-go. DO NOT EDIT.

package chat

import (
	"github.com/devigned/wasm-packs/compose/internal/example/domain/types"
	"go.bytecodealliance.org/cm"
	"unsafe"
)

// ChatResponseShape is used for storage in variant or result types.
type ChatResponseShape struct {
	_     cm.HostLayout
	shape [unsafe.Sizeof(ChatResponse{})]byte
}

func lift_ChatRequest(f0 *uint8, f1 uint32, f2 *types.Message, f3 uint32) (v types.ChatRequest) {
	v.Model = cm.LiftString[string](f0, f1)
	v.Messages = cm.LiftList[cm.List[types.Message]](f2, f3)
	return
}
