package actors

import (
	"fmt"

	"Lunarisnia/sdl-pong/internal/core"
	"Lunarisnia/sdl-pong/internal/dsu"
	"Lunarisnia/sdl-pong/internal/graphics"
	"github.com/veandco/go-sdl2/sdl"
)

type Movement struct {
	Direction dsu.Vector2i
}

var movement Movement = Movement{}

func NewPlayer(a *core.App, position dsu.Vector2i, texture *sdl.Texture) *Player {
	player := Player{
		Position: position,
		Texture:  texture,
		Speed:    8,
	}
	a.RegisterNode(&player)

	return &player
}

type Player struct {
	Position dsu.Vector2i
	Texture  *sdl.Texture
	Speed    int32
}

func (p *Player) OnStart() {
	fmt.Println("This start")
}

func (p *Player) OnUpdate() {
	_, _, width, height, err := p.Texture.Query()
	if err != nil {
		panic(err)
	}
	newPos := p.Position.Add(movement.Direction.MultiplyScalar(p.Speed))
	spritePosX := newPos.X + (width * 10.0)
	spritePosY := newPos.Y + (height * 10.0)
	if newPos.X >= 0 && spritePosX < core.ScreenWidth && newPos.Y >= 0 &&
		spritePosY < core.ScreenHeight {
		p.Position = newPos
	}
}

func (p *Player) OnRender(r *sdl.Renderer) {
	graphics.Blit(r, p.Texture, p.Position, 10.0)
}

func (p *Player) OnKeyDown(key *sdl.KeyboardEvent) {
	if key.Keysym.Scancode == sdl.SCANCODE_W {
		movement.Direction.Y = -1
	} else if key.Keysym.Scancode == sdl.SCANCODE_S {
		movement.Direction.Y = 1
	}
	if key.Keysym.Scancode == sdl.SCANCODE_A {
		movement.Direction.X = -1
	} else if key.Keysym.Scancode == sdl.SCANCODE_D {
		movement.Direction.X = 1
	}
}

func (p *Player) OnKeyUp(key *sdl.KeyboardEvent) {
	if key.Keysym.Scancode == sdl.SCANCODE_W || key.Keysym.Scancode == sdl.SCANCODE_S {
		movement.Direction.Y = 0
	}
	if key.Keysym.Scancode == sdl.SCANCODE_A || key.Keysym.Scancode == sdl.SCANCODE_D {
		movement.Direction.X = 0
	}
}
