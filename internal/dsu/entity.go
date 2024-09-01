package dsu

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Node interface {
	OnStart()
	OnUpdate()
}

type NodeInput interface {
	OnKeyDown(key *sdl.KeyboardEvent)
	OnKeyUp(key *sdl.KeyboardEvent)
}

type NodeRender interface {
	OnRender(r *sdl.Renderer)
}
