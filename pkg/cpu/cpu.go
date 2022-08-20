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
package cpu

import (
	"fmt"
	"gemu/pkg/boot"
	"gemu/pkg/mmu"
)

// Sharp SM83 CPU - https://gbdev.io/gb-opcodes/optables/errata
type CPU struct {
	// Registers
	Reg Registers

	// Memory
	Mem *mmu.MMU

	// Clock Cycles
	Cycles    uint32
	MaxCycles uint32

	// Halt flag
	Halted bool
}

// Registers
type Registers struct {
	A, F, B, C, D, E, H, L uint8
	SP, PC                 uint16
}

// Flags
const (
	FlagZ      = uint8(1 << 7)
	FlagN      = uint8(1 << 6)
	FlagH      = uint8(1 << 5)
	FlagC      = uint8(1 << 4)
	FlagMask   = uint8(FlagZ | FlagN | FlagH | FlagC)
	FlagUnused = uint8(0xF)
)

// Initializes the CPU
func (cpu *CPU) Init(mmu *mmu.MMU) {
	cpu.Mem = mmu
	cpu.Mem.Init()

	// Set initial registers to 0x00
	// This will allow our boot ROM will set these to what is expected after booting
	cpu.Reg.PC = 0x00
	cpu.Reg.SP = 0x00
	cpu.Reg.A = 0x00
	cpu.Reg.F = 0x00
	cpu.Reg.B = 0x00
	cpu.Reg.C = 0x00
	cpu.Reg.D = 0x00
	cpu.Reg.E = 0x00
	cpu.Reg.H = 0x00
	cpu.Reg.L = 0x00

	// 4.194304 MHz was the highest freq the DMG could run at.
	cpu.MaxCycles = 4194304

	// Load the boot ROM into memory
	fmt.Println("Loading boot ROM...")
	cpu.LoadBootROM()
}

// Loads the boot ROM into memory
func (cpu *CPU) LoadBootROM() {
	for addr, val := range boot.BootRom {
		//fmt.Printf("Loading boot ROM at 0x%X with value 0x%X\n", addr, val)
		cpu.Mem.Write(uint16(addr), val)
	}
}
