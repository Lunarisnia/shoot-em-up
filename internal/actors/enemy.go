package actors

import (
	"Lunarisnia/sdl-pong/internal/core"
	"Lunarisnia/sdl-pong/internal/dsu"
	"Lunarisnia/sdl-pong/internal/graphics"
	"github.com/veandco/go-sdl2/sdl"
)

func NewEnemy(
	a *core.App,
	position dsu.Vector2i,
	texture *sdl.Texture,
	bulletTexture *sdl.Texture,
) *Enemy {
	enemy := Enemy{
		app:           a,
		Position:      position,
		Texture:       texture,
		Speed:         8,
		scale:         2.0,
		bulletTexture: bulletTexture,
	}
	a.RegisterNode(&enemy)
	return &enemy
}

type Enemy struct {
	Position dsu.Vector2i
	Texture  *sdl.Texture
	Speed    int32

	app           *core.App
	direction     dsu.Vector2i
	bulletTexture *sdl.Texture
	scale         float32
}

func (e *Enemy) OnStart() {
}

func (e *Enemy) OnUpdate() {
	e.Position.X -= 1
}

func (e *Enemy) OnRender(r *sdl.Renderer) {
	graphics.Blit(r, e.Texture, e.Position, e.scale)
}
