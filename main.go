package main

import (
	"Lunarisnia/sdl-pong/internal/actors"
	"Lunarisnia/sdl-pong/internal/core"
	"Lunarisnia/sdl-pong/internal/dsu"
	"Lunarisnia/sdl-pong/internal/graphics"
	"Lunarisnia/sdl-pong/internal/inputs"
	"github.com/veandco/go-sdl2/sdl"
)

func initNodes(a *core.App) {
	playerSprite, err := graphics.LoadTexture(a.Renderer, "assets/player.png")
	if err != nil {
		panic(err)
	}
	_, _, width, height, err := playerSprite.Query()
	if err != nil {
		panic(err)
	}

	actors.NewPlayer(a,
		dsu.Vector2i{
			X: core.ScreenWidth/2 - width*2/2,
			Y: core.ScreenHeight/2 - height*2/2,
		},
		playerSprite,
	)
}

func main() {
	app := &core.App{}
	app.InitSDL()
	defer sdl.Quit()
	defer app.Window.Destroy()
	defer app.Renderer.Destroy()

	initNodes(app)

	app.Starts()

	tick := float64(sdl.GetTicks64())
	remainder := float64(0.0)

	running := true
	for running {
		graphics.PrepareScene(app.Renderer)

		app.Renders(app.Renderer)

		inputs.HandleInput(func() {
			running = false
		}, func(key *sdl.KeyboardEvent) {
			app.KeyboardInputs(key)
		})

		app.Updates()

		graphics.PresentScene(app.Renderer)

		capFramerate(&tick, &remainder)
	}
}

func capFramerate(previousTick *float64, remainder *float64) {
	var wait float64
	var latestTick float64

	wait = 16 + *remainder

	*remainder -= *remainder

	latestTick = float64(sdl.GetTicks64()) - *previousTick

	wait -= latestTick

	if wait < 1.0 {
		wait = 1.0
	}

	sdl.Delay(uint32(wait))

	*remainder += 0.667

	*previousTick = float64(sdl.GetTicks64())
}
