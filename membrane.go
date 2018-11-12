package gohdk

import (
	"bytes"
	"encoding/json"
	"syscall/js"
	"unsafe"
)

type HcGlobals struct {
	DnaName,
	DnaHash,
	AgentID,
	AgentAddress,
	AgentInitialHash,
	AgentLatestHash string
}

func InitGlobals(globals HcGlobals) error {

}

func PackJSON(args ...string) ([]int32, error) {
	for _, arg := range args {
		init := js.NewCallback()
	}
}
