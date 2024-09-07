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
	health int,
	layer int,
	targetLayer int,
) *Enemy {
	enemy := Enemy{
		Layer:         layer,
		TargetLayer:   targetLayer,
		app:           a,
		Position:      position,
		Texture:       texture,
		Speed:         8,
		scale:         2.0,
		Health:        health,
		bulletTexture: bulletTexture,
	}
	a.RegisterNode(&enemy)
	a.CollisionServer.RegisterNode(&enemy)
	enemy.OnStart()
	return &enemy
}

type Enemy struct {
	Position    dsu.Vector2i
	Texture     *sdl.Texture
	Speed       int32
	Index       int
	Health      int
	Layer       int
	TargetLayer int

	app           *core.App
	direction     dsu.Vector2i
	bulletTexture *sdl.Texture
	scale         float32
	shouldShoot   int
}

func (e *Enemy) OnStart() {
	e.direction.X = -1
	e.direction.Y = int32(rand.Intn(3) - 1)

	e.shouldShoot = 30
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

	e.shouldShoot--
	if e.shouldShoot <= 0 {
		e.shouldShoot = 30
		e.spawnBullet()
	}
}

func (e *Enemy) OnRender(r *sdl.Renderer) {
	if e.Position.X > -150 {
		graphics.Blit(r, e.Texture, e.Position, e.scale)
	} else {
		e.Free()
	}
}

func (e *Enemy) OnHit(collider *core.Collider) {
	if (*collider).GetTag() == "player" {
		e.Free()
		return
	}
	e.Health--
	if e.Health < 1 {
		e.Free()
	}
}

func (e *Enemy) GetMetadataForCollision() (int32, int32, int32, int32) {
	_, _, width, height, err := e.Texture.Query()
	if err != nil {
		panic(err)
	}

	return e.Position.X, e.Position.Y, width * 2, height * 2
}

func (e *Enemy) GetLayer() int {
	return e.Layer
}

func (e *Enemy) GetTargetLayer() int {
	return e.TargetLayer
}

func (e *Enemy) GetTag() string {
	return "enemy"
}

func (e *Enemy) Free() {
	for i, ev := range e.app.MainHooks {
		if *ev == e {
			e.app.MainHooks[i] = e.app.MainHooks[len(e.app.MainHooks)-1]
			e.app.MainHooks = e.app.MainHooks[:len(e.app.MainHooks)-1]
			break
		}
	}
	for i, ev := range e.app.CollisionServer.CollisionAreas {
		if *ev == e {
			e.app.CollisionServer.CollisionAreas[i] = e.app.CollisionServer.CollisionAreas[len(e.app.CollisionServer.CollisionAreas)-1]
			e.app.CollisionServer.CollisionAreas = e.app.CollisionServer.CollisionAreas[:len(e.app.CollisionServer.CollisionAreas)-1]
		}
	}
}

func (e *Enemy) spawnBullet() {
	_, _, width, height, err := e.Texture.Query()
	if err != nil {
		panic(err)
	}
	_, _, bulletWidth, bulletHeight, err := e.bulletTexture.Query()
	if err != nil {
		panic(err)
	}
	bulletSpawnPosition := dsu.Vector2i{
		X: e.Position.X,
		Y: e.Position.Y,
	}
	scatter := int32(rand.Intn(60+60) - 60)
	bulletOffset := dsu.Vector2i{
		X: (width*int32(e.scale) - 30) - bulletWidth/2,
		Y: (height * int32(e.scale) / 2) - bulletHeight/2 + scatter,
	}

	bulletSpawnPosition = bulletSpawnPosition.Add(bulletOffset)
	NewBullet(e.app, bulletSpawnPosition, e.bulletTexture, 7.0, dsu.Vector2i{X: -1, Y: 0}, 2, 0)
}
