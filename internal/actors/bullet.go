package actors

import (
	"fmt"

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
		Health:    1,
		Speed:     speed,
		direction: direction,
	}
	a.RegisterNode(&bullet)
	a.CollisionServer.RegisterNode(&bullet)
	bullet.OnStart()

	return &bullet
}

type Bullet struct {
	Position dsu.Vector2i
	Texture  *sdl.Texture
	Speed    int32
	Health   int

	app       *core.App
	direction dsu.Vector2i
}

func (b *Bullet) OnStart() {
}

func (b *Bullet) OnUpdate() {
}

func (b *Bullet) OnRender(r *sdl.Renderer) {
	if b.Position.X < core.ScreenWidth {
		graphics.Blit(r, b.Texture, b.Position, 1.0)
		b.Position = b.Position.Add(b.direction.MultiplyScalar(b.Speed))
		// TODO: Find out a way to free this memory
	}
}

func (b *Bullet) OnKeyDown(key *sdl.KeyboardEvent) {
}

func (b *Bullet) OnKeyUp(key *sdl.KeyboardEvent) {
}

func (b *Bullet) OnCollided(collider any) {
	// TODO: PROPER HANDLING
	fmt.Println("HIT AN ENEMY")
}

func (b *Bullet) GetMetadataForCollision() (int32, int32, int32, int32) {
	_, _, width, height, err := b.Texture.Query()
	if err != nil {
		panic(err)
	}

	return b.Position.X, b.Position.Y, width, height
}
