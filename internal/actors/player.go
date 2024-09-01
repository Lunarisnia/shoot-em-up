package actors

import (
	"fmt"

	"Lunarisnia/sdl-pong/internal/core"
	"Lunarisnia/sdl-pong/internal/dsu"
	"Lunarisnia/sdl-pong/internal/graphics"
	"github.com/veandco/go-sdl2/sdl"
)

func NewPlayer(a *core.App, position dsu.Vector2i, texture *sdl.Texture) *Player {
	player := Player{
		app:      a,
		Position: position,
		Texture:  texture,
		Speed:    8,
		scale:    2.0,
	}
	a.RegisterNode(&player)

	return &player
}

type Player struct {
	Position  dsu.Vector2i
	Texture   *sdl.Texture
	Speed     int32
	direction dsu.Vector2i
	app       *core.App
	scale     float32
}

func (p *Player) OnStart() {
	fmt.Println("This start")
}

func (p *Player) OnUpdate(r *sdl.Renderer) {
	graphics.Blit(r, p.Texture, p.Position, p.scale)

	_, _, width, height, err := p.Texture.Query()
	if err != nil {
		panic(err)
	}
	newPos := p.Position.Add(p.direction.MultiplyScalar(p.Speed))
	spritePosX := newPos.X + (width * int32(p.scale))
	spritePosY := newPos.Y + (height * int32(p.scale))
	if newPos.X >= 0 && spritePosX < core.ScreenWidth && newPos.Y >= 0 &&
		spritePosY < core.ScreenHeight {
		p.Position = newPos
	}
}

func (p *Player) OnKeyDown(key *sdl.KeyboardEvent) {
	if key.Keysym.Scancode == sdl.SCANCODE_W {
		p.direction.Y = -1
	} else if key.Keysym.Scancode == sdl.SCANCODE_S {
		p.direction.Y = 1
	}
	if key.Keysym.Scancode == sdl.SCANCODE_A {
		p.direction.X = -1
	} else if key.Keysym.Scancode == sdl.SCANCODE_D {
		p.direction.X = 1
	}

	if key.Keysym.Scancode == sdl.SCANCODE_SPACE {
		p.spawnBullet(p.app.Renderer)
	}
}

func (p *Player) OnKeyUp(key *sdl.KeyboardEvent) {
	if key.Keysym.Scancode == sdl.SCANCODE_W || key.Keysym.Scancode == sdl.SCANCODE_S {
		p.direction.Y = 0
	}
	if key.Keysym.Scancode == sdl.SCANCODE_A || key.Keysym.Scancode == sdl.SCANCODE_D {
		p.direction.X = 0
	}
}

func (p *Player) spawnBullet(r *sdl.Renderer) {
	// TODO: Draw my own bullet
	bulletSurface, err := sdl.CreateRGBSurface(0, 32, 32, 32, 0, 0, 0, 0)
	if err != nil {
		panic(err)
	}
	bulletSprite, err := r.CreateTextureFromSurface(bulletSurface)
	if err != nil {
		panic(err)
	}
	_, _, width, height, err := p.Texture.Query()
	if err != nil {
		panic(err)
	}
	_, _, bulletWidth, bulletHeight, err := bulletSprite.Query()
	if err != nil {
		panic(err)
	}
	NewBullet(p.app, dsu.Vector2i{
		X: p.Position.X + (width*int32(p.scale) + 30) - bulletWidth/2,
		Y: p.Position.Y + (height * int32(p.scale) / 2) - bulletHeight/2,
	}, bulletSprite, 10.0, dsu.Vector2i{X: 1, Y: 0})
}
