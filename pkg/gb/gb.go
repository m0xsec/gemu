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
	"gemu/pkg/cpu"
	"gemu/pkg/mmu"
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
}

// Init initializes the GameBoy, bringing subsystems online
func (gb *GameBoy) Init() {
	gb.cpu = new(cpu.CPU)
	gb.mmu = new(mmu.MMU)

	gb.cpu.Init(gb.mmu)
}

// Cycle represents a single GameBoy CPU/Emulation Cycle (Fetch/Decode/Execute)
func (gb *GameBoy) Cycle() error {
	// Fetch, Decode, and Execute
	return gb.cpu.Step()

	// TODO: other stuff will happen here, of course...
}
