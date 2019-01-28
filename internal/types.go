package gohdk

import (
	"syscall/js"
	"unsafe"

	// "github.com/gowasm/dom"
	"github.com/dennwc/dom"
)

var (
	_ dom.Event
	_ js.Value
	_ unsafe.Pointer
)

// Wstack is abstract representation of the WASM stack (growable; 64kB pagesize)
type Wstack struct {
	stack []uint64
}

func (s *wasmStack) Push(bt uint64) {
	s.stack = append(s.stack, bt)
}

func (s *wasmStack) Pop() uint64 {
	top := s.Top()

}
