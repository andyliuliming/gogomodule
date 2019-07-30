package mockoverlaymgrallinone
import (
	"github.com/andyliuliming/gogomodule/mockexternals/external1"
	"github.com/andyliuliming/gogomodule/mockexternals/external2"
	"github.com/andyliuliming/gogomodule/mockexternals/tinyexternal"
)

func NewExt1() *external1.External1{
	return &external1.External1{
	}
}

func NewExt2() *external2.External2{
	return &external2.External2{
	}
}

func NewTinyExternal() *tinyexternal.TinyExternal{
	return &tinyexternal.TinyExternal{
	}
}

