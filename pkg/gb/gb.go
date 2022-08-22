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
package gb

import (
	"fmt"
	"gemu/pkg/cpu"
	"gemu/pkg/mmu"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// GameBoy represents the GameBoy hardware
type GameBoy struct {
	// The heart of the Gameboy, the CPU.
	// The CPU is responsible for decoding and executing instructions.
	// The DMG-01 had a Sharp LR35902 CPU (speculated to be a SM83 core), which is a hybrid of the Z80 and the 8080.
	cpu *cpu.CPU

	// The DMG-01 didn't have an actual Memory Management Unit (MMU), but it had a memory-mapped I/O system with a single RAM chip.
	// To make emulation easier, we will define a MMU.
	// The MMU is responsible for mapping memory addresses to actual memory locations.
	mmu *mmu.MMU

	// nextFrame represents the SDL Texture channel that will be used by the renderer to display the Gameboy screen
	nextFrame chan *sdl.Surface

	// Temp message to display while the PPU is not implemented
	ppuWarning *sdl.Surface
}

// Run will start up the Gameboy Emulator
func (gb *GameBoy) Run(gbStopped chan struct{}, stopGB chan struct{}) error {
	emulating := true

	// Stop channel monitoring
	go func(stopped chan struct{}, stop chan struct{}) {
		<-stop
		emulating = false
		close(stopped)
	}(gbStopped, stopGB)

	// GameBoy CPU Cycle
	for emulating {
		err := gb.cycle()
		if err != nil {
			//fmt.Printf("[GB Cycle] Error: %s\n", err)
		}

	}

	return nil
}

// Init initializes the GameBoy, bringing subsystems online
func (gb *GameBoy) Init(nextFrame chan *sdl.Surface) {
	// Setup a temporary warning message to indicate that the PPU is not yet implemented
	font, err := ttf.OpenFont("./fonts/GameBoy.ttf", 20)
	if err != nil {
		fmt.Println("Failed to load font: " + err.Error())
	}

	gb.ppuWarning, err = font.RenderUTF8Blended("PPU NOT YET IMPLEMENTED", sdl.Color{R: 255, G: 255, B: 255, A: 255})
	if err != nil {
		fmt.Println("Failed to render surface: " + err.Error())
	}
	font.Close()

	// Setup Gameboy subsystems
	gb.cpu = new(cpu.CPU)
	gb.mmu = new(mmu.MMU)
	gb.nextFrame = nextFrame

	// Init Gameboy subsystems <3
	gb.cpu.Init(gb.mmu)
}

// Cycle represents a single GameBoy CPU/Emulation Cycle (Fetch/Decode/Execute)
func (gb *GameBoy) cycle() error {
	// Fetch, Decode, and Execute
	err := gb.cpu.Step()

	if err != nil {
		return err
	}

	// WARNING: This is a blocking operation !!
	// As long as the emulator doesn't run too fast, it shouldn't matter.
	select {
	case gb.nextFrame <- gb.ppuWarning:
		//fmt.Println("SENT FRAME")
	default:
		//fmt.Println("FRAME CHAN BLOCK")
	}
	//gb.nextFrame <- gb.ppuWarning

	// TODO: other stuff will happen here, of course...

	return nil
}
