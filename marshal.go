package gohdk

// +build js,wasm

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/gonuts/binary"
)

var (
	errArgMarshalErr = errors.New("unable to convert argument to int32")
	errEncodingErr   = errors.New("unable to encode into binary representation")
)

type ArgMarshaler interface {
	MarshalArg() WasmArg
}

type WasmArg int32

type Arg interface {
	GoType() string
	ValString() string
}

type GoInt int

func (i GoInt) MarshalArg() WasmArg {
	return WasmArg(i)
}

func (i GoInt) GoType() string {
	return fmt.Sprintf("%T", i)
}

func (i GoInt) ValString() string {
	return string(i)
}

func marshalWASMargs(args ...Arg) []byte {
	var argbuff []byte
	var writer = bytes.NewBuffer(argbuff)
	var encoder = binary.NewEncoder(writer)
	for _, v := range args {
		errArgMarshalErr = encoder.Encode(v)
	}
	return writer.Bytes()
}
