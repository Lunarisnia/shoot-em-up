package actors

import (
	"Lunarisnia/sdl-pong/internal/core"
	"Lunarisnia/sdl-pong/internal/dsu"
	"Lunarisnia/sdl-pong/internal/graphics"
	"github.com/veandco/go-sdl2/sdl"
)

func NewBullet(
	a *core.App,
	position dsu.Vector2i,
	texture *sdl.Texture,
	speed int32,
	direction dsu.Vector2i,
) *Bullet {
	bullet := Bullet{
		Position:  position,
		Texture:   texture,
		Speed:     speed,
		direction: direction,
	}
	a.RegisterNode(&bullet)

	return &bullet
}

type Bullet struct {
	Position dsu.Vector2i
	Texture  *sdl.Texture
	Speed    int32

	app       *core.App
	direction dsu.Vector2i
}

func (b *Bullet) OnStart() {
}

func (b *Bullet) OnUpdate(r *sdl.Renderer) {
	graphics.Blit(r, b.Texture, b.Position, 1.0)

	b.Position = b.Position.Add(b.direction.MultiplyScalar(b.Speed))
}

func (b *Bullet) OnKeyDown(key *sdl.KeyboardEvent) {
}

func (b *Bullet) OnKeyUp(key *sdl.KeyboardEvent) {
}
