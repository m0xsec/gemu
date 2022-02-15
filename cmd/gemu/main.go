/*
	 ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄       ▄▄  ▄         ▄
	▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░▌     ▐░░▌▐░▌       ▐░▌
	▐░█▀▀▀▀▀▀▀▀▀ ▐░█▀▀▀▀▀▀▀▀▀ ▐░▌░▌   ▐░▐░▌▐░▌       ▐░▌
	▐░▌          ▐░▌          ▐░▌▐░▌ ▐░▌▐░▌▐░▌       ▐░▌
	▐░▌ ▄▄▄▄▄▄▄▄ ▐░█▄▄▄▄▄▄▄▄▄ ▐░▌ ▐░▐░▌ ▐░▌▐░▌       ▐░▌
	▐░▌▐░░░░░░░░▌▐░░░░░░░░░░░▌▐░▌  ▐░▌  ▐░▌▐░▌       ▐░▌
	▐░▌ ▀▀▀▀▀▀█░▌▐░█▀▀▀▀▀▀▀▀▀ ▐░▌   ▀   ▐░▌▐░▌       ▐░▌
	▐░▌       ▐░▌▐░▌          ▐░▌       ▐░▌▐░▌       ▐░▌
	▐░█▄▄▄▄▄▄▄█░▌▐░█▄▄▄▄▄▄▄▄▄ ▐░▌       ▐░▌▐░█▄▄▄▄▄▄▄█░▌
	▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░▌       ▐░▌▐░░░░░░░░░░░▌
	▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀         ▀  ▀▀▀▀▀▀▀▀▀▀▀
					the GameBoy Emulator
							m0x <3
*/
package main

import (
	"fmt"
	"gemu/pkg/gb"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {
	fmt.Println("gemu")

	// TODO: Initialize GameBoy
	// Subsystems, etc, etc
	gemu := gb.GameBoy{}
	gemu.Init()

	// Initialize SDL2
	fmt.Println("Initializing SDL2...")
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	// Initialize SDL2 TTF
	fmt.Println("Initializing SDL2 TTF...")
	err := ttf.Init()
	if err != nil {
		fmt.Println("Failed to initialize TTF: " + err.Error())
	}

	// Create SDL2 window
	window, err := sdl.CreateWindow("gemu", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 600, 800, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	// Create SDL2 renderer
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	// Emulator loop
	emulating := true
	for emulating {
		// TODO: GameBoy CPU/Emulation Cycle (Fetch/Decode/Execute)

		//renderer.SetDrawColor(0, 0, 0, 255)
		//renderer.Clear()

		// TODO: Rendering

		// Render to screen <3
		//renderer.Present()

		// Event handling
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				println("kthxbai<3")
				emulating = false

			case *sdl.KeyboardEvent:
				switch t.Keysym.Sym {
				case sdl.K_ESCAPE:
					println("kthxbai<3")
					emulating = false
				}
			}
		}

		// TODO: Maintain FPS/Timers/Cycle Rate?
	}
}
