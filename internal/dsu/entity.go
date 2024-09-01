package dsu

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Node interface {
	OnStart()
	OnUpdate(r *sdl.Renderer)
}

type NodeInput interface {
	OnKeyDown(key *sdl.KeyboardEvent)
	OnKeyUp(key *sdl.KeyboardEvent)
}
