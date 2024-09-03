package actors

import (
	"math/rand"

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
	enemy.OnStart()
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
	e.direction.X = -1
	e.direction.Y = int32(rand.Intn(3) - 1)
}

func (e *Enemy) OnUpdate() {
	_, _, _, height, err := e.Texture.Query()
	if err != nil {
		panic(err)
	}
	newPosition := e.Position.Add(e.direction.MultiplyScalar(5))
	if newPosition.Y < 0 || newPosition.Y > (core.ScreenHeight-height*2) {
		e.direction.Y *= -1
	}
	e.Position = newPosition
}

func (e *Enemy) OnRender(r *sdl.Renderer) {
	graphics.Blit(r, e.Texture, e.Position, e.scale)
}
