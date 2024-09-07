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
	layer int,
	targetLayer int,
) *Bullet {
	bullet := Bullet{
		Layer:       layer,
		TargetLayer: targetLayer,
		Position:    position,
		Texture:     texture,
		Health:      1,
		Speed:       speed,
		direction:   direction,
		app:         a,
	}
	a.RegisterNode(&bullet)
	a.CollisionServer.RegisterNode(&bullet)
	bullet.OnStart()

	return &bullet
}

type Bullet struct {
	Position    dsu.Vector2i
	Texture     *sdl.Texture
	Speed       int32
	Health      int
	Layer       int
	TargetLayer int

	app       *core.App
	direction dsu.Vector2i
}

func (b *Bullet) OnStart() {
}

func (b *Bullet) OnUpdate() {
}

func (b *Bullet) OnRender(r *sdl.Renderer) {
	if b.Position.X < core.ScreenWidth && b.Position.X > -30 {
		graphics.Blit(r, b.Texture, b.Position, 1.0)
		b.Position = b.Position.Add(b.direction.MultiplyScalar(b.Speed))
	} else {
		b.Free()
	}
}

func (b *Bullet) OnKeyDown(key *sdl.KeyboardEvent) {
}

func (b *Bullet) OnKeyUp(key *sdl.KeyboardEvent) {
}

func (b *Bullet) OnCollided(area *core.CollisionArea) {
	b.Free()
}

func (b *Bullet) GetMetadataForCollision() (int32, int32, int32, int32) {
	_, _, width, height, err := b.Texture.Query()
	if err != nil {
		panic(err)
	}

	return b.Position.X, b.Position.Y, width, height
}

func (b *Bullet) GetLayer() int {
	return b.Layer
}

func (b *Bullet) GetTargetLayer() int {
	return b.TargetLayer
}

func (b *Bullet) GetTag() string {
	return "bullet"
}

func (b *Bullet) Free() {
	for i, e := range b.app.MainHooks {
		if *e == b {
			b.app.MainHooks[i] = b.app.MainHooks[len(b.app.MainHooks)-1]
			b.app.MainHooks = b.app.MainHooks[:len(b.app.MainHooks)-1]
			break
		}
	}
	for i, e := range b.app.KeyboardInputHooks {
		if *e == b {
			b.app.KeyboardInputHooks[i] = b.app.KeyboardInputHooks[len(b.app.KeyboardInputHooks)-1]
			b.app.KeyboardInputHooks = b.app.KeyboardInputHooks[:len(b.app.KeyboardInputHooks)-1]
		}
	}
	for i, e := range b.app.CollisionServer.Colliders {
		if *e == b {
			b.app.CollisionServer.Colliders[i] = b.app.CollisionServer.Colliders[len(b.app.CollisionServer.Colliders)-1]
			b.app.CollisionServer.Colliders = b.app.CollisionServer.Colliders[:len(b.app.CollisionServer.Colliders)-1]
		}
	}
}
