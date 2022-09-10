/*
		gemu - the gameboy emulator
				<3 m0x
	    __________________________
	   |                          |
	   | .----------------------. |
	   | |  .----------------.  | |
	   | |  |                |  | |
	   | |))|                |  | |
	   | |  |                |  | |
	   | |  |                |  | |
	   | |  |                |  | |
	   | |  |                |  | |
	   | |  |                |  | |
	   | |  '----------------'  | |
	   | |__GAME BOY____________/ |
	   |          ________        |
	   |    .    (Nintendo)       |
	   |  _| |_   """"""""   .-.  |
	   |-[_   _]-       .-. (   ) |
	   |   |_|         (   ) '-'  |
	   |    '           '-'   A   |
	   |                 B        |
	   |          ___   ___       |
	   |         (___) (___)  ,., |
	   |        select start ;:;: |
	   |                    ,;:;' /
	   |                   ,:;:'.'
	   '-----------------------`
*/
package render

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Init will start SDL and related subsystems
func Init() error {
	// Initialize SDL2
	fmt.Println("Initializing SDL2...")
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return err
	}

	// Initialize SDL2 TTF
	fmt.Println("Initializing SDL2 TTF...")
	err := ttf.Init()
	if err != nil {
		fmt.Println("Failed to initialize TTF: " + err.Error())
		return err
	}

	return nil
}

// Run starts the rendering loop, which handles SDL events and renders the gameboy screen
func Run(frame chan *sdl.Surface, renderStopped chan struct{}, stopRender chan struct{}) error {
	//runtime.LockOSThread()
	rendering := true

	// Define rendering FPS limits
	// 60hz is pretty close to the 59.73Hz vertical sync of the Gameboy
	fps := uint64(60)         // Frame per second maxmimum
	tpp := uint64(1000 / fps) // Ticks per frame

	// Create SDL2 window
	window, err := sdl.CreateWindow("gemu", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 600, 800, sdl.WINDOW_SHOWN)
	if err != nil {
		return err
	}

	// Create SDL2 renderer for our window, without V-Sync (since we want to control the FPS)
	//renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE) // WSL2 doesn't support hardware acceleration
	if err != nil {
		return err
	}

	// Define our SDL Surface, which will represent the rendered version of the Gameboy's VRAM
	frameSurf := new(sdl.Surface)

	// Stop channel monitoring
	go func(stopped chan struct{}, stop chan struct{}) {
		<-stop

		// Cleanup
		window.Destroy()
		renderer.Destroy()
		ttf.Quit()
		sdl.Quit()

		rendering = false
		//runtime.UnlockOSThread()
		close(stopped)
	}(renderStopped, stopRender)

	// Rendering loop
	for rendering {
		// Make note of when this frame iternation started, so we can control FPS
		frameStart := sdl.GetTicks64()

		// Get the next frame from the emulator, in a non-blocking way
		select {
		case f := <-frame:
			frameSurf = f
		default:
			//fmt.Println("[RENDER] NO FRAME IN CHAN")
		}

		// Create our frame texture
		texture, err := renderer.CreateTextureFromSurface(frameSurf)
		if err != nil {
			return err
		}
		texture.SetBlendMode(sdl.BLENDMODE_BLEND)
		texture.SetAlphaMod(200)

		// Update our renderer with the new frame texture
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()
		renderer.Copy(texture, nil, &sdl.Rect{X: 150, Y: 250, W: 200, H: 20})
		renderer.Present()
		texture.Destroy()

		// How long did that take? Do we need to delay to maintain 60fps
		frameTime := sdl.GetTicks64() - frameStart
		if frameTime < tpp {
			//fmt.Println("Render delay to maintain 60fps")
			sdl.Delay(uint32(tpp - frameTime))
		}

		// Event handling
		// TODO: Event handling should probably go in its own routine...
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				println("kthxbai<3")
				close(renderStopped)
				rendering = false

			case *sdl.KeyboardEvent:
				switch t.Keysym.Sym {
				case sdl.K_ESCAPE:
					println("kthxbai<3")
					close(renderStopped)
					rendering = false
				}
			}
		}
	}
	return nil
}
