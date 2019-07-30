package client

import (
	"github.com/andyliuliming/gogomodule/mockexternals/tinyexternal"
)

func NewTinyExternal() *tinyexternal.TinyExternal{
	return &tinyexternal.TinyExternal{
	}
}

func ClientFunc() {
	NewTinyExternal()
}