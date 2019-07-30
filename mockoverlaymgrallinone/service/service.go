package service

import (
	"github.com/andyliuliming/gogomodule/mockexternals/external1"
	"github.com/andyliuliming/gogomodule/mockexternals/external2"
)

func NewExt1() *external1.External1{
	return &external1.External1{
	}
}

func NewExt2() *external2.External2{
	return &external2.External2{
	}
}


