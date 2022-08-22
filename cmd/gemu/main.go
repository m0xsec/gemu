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
package main

import (
	"fmt"
	"gemu/pkg/gb"
	"gemu/pkg/render"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	fmt.Println("gemu")

	// Setup communication channels
	renderFrame := make(chan *sdl.Surface, 4)
	renderStopped := make(chan struct{})
	stopRender := make(chan struct{})
	gbStopped := make(chan struct{})
	stopGB := make(chan struct{})

	// Initialize SDL
	render.Init()

	// Initialize GameBoy
	gemu := gb.GameBoy{}
	gemu.Init(renderFrame)

	// Launch Renderer and Emulator :3
	go render.Run(renderFrame, renderStopped, stopRender)
	go gemu.Run(gbStopped, stopGB)

	// Wait to close, gracefully <3
	select {
	case <-renderStopped:
		close(stopGB)
		<-gbStopped

	case <-gbStopped:
		close(stopRender)
		<-renderStopped
	}
}
