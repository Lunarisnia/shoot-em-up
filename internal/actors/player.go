package actors

import (
	"fmt"

	"Lunarisnia/sdl-pong/internal/core"
	"Lunarisnia/sdl-pong/internal/dsu"
	"Lunarisnia/sdl-pong/internal/graphics"
	"github.com/veandco/go-sdl2/sdl"
)

func NewPlayer(a *core.App, position dsu.Vector2i, texture *sdl.Texture) *Player {
	bulletTexture, err := graphics.LoadTexture(a.Renderer, "assets/bullet.png")
	if err != nil {
		panic(err)
	}
	player := Player{
		app:           a,
		Position:      position,
		Texture:       texture,
		Speed:         8,
		scale:         2.0,
		bulletTexture: bulletTexture,
	}
	a.RegisterNode(&player)

	return &player
}

type movementInput struct {
	Up    int
	Down  int
	Left  int
	Right int
}

type Player struct {
	Position dsu.Vector2i
	Texture  *sdl.Texture
	Speed    int32

	movementInput movementInput
	direction     dsu.Vector2i
	app           *core.App
	scale         float32
	bulletTexture *sdl.Texture
	firing        bool
	reloadSpeed   int
}

func (p *Player) OnStart() {
	fmt.Println("This start")
}

func (p *Player) OnUpdate() {
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

	if p.firing {
		p.spawnBullet()
	}

	if p.movementInput.Up == 1 {
		fmt.Println("MovementInput: ", p.movementInput)
		p.direction.Y = -1
	} else if p.movementInput.Down == 1 {
		p.direction.Y = 1
	} else {
		p.direction.Y = 0
	}
	if p.movementInput.Left == 1 {
		p.direction.X = -1
	} else if p.movementInput.Right == 1 {
		p.direction.X = 1
	} else {
		p.direction.X = 0
	}
}

func (p *Player) OnRender(r *sdl.Renderer) {
	graphics.Blit(r, p.Texture, p.Position, p.scale)
}

func (p *Player) OnKeyDown(key *sdl.KeyboardEvent) {
	if key.Keysym.Scancode == sdl.SCANCODE_W {
		p.movementInput.Up = 1
	} else if key.Keysym.Scancode == sdl.SCANCODE_S {
		p.movementInput.Down = 1
	}
	if key.Keysym.Scancode == sdl.SCANCODE_A {
		p.movementInput.Left = 1
	} else if key.Keysym.Scancode == sdl.SCANCODE_D {
		p.movementInput.Right = 1
	}

	if key.Keysym.Scancode == sdl.SCANCODE_SPACE {
		p.firing = true
	}
}

func (p *Player) OnKeyUp(key *sdl.KeyboardEvent) {
	if key.Keysym.Scancode == sdl.SCANCODE_W {
		p.movementInput.Up = 0
	} else if key.Keysym.Scancode == sdl.SCANCODE_S {
		p.movementInput.Down = 0
	}
	if key.Keysym.Scancode == sdl.SCANCODE_A {
		p.movementInput.Left = 0
	} else if key.Keysym.Scancode == sdl.SCANCODE_D {
		p.movementInput.Right = 0
	}

	if key.Keysym.Scancode == sdl.SCANCODE_SPACE {
		p.firing = false
	}
}

func (p *Player) spawnBullet() {
	_, _, width, height, err := p.Texture.Query()
	if err != nil {
		panic(err)
	}
	_, _, bulletWidth, bulletHeight, err := p.bulletTexture.Query()
	if err != nil {
		panic(err)
	}
	bulletSpawnPosition := dsu.Vector2i{
		X: p.Position.X,
		Y: p.Position.Y,
	}
	bulletOffset := dsu.Vector2i{
		X: (width*int32(p.scale) + 30) - bulletWidth/2,
		Y: (height * int32(p.scale) / 2) - bulletHeight/2,
	}

	bulletSpawnPosition = bulletSpawnPosition.Add(bulletOffset)
	NewBullet(p.app, bulletSpawnPosition, p.bulletTexture, 10.0, dsu.Vector2i{X: 1, Y: 0})
}
