package actors

import (
	"math/rand"

	"Lunarisnia/sdl-pong/internal/core"
	"Lunarisnia/sdl-pong/internal/dsu"
	"github.com/veandco/go-sdl2/sdl"
)

func NewEnemySpawner(
	a *core.App,
	interval int,
	enemyTexture *sdl.Texture,
	bulletTexture *sdl.Texture,
) *EnemySpawner {
	enemySpawner := EnemySpawner{
		Interval:      interval,
		EnemyTexture:  enemyTexture,
		BulletTexture: bulletTexture,

		app: a,
	}
	a.RegisterNode(&enemySpawner)
	return &enemySpawner
}

type EnemySpawner struct {
	Interval      int
	EnemyTexture  *sdl.Texture
	BulletTexture *sdl.Texture

	app   *core.App
	timer int
}

func (e *EnemySpawner) OnRender(r *sdl.Renderer) {
}

func (e *EnemySpawner) OnStart() {
	e.timer = e.Interval
}

func (e *EnemySpawner) OnUpdate() {
	if e.timer <= 0 {
		_, _, _, height, err := e.EnemyTexture.Query()
		if err != nil {
			panic(err)
		}
		NewEnemy(e.app, dsu.Vector2i{
			X: core.ScreenWidth,
			Y: int32(rand.Intn(int(core.ScreenHeight - height*2.0))),
		}, e.EnemyTexture, e.BulletTexture, 5, 2, 0)
		e.timer = e.Interval
		return
	}

	e.timer--
}
