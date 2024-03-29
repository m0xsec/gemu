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
	"time"
)

// The DMG-01 had a Sharp LR35902 CPU (speculated to be a SM83 core), which is a hybrid of the Z80 and the 8080
// https://gbdev.io/gb-opcodes/optables/errata
type CPU struct {
	// Registers
	reg Registers

	// Memory
	mem *mmu.MMU

	// Clock Cycles
	// Interesting discussion - https://www.reddit.com/r/EmuDev/comments/4o2t6k/how_do_you_emulate_specific_cpu_speeds/
	cycles    uint32
	maxCycles uint32

	// Halt flag
	halted bool
}

// Initializes the CPU
func (cpu *CPU) Init(mmu *mmu.MMU) {
	cpu.mem = mmu
	cpu.mem.Init()

	/*
		Set initial registers to 0x00 - The DMG-01 power up sequence, per PanDocs, is:
		https://gbdev.io/pandocs/Power_Up_Sequence.html
		A = 0x01
		F = 0xB0
		B = 0x00
		C = 0x13
		D = 0x00
		E = 0xD8
		H = 0x01
		L = 0x4D
		PC = 0x0100
		SP = 0xFFFE

		This should be what the boot ROM does.
	*/
	cpu.reg.A = 0x00
	cpu.reg.F = 0x00
	cpu.reg.B = 0x00
	cpu.reg.C = 0x00
	cpu.reg.D = 0x00
	cpu.reg.E = 0x00
	cpu.reg.H = 0x00
	cpu.reg.L = 0x00
	cpu.reg.PC = 0x0000
	cpu.reg.SP = 0x0000

	// 4.194304 MHz was the highest freq the DMG could run at.
	cpu.maxCycles = 4194304

	cpu.halted = false

	// Load the boot ROM into memory
	fmt.Println("Loading boot ROM...")
	cpu.LoadBootROM()

	// Loads cartridge into memory
	// TODO: Implement cartridge loading
}

// Loads the boot ROM into memory
func (cpu *CPU) LoadBootROM() {
	for addr, val := range boot.BootRom {
		//fmt.Printf("Loading boot ROM at 0x%X with value 0x%X\n", addr, val)
		cpu.mem.Write(uint16(addr), val)
	}
	cpu.reg.PC = 0x0000
}

// Step the CPU for a single instruction - Fetch, decode, execute
func (cpu *CPU) Step() error {
	// Is the CPU halted?
	if !cpu.halted {
		op := cpu.fetch()
		instruction, valid := opcodes[op]
		if !valid {
			cpu.reg.PC++
			return fmt.Errorf("opcode not implmented: 0x%x", op)
		}

		// Execute opcode
		cpu.cycles += instruction.cycles
		instruction.execute(cpu)

		// Bits 0-3 of the Flag register are always zero, as they are unused.
		cpu.reg.F &^= FlagUnused
	} else {
		// NOP NOP bby ~
		cpu.cycles += 4
	}

	// Check CPU Cycles
	// TODO: Will need to check that this timeing is accurate...
	//fmt.Printf("Cycles: %d, Max: %d\n", cpu.cycles, cpu.maxCycles)
	if cpu.cycles > cpu.maxCycles {
		fmt.Println("[CPU Timer]: Enforcing 4.194304 Mhz")
		time.Sleep(1 * time.Second)
		cpu.cycles = 0
	}

	return nil
}

// Fetches the next opcode from memory
func (cpu *CPU) fetch() uint8 {
	op := cpu.mem.Read(cpu.reg.PC)
	//fmt.Printf("PC = 0x%X, op = 0x%X\n", cpu.reg.PC, op)
	return op
}
