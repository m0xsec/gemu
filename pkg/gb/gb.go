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

	// Fetch opcode
	op := gb.mmu.Read(gb.cpu.Reg.PC)
	gb.cpu.Reg.PC++
	fmt.Printf("Opcode: 0x%x\n", op)

	// Decode opcode
	instruction, valid := cpu.Opcodes[op]
	if !valid {
		return fmt.Errorf("opcode not implmented: 0x%x", op)
	}

	// Execute opcode
	gb.cpu.Cycles += instruction.Cycles
	instruction.Execute(gb.cpu)

	return nil
}
