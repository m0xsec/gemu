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
	if err := gemu.Init(renderFrame); err != nil {
		fmt.Println("[!] gemu init failed - " + err.Error())
		return
	}

	// Launch Renderer and Emulator :3
	go func() {
		err := render.Run(renderFrame, renderStopped, stopRender)
		if err != nil {
			fmt.Println("[!] render routine failed - " + err.Error())
			close(renderStopped)
		}
	}()
	go func() {
		err := gemu.Run(gbStopped, stopGB)
		if err != nil {
			fmt.Println("[!] gemu routine failed - " + err.Error())
			close(gbStopped)
		}
	}()

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
