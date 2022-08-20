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
	cpu *cpu.CPU
	mmu *mmu.MMU
}

func (gb *GameBoy) Init() {
	gb.cpu = new(cpu.CPU)
	gb.mmu = new(mmu.MMU)

	gb.cpu.Init(gb.mmu)
}

// GameBoy CPU/Emulation Cycle (Fetch/Decode/Execute)
func (gb *GameBoy) Cycle() error {
	// Fetch, Decode, and Execute
	return gb.cpu.Step()

	// TODO: other stuff will happen here, of course...
}
