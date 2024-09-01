package inputs

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func HandleInput(quit func(), onInput func(key *sdl.KeyboardEvent)) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.KeyboardEvent:
			keyEvent := event.(*sdl.KeyboardEvent)
			onInput(keyEvent)
		case *sdl.QuitEvent:
			fmt.Println("Quit")
			quit()
		default:
			break
		}
	}
}
