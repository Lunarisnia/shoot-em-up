package core

import (
	"Lunarisnia/sdl-pong/internal/dsu"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

type App struct {
	Window          *sdl.Window
	Renderer        *sdl.Renderer
	CollisionServer *CollisionServer

	MainHooks          []*dsu.Node
	KeyboardInputHooks []*dsu.NodeInput
}

func (a *App) InitSDL() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	var err error
	a.Window, err = sdl.CreateWindow(
		"main",
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		ScreenWidth,
		ScreenHeight,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		panic(err)
	}

	a.Renderer, err = sdl.CreateRenderer(a.Window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	a.CollisionServer = NewCollisionServer()
}

func (a *App) RegisterNode(e interface{}) {
	if ev, ok := e.(dsu.Node); ok {
		a.MainHooks = append(a.MainHooks, &ev)
	}
	if ev, ok := e.(dsu.NodeInput); ok {
		a.KeyboardInputHooks = append(a.KeyboardInputHooks, &ev)
	}
}

func (a *App) Updates() {
	for _, event := range a.MainHooks {
		(*event).OnUpdate()
	}
}

func (a *App) Renders(r *sdl.Renderer) {
	for _, event := range a.MainHooks {
		(*event).OnRender(r)
	}
}

func (a *App) Starts() {
	for _, event := range a.MainHooks {
		(*event).OnStart()
	}
}

func (a *App) KeyboardInputs(key *sdl.KeyboardEvent) {
	for _, event := range a.KeyboardInputHooks {
		if key.GetType() == sdl.KEYDOWN {
			(*event).OnKeyDown(key)
		} else {
			(*event).OnKeyUp(key)
		}
	}
}
