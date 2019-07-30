package client
import (
"github.com/andyliuliming/gogomodule/mockexternals/tinyexternal"
)

func ClientFunc() {
	NewTinyExternal()
}

func NewTinyExternal() *tinyexternal.TinyExternal{
	return &tinyexternal.TinyExternal{
	}
}