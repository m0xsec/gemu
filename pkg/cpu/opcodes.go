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

import "fmt"

// https://gbdev.io/gb-opcodes/optables/
// https://gbdev.io/gb-opcodes/Opcodes.json
// https://www.pastraiser.com/cpu/gameboy/gameboy_opcodes.html
// http://marc.rawer.de/Gameboy/Docs/GBCPUman.pdf

var opcodes = map[uint8]struct {
	name    string
	cycles  uint32
	execute func(cpu *CPU)
}{

	/////////////////////////////////////////////////////////////////////////////////////////
	// 8-bit load/store/move instructions
	// Opcode - Mnemonic - Description
	// Cycles: n
	// Bytes: n
	// Flags: Z N H C
	/////////////////////////////////////////////////////////////////////////////////////////

	// 0x02 - LD (BC), A - Load A into memory at address BC
	// Cycles: 8
	// Bytes: 1
	// Flags: - - - -
	0x02: {name: "LD (BC), A", cycles: 8, execute: func(cpu *CPU) {
		cpu.mem.Write(cpu.BC(), cpu.reg.A)
		cpu.reg.PC++
	}},

	// 0x06 - LD B, d8 - Load immediate 8-bit value into register B
	// Cycles: 8
	// Bytes: 2
	// Flags: - - - -
	0x06: {name: "LD B, d8", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.B = cpu.mem.Read(cpu.reg.PC + 1)
		cpu.reg.PC += 2
	}},

	// 0x0A - LD A, (BC) - Load memory at address BC into register A
	// Cycles: 8
	// Bytes: 1
	// Flags: - - - -
	0x0A: {name: "LD A, (BC)", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.A = cpu.mem.Read(cpu.BC())
		cpu.reg.PC++
	}},

	// 0x0E - LD C, d8 - Load immediate 8-bit value into register C
	// Cycles: 8
	// Bytes: 2
	// Flags: - - - -
	0x0E: {name: "LD C, d8", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.C = cpu.mem.Read(cpu.reg.PC + 1)
		cpu.reg.PC += 2
	}},

	// 0x12 - LD (DE), A - Load A into memory at address DE
	// Cycles: 8
	// Bytes: 1
	// Flags: - - - -
	0x12: {name: "LD (DE), A", cycles: 8, execute: func(cpu *CPU) {
		cpu.mem.Write(cpu.DE(), cpu.reg.A)
		cpu.reg.PC++
	}},

	// 0x16 - LD D, d8 - Load immediate 8-bit value into register D
	// Cycles: 8
	// Bytes: 2
	// Flags: - - - -
	0x16: {name: "LD D, d8", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.D = cpu.mem.Read(cpu.reg.PC + 1)
		cpu.reg.PC += 2
	}},

	// 0x1A - LD A, (DE) - Load memory at address DE into register A
	// Cycles: 8
	// Bytes: 1
	// Flags: - - - -
	0x1A: {name: "LD A, (DE)", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.A = cpu.mem.Read(cpu.DE())
		cpu.reg.PC++
	}},

	// 0x1E - LD E, d8 - Load immediate 8-bit value into register E
	// Cycles: 8
	// Bytes: 2
	// Flags: - - - -
	0x1E: {name: "LD E, d8", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.E = cpu.mem.Read(cpu.reg.PC + 1)
		cpu.reg.PC += 2
	}},

	// 0x22 - LD (HL+), A - Load A into memory at address HL
	// Cycles: 8
	// Bytes: 1
	// Flags: - - - -
	0x22: {name: "LD (HL+), A", cycles: 8, execute: func(cpu *CPU) {
		cpu.mem.Write(cpu.HL(), cpu.reg.A)
		cpu.reg.PC++
	}},

	// 0x26 - LD H, d8 - Load immediate 8-bit value into register H
	// Cycles: 8
	// Bytes: 2
	// Flags: - - - -
	0x26: {name: "LD H, d8", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.H = cpu.mem.Read(cpu.reg.PC + 1)
		cpu.reg.PC += 2
	}},

	// 0x2A - LD A, (HL+) - Load memory at address HL into register A
	// Cycles: 8
	// Bytes: 1
	// Flags: - - - -
	0x2A: {name: "LD A, (HL+)", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.A = cpu.mem.Read(cpu.HL())
		cpu.reg.PC++
	}},

	// 0x2E - LD L, d8 - Load immediate 8-bit value into register L
	// Cycles: 8
	// Bytes: 2
	// Flags: - - - -
	0x2E: {name: "LD L, d8", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.L = cpu.mem.Read(cpu.reg.PC + 1)
		cpu.reg.PC += 2
	}},

	// 0x32 - LD (HL-), A - Load A into memory at address HL
	// Cycles: 8
	// Bytes: 1
	// Flags: - - - -
	0x32: {name: "LD (HL-), A", cycles: 8, execute: func(cpu *CPU) {
		cpu.mem.Write(cpu.HL(), cpu.reg.A)
		cpu.reg.PC++
	}},

	// 0x36 - LD (HL),d8 - Load immediate 8-bit value into memory at address HL
	// Cycles: 8
	// Bytes: 2
	// Flags: - - - -
	0x36: {name: "LD (HL), d8", cycles: 8, execute: func(cpu *CPU) {
		cpu.mem.Write(cpu.HL(), cpu.mem.Read(cpu.reg.PC+1))
		cpu.reg.PC += 2
	}},

	// 0x3A - LD A, (HL-) - Load memory at address HL into register A
	// Cycles: 8
	// Bytes: 1
	// Flags: - - - -
	0x3A: {name: "LD A, (HL-)", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.A = cpu.mem.Read(cpu.HL())
		cpu.reg.PC++
	}},

	// 0x3E - LD A, d8 - Load immediate 8-bit value into register A
	// Cycles: 8
	// Bytes: 2
	// Flags: - - - -
	0x3E: {name: "LD A, d8", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.A = cpu.mem.Read(cpu.reg.PC + 1)
		cpu.reg.PC += 2
	}},

	// 0x40 - LD B, B - Load register B into register B
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x40: {name: "LD B, B", cycles: 4, execute: func(cpu *CPU) {
		/*cpu.reg.B = cpu.reg.B*/
		cpu.reg.PC++
	}},

	// 0x41 - LD B, C - Load register C into register B
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x41: {name: "LD B, C", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.B = cpu.reg.C
		cpu.reg.PC++
	}},

	// 0x42 - LD B, D - Load register D into register B
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x42: {name: "LD B, D", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.B = cpu.reg.D
		cpu.reg.PC++
	}},

	// 0x43 - LD B, E - Load register E into register B
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x43: {name: "LD B, E", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.B = cpu.reg.E
		cpu.reg.PC++
	}},

	// 0x44 - LD B, H - Load register H into register B
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x44: {name: "LD B, H", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.B = cpu.reg.H
		cpu.reg.PC++
	}},

	// 0x45 - LD B, L - Load register L into register B
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x45: {name: "LD B, L", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.B = cpu.reg.L
		cpu.reg.PC++
	}},

	// 0x46 - LD B, (HL) - Load memory at address HL into register B
	// Cycles: 8
	// Bytes: 1
	// Flags: - - - -
	0x46: {name: "LD B, (HL)", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.B = cpu.mem.Read(cpu.HL())
		cpu.reg.PC++
	}},

	// 0x47 - LD B, A - Load register A into register B
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x47: {name: "LD B, A", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.B = cpu.reg.A
		cpu.reg.PC++
	}},

	// 0x48 - LD C, B - Load register B into register C
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x48: {name: "LD C, B", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.C = cpu.reg.B
		cpu.reg.PC++
	}},

	// 0x49 - LD C, C - Load register C into register C
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x49: {name: "LD C, C", cycles: 4, execute: func(cpu *CPU) {
		/*cpu.reg.C = cpu.reg.C*/
		cpu.reg.PC++
	}},

	// 0x4A - LD C, D - Load register D into register C
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x4A: {name: "LD C, D", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.C = cpu.reg.D
		cpu.reg.PC++
	}},

	// 0x4B - LD C, E - Load register E into register C
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x4B: {name: "LD C, E", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.C = cpu.reg.E
		cpu.reg.PC++
	}},

	// 0x4C - LD C, H - Load register H into register C
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x4C: {name: "LD C, H", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.C = cpu.reg.H
		cpu.reg.PC++
	}},

	// 0x4D - LD C, L - Load register L into register C
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x4D: {name: "LD C, L", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.C = cpu.reg.L
		cpu.reg.PC++
	}},

	// 0x4E - LD C, (HL) - Load memory at address HL into register C
	// Cycles: 8
	// Bytes: 1
	// Flags: - - - -
	0x4E: {name: "LD C, (HL)", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.C = cpu.mem.Read(cpu.HL())
		cpu.reg.PC++
	}},

	// 0x4F - LD C, A - Load register A into register C
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x4F: {name: "LD C, A", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.C = cpu.reg.A
		cpu.reg.PC++
	}},

	// 0x50 - LD D, B - Load register B into register D
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x50: {name: "LD D, B", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.D = cpu.reg.B
		cpu.reg.PC++
	}},

	// 0x51 - LD D, C - Load register C into register D
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x51: {name: "LD D, C", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.D = cpu.reg.C
		cpu.reg.PC++
	}},

	// 0x52 - LD D, D - Load register D into register D
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x52: {name: "LD D, D", cycles: 4, execute: func(cpu *CPU) {
		/*cpu.reg.D = cpu.reg.D*/
		cpu.reg.PC++
	}},

	// 0x53 - LD D, E - Load register E into register D
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x53: {name: "LD D, E", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.D = cpu.reg.E
		cpu.reg.PC++
	}},

	// 0x54 - LD D, H - Load register H into register D
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x54: {name: "LD D, H", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.D = cpu.reg.H
		cpu.reg.PC++
	}},

	// 0x55 - LD D, L - Load register L into register D
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x55: {name: "LD D, L", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.D = cpu.reg.L
		cpu.reg.PC++
	}},

	// 0x56 - LD D, (HL) - Load memory at address HL into register D
	// Cycles: 8
	// Bytes: 1
	// Flags: - - - -
	0x56: {name: "LD D, (HL)", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.D = cpu.mem.Read(cpu.HL())
		cpu.reg.PC++
	}},

	// 0x57 - LD D, A - Load register A into register D
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x57: {name: "LD D, A", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.D = cpu.reg.A
		cpu.reg.PC++
	}},

	// 0x58 - LD E, B - Load register B into register E
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x58: {name: "LD E, B", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.E = cpu.reg.B
		cpu.reg.PC++
	}},

	// 0x59 - LD E, C - Load register C into register E
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x59: {name: "LD E, C", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.E = cpu.reg.C
		cpu.reg.PC++
	}},

	// 0x5A - LD E, D - Load register D into register E
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x5A: {name: "LD E, D", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.E = cpu.reg.D
		cpu.reg.PC++
	}},

	// 0x5B - LD E, E - Load register E into register E
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x5B: {name: "LD E, E", cycles: 4, execute: func(cpu *CPU) {
		/*cpu.reg.E = cpu.reg.E*/
		cpu.reg.PC++
	}},

	// 0x5C - LD E, H - Load register H into register E
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x5C: {name: "LD E, H", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.E = cpu.reg.H
		cpu.reg.PC++
	}},

	// 0x5D - LD E, L - Load register L into register E
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x5D: {name: "LD E, L", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.E = cpu.reg.L
		cpu.reg.PC++
	}},

	// 0x5E - LD E, (HL) - Load memory at address HL into register E
	// Cycles: 8
	// Bytes: 1
	// Flags: - - - -
	0x5E: {name: "LD E, (HL)", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.E = cpu.mem.Read(cpu.HL())
		cpu.reg.PC++
	}},

	// 0x5F - LD E, A - Load register A into register E
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x5F: {name: "LD E, A", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.E = cpu.reg.A
		cpu.reg.PC++
	}},

	// 0x60 - LD H, B - Load register B into register H
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x60: {name: "LD H, B", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.H = cpu.reg.B
		cpu.reg.PC++
	}},

	// 0x61 - LD H, C - Load register C into register H
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x61: {name: "LD H, C", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.H = cpu.reg.C
		cpu.reg.PC++
	}},

	// 0x62 - LD H, D - Load register D into register H
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x62: {name: "LD H, D", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.H = cpu.reg.D
		cpu.reg.PC++
	}},

	// 0x63 - LD H, E - Load register E into register H
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x63: {name: "LD H, E", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.H = cpu.reg.E
		cpu.reg.PC++
	}},

	// 0x64 - LD H, H - Load register H into register H
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x64: {name: "LD H, H", cycles: 4, execute: func(cpu *CPU) {
		/*cpu.reg.H = cpu.reg.H*/
		cpu.reg.PC++
	}},

	// 0x65 - LD H, L - Load register L into register H
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x65: {name: "LD H, L", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.H = cpu.reg.L
		cpu.reg.PC++
	}},

	// 0x66 - LD H, (HL) - Load memory at address HL into register H
	// Cycles: 8
	// Bytes: 1
	// Flags: - - - -
	0x66: {name: "LD H, (HL)", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.H = cpu.mem.Read(cpu.HL())
		cpu.reg.PC++
	}},

	// 0x67 - LD H, A - Load register A into register H
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x67: {name: "LD H, A", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.H = cpu.reg.A
		cpu.reg.PC++
	}},

	// 0x68 - LD L, B - Load register B into register L
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x68: {name: "LD L, B", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.L = cpu.reg.B
		cpu.reg.PC++
	}},

	// 0x69 - LD L, C - Load register C into register L
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x69: {name: "LD L, C", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.L = cpu.reg.C
		cpu.reg.PC++
	}},

	// 0x6A - LD L, D - Load register D into register L
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x6A: {name: "LD L, D", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.L = cpu.reg.D
		cpu.reg.PC++
	}},

	// 0x6B - LD L, E - Load register E into register L
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x6B: {name: "LD L, E", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.L = cpu.reg.E
		cpu.reg.PC++
	}},

	// 0x6C - LD L, H - Load register H into register L
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x6C: {name: "LD L, H", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.L = cpu.reg.H
		cpu.reg.PC++
	}},

	// 0x6D - LD L, L - Load register L into register L
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x6D: {name: "LD L, L", cycles: 4, execute: func(cpu *CPU) {
		/*cpu.reg.L = cpu.reg.L*/
		cpu.reg.PC++
	}},

	// 0x6E - LD L, (HL) - Load memory at address HL into register L
	// Cycles: 8
	// Bytes: 1
	// Flags: - - - -
	0x6E: {name: "LD L, (HL)", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.L = cpu.mem.Read(cpu.HL())
		cpu.reg.PC++
	}},

	// 0x6F - LD L, A - Load register A into register L
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x6F: {name: "LD L, A", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.L = cpu.reg.A
		cpu.reg.PC++
	}},

	// 0x70 - LD (HL), B - Load register B into memory at address HL
	// Cycles: 8
	// Bytes: 1
	// Flags: - - - -
	0x70: {name: "LD (HL), B", cycles: 8, execute: func(cpu *CPU) {
		cpu.mem.Write(cpu.HL(), cpu.reg.B)
		cpu.reg.PC++
	}},

	// 0x71 - LD (HL), C - Load register C into memory at address HL
	// Cycles: 8
	// Bytes: 1
	// Flags: - - - -
	0x71: {name: "LD (HL), C", cycles: 8, execute: func(cpu *CPU) {
		cpu.mem.Write(cpu.HL(), cpu.reg.C)
		cpu.reg.PC++
	}},

	// 0x72 - LD (HL), D - Load register D into memory at address HL
	// Cycles: 8
	// Bytes: 1
	// Flags: - - - -
	0x72: {name: "LD (HL), D", cycles: 8, execute: func(cpu *CPU) {
		cpu.mem.Write(cpu.HL(), cpu.reg.D)
		cpu.reg.PC++
	}},

	// 0x73 - LD (HL), E - Load register E into memory at address HL
	// Cycles: 8
	// Bytes: 1
	// Flags: - - - -
	0x73: {name: "LD (HL), E", cycles: 8, execute: func(cpu *CPU) {
		cpu.mem.Write(cpu.HL(), cpu.reg.E)
		cpu.reg.PC++
	}},

	// 0x74 - LD (HL), H - Load register H into memory at address HL
	// Cycles: 8
	// Bytes: 1
	// Flags: - - - -
	0x74: {name: "LD (HL), H", cycles: 8, execute: func(cpu *CPU) {
		cpu.mem.Write(cpu.HL(), cpu.reg.H)
		cpu.reg.PC++
	}},

	// 0x75 - LD (HL), L - Load register L into memory at address HL
	// Cycles: 8
	// Bytes: 1
	// Flags: - - - -
	0x75: {name: "LD (HL), L", cycles: 8, execute: func(cpu *CPU) {
		cpu.mem.Write(cpu.HL(), cpu.reg.L)
		cpu.reg.PC++
	}},

	// 0x76 - HALT - Halt the CPU until an interrupt occurs
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x77: {name: "LD (HL), A", cycles: 8, execute: func(cpu *CPU) {
		cpu.mem.Write(cpu.HL(), cpu.reg.A)
		cpu.reg.PC++
	}},

	// 0x78 - LD A, B - Load register B into register A
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x78: {name: "LD A, B", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.A = cpu.reg.B
		cpu.reg.PC++
	}},

	// 0x79 - LD A, C - Load register C into register A
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x79: {name: "LD A, C", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.A = cpu.reg.C
		cpu.reg.PC++
	}},

	// 0x7A - LD A, D - Load register D into register A
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x7A: {name: "LD A, D", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.A = cpu.reg.D
		cpu.reg.PC++
	}},

	// 0x7B - LD A, E - Load register E into register A
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x7B: {name: "LD A, E", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.A = cpu.reg.E
		cpu.reg.PC++
	}},

	// 0x7C - LD A, H - Load register H into register A
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x7C: {name: "LD A, H", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.A = cpu.reg.H
		cpu.reg.PC++
	}},

	// 0x7D - LD A, L - Load register L into register A
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x7D: {name: "LD A, L", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.A = cpu.reg.L
		cpu.reg.PC++
	}},

	// 0x7E - LD A, (HL) - Load memory at address HL into register A
	// Cycles: 8
	// Bytes: 1
	// Flags: - - - -
	0x7E: {name: "LD A, (HL)", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.A = cpu.mem.Read(cpu.HL())
		cpu.reg.PC++
	}},

	// 0x7F - LD A, A - Load register A into register A
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x7F: {name: "LD A, A", cycles: 4, execute: func(cpu *CPU) {
		/*cpu.reg.A = cpu.reg.A*/
		cpu.reg.PC++
	}},

	// 0xE0 - LDH (n), A - Load register A into memory at address 0xFF00 + n
	// Cycles: 12
	// Bytes: 2
	// Flags: - - - -
	0xE0: {name: "LDH (n), A", cycles: 12, execute: func(cpu *CPU) {
		cpu.mem.Write(0xFF00+uint16(cpu.mem.Read(cpu.reg.PC+1)), cpu.reg.A)
		cpu.reg.PC += 2
	}},

	// 0xE2 - LD (C), A - Load register A into memory at address 0xFF00 + C
	// Cycles: 8
	// Bytes: 1
	// Flags: - - - -
	0xE2: {name: "LD (C), A", cycles: 8, execute: func(cpu *CPU) {
		cpu.mem.Write(0xFF00+uint16(cpu.reg.C), cpu.reg.A)
		cpu.reg.PC += 2
	}},

	// 0xEA - LD (a16), A - Load register A into memory at the absolute 16-bit address a16
	// Cycles: 16
	// Bytes: 3
	// Flags: - - - -
	0xEA: {name: "LD (a16), A", cycles: 16, execute: func(cpu *CPU) {
		addr := uint16(cpu.mem.Read(cpu.reg.PC+1)) | uint16(cpu.mem.Read(cpu.reg.PC+2))<<8
		cpu.mem.Write(addr, cpu.reg.A)
		cpu.reg.PC += 3
	}},

	// 0xF0 - LDH A, (n) - Load memory at address 0xFF00 + n into register A
	// Cycles: 12
	// Bytes: 2
	// Flags: - - - -
	0xF0: {name: "LDH A, (a8)", cycles: 12, execute: func(cpu *CPU) {
		addr := uint16(cpu.mem.Read(cpu.reg.PC + 1))
		cpu.reg.A = cpu.mem.Read(0xFF00 + addr)
		cpu.reg.PC += 2
	}},

	// 0xF2 - LD A, (C) - Load memory at address 0xFF00 + C into register A
	// Cycles: 8
	// Bytes: 1
	// Flags: - - - -
	0xF2: {name: "LD A, (C)", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.A = cpu.mem.Read(0xFF00 + uint16(cpu.reg.C))
		cpu.reg.PC += 2
	}},

	// 0xFA - LD A, (a16) - Load memory at the absolute 16-bit address a16 into register A
	// Cycles: 16
	// Bytes: 3
	// Flags: - - - -
	0xFA: {name: "LD A, (a16)", cycles: 16, execute: func(cpu *CPU) {
		addr := uint16(cpu.mem.Read(cpu.reg.PC+1)) | uint16(cpu.mem.Read(cpu.reg.PC+2))<<8
		cpu.reg.A = cpu.mem.Read(addr)
		cpu.reg.PC += 3
	}},

	/////////////////////////////////////////////////////////////////////////////////////////
	// 16bit load/store/move instructions
	// Opcode - Mnemonic - Description
	// Cycles: n
	// Bytes: n
	// Flags: Z N H C
	/////////////////////////////////////////////////////////////////////////////////////////

	// 0x01 - LD BC, d16 - Load 16-bit immediate value d16 into register BC
	// Cycles: 12
	// Bytes: 3
	// Flags: - - - -
	0x01: {name: "LD BC, d16", cycles: 12, execute: func(cpu *CPU) {
		cpu.SetBC(uint16(cpu.mem.Read(cpu.reg.PC+1)) | uint16(cpu.mem.Read(cpu.reg.PC+2))<<8)
		cpu.reg.PC += 3
	}},

	// 0x11 - LD DE, d16 - Load 16-bit immediate value d16 into register DE
	// Cycles: 12
	// Bytes: 3
	// Flags: - - - -
	0x11: {name: "LD DE, d16", cycles: 12, execute: func(cpu *CPU) {
		cpu.SetDE(uint16(cpu.mem.Read(cpu.reg.PC+1)) | uint16(cpu.mem.Read(cpu.reg.PC+2))<<8)
		cpu.reg.PC += 3
	}},

	// 0x21 - LD HL, d16 - Load 16-bit immediate value d16 into register HL
	// Cycles: 12
	// Bytes: 3
	// Flags: - - - -
	0x21: {name: "LD HL, d16", cycles: 12, execute: func(cpu *CPU) {
		cpu.SetHL(uint16(cpu.mem.Read(cpu.reg.PC+1)) | uint16(cpu.mem.Read(cpu.reg.PC+2))<<8)
		cpu.reg.PC += 3
	}},

	// 0x31 - LD SP, d16 - Load 16-bit immediate value d16 into register SP
	// Cycles: 12
	// Bytes: 3
	// Flags: - - - -
	0x31: {name: "LD SP, d16", cycles: 12, execute: func(cpu *CPU) {
		cpu.reg.SP = uint16(cpu.mem.Read(cpu.reg.PC+1)) | uint16(cpu.mem.Read(cpu.reg.PC+2))<<8
		cpu.reg.PC += 3
	}},

	// 0xC1 - POP BC - Pop register BC from stack
	// Cycles: 12
	// Bytes: 1
	// Flags: - - - -
	0xC1: {name: "POP BC", cycles: 12, execute: func(cpu *CPU) {
		cpu.reg.C = cpu.stackPop()
		cpu.reg.B = cpu.stackPop()
		cpu.reg.PC++
	}},

	// 0xD1 - POP DE - Pop register DE from stack
	// Cycles: 12
	// Bytes: 1
	// Flags: - - - -
	0xD1: {name: "POP DE", cycles: 12, execute: func(cpu *CPU) {
		cpu.reg.E = cpu.stackPop()
		cpu.reg.D = cpu.stackPop()
		cpu.reg.PC++
	}},

	// 0xE1 - POP HL - Pop register HL from stack
	// Cycles: 12
	// Bytes: 1
	// Flags: - - - -
	0xE1: {name: "POP HL", cycles: 12, execute: func(cpu *CPU) {
		cpu.reg.L = cpu.stackPop()
		cpu.reg.H = cpu.stackPop()
		cpu.reg.PC++
	}},

	// 0xF1 - POP AF - Pop register AF from stack
	// Cycles: 12
	// Bytes: 1
	// Flags: Z N H C
	0xF1: {name: "POP AF", cycles: 12, execute: func(cpu *CPU) {
		// Since the F register (Flag register) is popped, all flags are changed.
		cpu.reg.F = cpu.stackPop()
		cpu.reg.A = cpu.stackPop()
		cpu.reg.PC++
	}},

	// 0xC5 - PUSH BC - Push register BC onto stack
	// Cycles: 16
	// Bytes: 1
	// Flags: - - - -
	0xC5: {name: "PUSH BC", cycles: 16, execute: func(cpu *CPU) {
		cpu.stackPush(cpu.reg.B)
		cpu.stackPush(cpu.reg.C)
		cpu.reg.PC++
	}},

	// 0xD5 - PUSH DE - Push register DE onto stack
	// Cycles: 16
	// Bytes: 1
	// Flags: - - - -
	0xD5: {name: "PUSH DE", cycles: 16, execute: func(cpu *CPU) {
		cpu.stackPush(cpu.reg.D)
		cpu.stackPush(cpu.reg.E)
		cpu.reg.PC++
	}},

	// 0xE5 - PUSH HL - Push register HL onto stack
	// Cycles: 16
	// Bytes: 1
	// Flags: - - - -
	0xE5: {name: "PUSH HL", cycles: 16, execute: func(cpu *CPU) {
		cpu.stackPush(cpu.reg.H)
		cpu.stackPush(cpu.reg.L)
		cpu.reg.PC++
	}},

	// 0xF5 - PUSH AF - Push register AF onto stack
	// Cycles: 16
	// Bytes: 1
	// Flags: - - - -
	0xF5: {name: "PUSH AF", cycles: 16, execute: func(cpu *CPU) {
		cpu.stackPush(cpu.reg.A)
		cpu.stackPush(cpu.reg.F)
		cpu.reg.PC++
	}},

	// 0x08 - LD (a16), SP - Load SP into memory address a16
	// Cycles: 20
	// Bytes: 3
	// Flags: - - - -
	0x08: {name: "LD (a16), SP", cycles: 20, execute: func(cpu *CPU) {
		addr := uint16(cpu.mem.Read(cpu.reg.PC+1)) + uint16(cpu.mem.Read(cpu.reg.PC+2))<<8
		cpu.mem.Write(addr, uint8(cpu.reg.SP&0xFF))
		cpu.mem.Write(addr+1, uint8(cpu.reg.SP>>8)&0xFF)
		cpu.reg.PC += 3
	}},

	// 0xF8 - LD HL, SP+r8 - Load SP+r8 into register HL
	// Cycles: 12
	// Bytes: 2
	// Flags: 0 0 H C
	0xF8: {name: "LD HL, SP+r8", cycles: 12, execute: func(cpu *CPU) {
		// Bit clear flags
		cpu.reg.F &= ^FlagMask

		// Get the value of r8 and SP+r8
		r8 := uint8(cpu.mem.Read(cpu.reg.PC + 1))
		spr8 := cpu.reg.SP + uint16(r8)

		// Set flags
		if ((cpu.reg.SP & 0xF) + (uint16(r8) & 0xF)) > 0xF {
			cpu.reg.F |= FlagH
		}
		if ((cpu.reg.SP & 0xFF) + (uint16(r8) & 0xFF)) > 0xFF {
			cpu.reg.F |= FlagC
		}

		// Set HL
		cpu.SetHL(spr8)

		cpu.reg.PC += 2
	}},

	// 0xF9 - LD SP, HL - Load register HL into register SP
	// Cycles: 8
	// Bytes: 1
	// Flags: - - - -
	0xF9: {name: "LD SP, HL", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.SP = cpu.HL()
		cpu.reg.PC++
	}},

	// TODO: 8bit arithmetic/logic instructions
	/////////////////////////////////////////////////////////////////////////////////////////
	// 8bit arithmetic/logic instructions
	// Opcode - Mnemonic - Description
	// Cycles: n
	// Bytes: n
	// Flags: Z N H C
	/////////////////////////////////////////////////////////////////////////////////////////

	// 0x04 - INC B - Increment register B
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 0 H -
	0x04: {name: "INC B", cycles: 4, execute: func(cpu *CPU) {
		cpu.Inc8(&cpu.reg.B)
		cpu.reg.PC++
	}},

	// 0x0C - INC C - Increment register C
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 0 H -
	0x0C: {name: "INC C", cycles: 4, execute: func(cpu *CPU) {
		cpu.Inc8(&cpu.reg.C)
		cpu.reg.PC++
	}},

	// 0x14 - INC D - Increment register D
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 0 H -
	0x14: {name: "INC D", cycles: 4, execute: func(cpu *CPU) {
		cpu.Inc8(&cpu.reg.D)
		cpu.reg.PC++
	}},

	// 0x1C - INC E - Increment register E
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 0 H -
	0x1C: {name: "INC E", cycles: 4, execute: func(cpu *CPU) {
		cpu.Inc8(&cpu.reg.E)
		cpu.reg.PC++
	}},

	// 0x24 - INC H - Increment register H
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 0 H -
	0x24: {name: "INC H", cycles: 4, execute: func(cpu *CPU) {
		cpu.Inc8(&cpu.reg.H)
		cpu.reg.PC++
	}},

	// 0x2C - INC L - Increment register L
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 0 H -
	0x2C: {name: "INC L", cycles: 4, execute: func(cpu *CPU) {
		cpu.Inc8(&cpu.reg.L)
		cpu.reg.PC++
	}},

	// 0x34 - INC (HL) - Increment memory address HL
	// Cycles: 12
	// Bytes: 1
	// Flags: Z 0 H -
	0x34: {name: "INC (HL)", cycles: 12, execute: func(cpu *CPU) {
		addr := cpu.HL()
		val := cpu.mem.Read(addr)
		cpu.Inc8(&val)
		cpu.mem.Write(addr, val)
		cpu.reg.PC++
	}},

	// 0x3C - INC A - Increment register A
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 0 H -
	0x3C: {name: "INC A", cycles: 4, execute: func(cpu *CPU) {
		cpu.Inc8(&cpu.reg.A)
		cpu.reg.PC++
	}},

	// 0x05 - DEC B - Decrement register B
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 1 H -
	0x05: {name: "DEC B", cycles: 4, execute: func(cpu *CPU) {
		cpu.Dec8(&cpu.reg.B)
		cpu.reg.PC++
	}},

	// 0x0D - DEC C - Decrement register C
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 1 H -
	0x0D: {name: "DEC C", cycles: 4, execute: func(cpu *CPU) {
		cpu.Dec8(&cpu.reg.C)
		cpu.reg.PC++
	}},

	// 0x15 - DEC D - Decrement register D
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 1 H -
	0x15: {name: "DEC D", cycles: 4, execute: func(cpu *CPU) {
		cpu.Dec8(&cpu.reg.D)
		cpu.reg.PC++
	}},

	// 0x1D - DEC E - Decrement register E
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 1 H -
	0x1D: {name: "DEC E", cycles: 4, execute: func(cpu *CPU) {
		cpu.Dec8(&cpu.reg.E)
		cpu.reg.PC++
	}},

	// 0x25 - DEC H - Decrement register H
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 1 H -
	0x25: {name: "DEC H", cycles: 4, execute: func(cpu *CPU) {
		cpu.Dec8(&cpu.reg.H)
		cpu.reg.PC++
	}},

	// 0x2D - DEC L - Decrement register L
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 1 H -
	0x2D: {name: "DEC L", cycles: 4, execute: func(cpu *CPU) {
		cpu.Dec8(&cpu.reg.L)
		cpu.reg.PC++
	}},

	// 0x35 - DEC (HL) - Decrement memory address HL
	// Cycles: 12
	// Bytes: 1
	// Flags: Z 1 H -
	0x35: {name: "DEC (HL)", cycles: 12, execute: func(cpu *CPU) {
		addr := cpu.HL()
		val := cpu.mem.Read(addr)
		cpu.Dec8(&val)
		cpu.mem.Write(addr, val)
		cpu.reg.PC++
	}},

	// 0x3D - DEC A - Decrement register A
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 1 H -
	0x3D: {name: "DEC A", cycles: 4, execute: func(cpu *CPU) {
		cpu.Dec8(&cpu.reg.A)
		cpu.reg.PC++
	}},

	// 0x80 - ADD A, B - Add register B to register A
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 0 H C
	0x80: {name: "ADD A, B", cycles: 4, execute: func(cpu *CPU) {
		cpu.Add8(&cpu.reg.A, cpu.reg.B, false)
		cpu.reg.PC++
	}},

	// 0x81 - ADD A, C - Add register C to register A
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 0 H C
	0x81: {name: "ADD A, C", cycles: 4, execute: func(cpu *CPU) {
		cpu.Add8(&cpu.reg.A, cpu.reg.C, false)
		cpu.reg.PC++
	}},

	// 0x82 - ADD A, D - Add register D to register A
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 0 H C
	0x82: {name: "ADD A, D", cycles: 4, execute: func(cpu *CPU) {
		cpu.Add8(&cpu.reg.A, cpu.reg.D, false)
		cpu.reg.PC++
	}},

	// 0x83 - ADD A, E - Add register E to register A
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 0 H C
	0x83: {name: "ADD A, E", cycles: 4, execute: func(cpu *CPU) {
		cpu.Add8(&cpu.reg.A, cpu.reg.E, false)
		cpu.reg.PC++
	}},

	// 0x84 - ADD A, H - Add register H to register A
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 0 H C
	0x84: {name: "ADD A, H", cycles: 4, execute: func(cpu *CPU) {
		cpu.Add8(&cpu.reg.A, cpu.reg.H, false)
		cpu.reg.PC++
	}},

	// 0x85 - ADD A, L - Add register L to register A
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 0 H C
	0x85: {name: "ADD A, L", cycles: 4, execute: func(cpu *CPU) {
		cpu.Add8(&cpu.reg.A, cpu.reg.L, false)
		cpu.reg.PC++
	}},

	// 0x86 - ADD A, (HL) - Add value at memory address HL to register A
	// Cycles: 8
	// Bytes: 1
	// Flags: Z 0 H C
	0x86: {name: "ADD A, (HL)", cycles: 8, execute: func(cpu *CPU) {
		cpu.Add8(&cpu.reg.A, cpu.mem.Read(cpu.HL()), false)
		cpu.reg.PC++
	}},

	// 0x87 - ADD A, A - Add register A to register A
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 0 H C
	0x87: {name: "ADD A, A", cycles: 4, execute: func(cpu *CPU) {
		cpu.Add8(&cpu.reg.A, cpu.reg.A, false)
		cpu.reg.PC++
	}},

	// 0x88 - ADC A, B - Add register B to register A with carry flag
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 0 H C
	0x88: {name: "ADC A, B", cycles: 4, execute: func(cpu *CPU) {
		cpu.Add8(&cpu.reg.A, cpu.reg.B, true)
		cpu.reg.PC++
	}},

	// 0x89 - ADC A, C - Add register C to register A with carry flag
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 0 H C
	0x89: {name: "ADC A, C", cycles: 4, execute: func(cpu *CPU) {
		cpu.Add8(&cpu.reg.A, cpu.reg.C, true)
		cpu.reg.PC++
	}},

	// 0x8A - ADC A, D - Add register D to register A with carry flag
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 0 H C
	0x8A: {name: "ADC A, D", cycles: 4, execute: func(cpu *CPU) {
		cpu.Add8(&cpu.reg.A, cpu.reg.D, true)
		cpu.reg.PC++
	}},

	// 0x8B - ADC A, E - Add register E to register A with carry flag
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 0 H C
	0x8B: {name: "ADC A, E", cycles: 4, execute: func(cpu *CPU) {
		cpu.Add8(&cpu.reg.A, cpu.reg.E, true)
		cpu.reg.PC++
	}},

	// 0x8C - ADC A, H - Add register H to register A with carry flag
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 0 H C
	0x8C: {name: "ADC A, H", cycles: 4, execute: func(cpu *CPU) {
		cpu.Add8(&cpu.reg.A, cpu.reg.H, true)
		cpu.reg.PC++
	}},

	// 0x8D - ADC A, L - Add register L to register A with carry flag
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 0 H C
	0x8D: {name: "ADC A, L", cycles: 4, execute: func(cpu *CPU) {
		cpu.Add8(&cpu.reg.A, cpu.reg.L, true)
		cpu.reg.PC++
	}},

	// 0x8E - ADC A, (HL) - Add value at memory address HL to register A with carry flag
	// Cycles: 8
	// Bytes: 1
	// Flags: Z 0 H C
	0x8E: {name: "ADC A, (HL)", cycles: 8, execute: func(cpu *CPU) {
		cpu.Add8(&cpu.reg.A, cpu.mem.Read(cpu.HL()), true)
		cpu.reg.PC++
	}},

	// 0x8F - ADC A, A - Add register A to register A with carry flag
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 0 H C
	0x8F: {name: "ADC A, A", cycles: 4, execute: func(cpu *CPU) {
		cpu.Add8(&cpu.reg.A, cpu.reg.A, true)
		cpu.reg.PC++
	}},

	// 0x2F - CPL - Complement register A
	// Cycles: 4
	// Bytes: 1
	// Flags: - 1 1 -
	0x2F: {name: "CPL", cycles: 4, execute: func(cpu *CPU) {
		// Set flags N, H
		cpu.reg.F |= FlagN
		cpu.reg.F |= FlagH

		// Complement register A
		cpu.reg.A = ^cpu.reg.A
		cpu.reg.PC++
	}},

	// 0x3F - CCF - Complement carry flag
	// Cycles: 4
	// Bytes: 1
	// Flags: - 0 0 C
	0x3F: {name: "CCF", cycles: 4, execute: func(cpu *CPU) {
		// If C flag is set, reset it
		if cpu.reg.F&FlagC != 0 {
			cpu.reg.F &= ^FlagC
		} else {
			// Otherwise set it
			cpu.reg.F |= FlagC
		}

		// Complement carry flag
		cpu.reg.F ^= FlagC
		cpu.reg.PC++
	}},

	// 0x90 - SUB B - Subtract register B from register A
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 1 H C
	0x90: {name: "SUB B", cycles: 4, execute: func(cpu *CPU) {
		cpu.Sub8(&cpu.reg.A, cpu.reg.B, false)
		cpu.reg.PC++
	}},

	// 0x91 - SUB C - Subtract register C from register A
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 1 H C
	0x91: {name: "SUB C", cycles: 4, execute: func(cpu *CPU) {
		cpu.Sub8(&cpu.reg.A, cpu.reg.C, false)
		cpu.reg.PC++
	}},

	// 0x92 - SUB D - Subtract register D from register A
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 1 H C
	0x92: {name: "SUB D", cycles: 4, execute: func(cpu *CPU) {
		cpu.Sub8(&cpu.reg.A, cpu.reg.D, false)
		cpu.reg.PC++
	}},

	// 0x93 - SUB E - Subtract register E from register A
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 1 H C
	0x93: {name: "SUB E", cycles: 4, execute: func(cpu *CPU) {
		cpu.Sub8(&cpu.reg.A, cpu.reg.E, false)
		cpu.reg.PC++
	}},

	// 0x94 - SUB H - Subtract register H from register A
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 1 H C
	0x94: {name: "SUB H", cycles: 4, execute: func(cpu *CPU) {
		cpu.Sub8(&cpu.reg.A, cpu.reg.H, false)
		cpu.reg.PC++
	}},

	// 0x95 - SUB L - Subtract register L from register A
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 1 H C
	0x95: {name: "SUB L", cycles: 4, execute: func(cpu *CPU) {
		cpu.Sub8(&cpu.reg.A, cpu.reg.L, false)
		cpu.reg.PC++
	}},

	// 0x96 - SUB (HL) - Subtract value at memory address HL from register A
	// Cycles: 8
	// Bytes: 1
	// Flags: Z 1 H C
	0x96: {name: "SUB (HL)", cycles: 8, execute: func(cpu *CPU) {
		cpu.Sub8(&cpu.reg.A, cpu.mem.Read(cpu.HL()), false)
		cpu.reg.PC++
	}},

	// 0x97 - SUB A - Subtract register A from register A
	// Cycles: 4
	// Bytes: 1
	// Flags: 1 1 0 0
	0x97: {name: "SUB A", cycles: 4, execute: func(cpu *CPU) {
		cpu.Sub8(&cpu.reg.A, cpu.reg.A, false)
		cpu.reg.PC++
	}},

	// 0x98 - SBC A, B - Subtract register B from register A with carry flag
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 1 H C
	0x98: {name: "SBC A, B", cycles: 4, execute: func(cpu *CPU) {
		cpu.Sub8(&cpu.reg.A, cpu.reg.B, true)
		cpu.reg.PC++
	}},

	// 0x99 - SBC A, C - Subtract register C from register A with carry flag
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 1 H C
	0x99: {name: "SBC A, C", cycles: 4, execute: func(cpu *CPU) {
		cpu.Sub8(&cpu.reg.A, cpu.reg.C, true)
		cpu.reg.PC++
	}},

	// 0x9A - SBC A, D - Subtract register D from register A with carry flag
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 1 H C
	0x9A: {name: "SBC A, D", cycles: 4, execute: func(cpu *CPU) {
		cpu.Sub8(&cpu.reg.A, cpu.reg.D, true)
		cpu.reg.PC++
	}},

	// 0x9B - SBC A, E - Subtract register E from register A with carry flag
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 1 H C
	0x9B: {name: "SBC A, E", cycles: 4, execute: func(cpu *CPU) {
		cpu.Sub8(&cpu.reg.A, cpu.reg.E, true)
		cpu.reg.PC++
	}},

	// 0x9C - SBC A, H - Subtract register H from register A with carry flag
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 1 H C
	0x9C: {name: "SBC A, H", cycles: 4, execute: func(cpu *CPU) {
		cpu.Sub8(&cpu.reg.A, cpu.reg.H, true)
		cpu.reg.PC++
	}},

	// 0x9D - SBC A, L - Subtract register L from register A with carry flag
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 1 H C
	0x9D: {name: "SBC A, L", cycles: 4, execute: func(cpu *CPU) {
		cpu.Sub8(&cpu.reg.A, cpu.reg.L, true)
		cpu.reg.PC++
	}},

	// 0x9E - SBC A, (HL) - Subtract value at memory address HL from register A with carry flag
	// Cycles: 8
	// Bytes: 1
	// Flags: Z 1 H C
	0x9E: {name: "SBC A, (HL)", cycles: 8, execute: func(cpu *CPU) {
		cpu.Sub8(&cpu.reg.A, cpu.mem.Read(cpu.HL()), true)
		cpu.reg.PC++
	}},

	// 0x9F - SBC A, A - Subtract register A from register A with carry flag
	// Cycles: 4
	// Bytes: 1
	// Flags: Z 1 H C
	0x9F: {name: "SBC A, A", cycles: 4, execute: func(cpu *CPU) {
		cpu.Sub8(&cpu.reg.A, cpu.reg.A, true)
		cpu.reg.PC++
	}},

	// 0x27 - DAA - Decimal adjust register A
	// Cycles: 4
	// Bytes: 1
	// Flags: Z - 0 C
	0x27: {name: "DAA", cycles: 4, execute: func(cpu *CPU) {
		// Implementation pulled from AWJ's post #433 here - https://forums.nesdev.org/viewtopic.php?f=20&t=15944
		// thank you <3
		// TODO: If this fails, it is probably due to how H and N flags are set in the other instructions.
		//	     DAA is the only thing that actually uses those flags!
		/*
			// note: assumes a is a uint8_t and wraps from 0xff to 0
			if (!n_flag) {  // after an addition, adjust if (half-)carry occurred or if result is out of bounds
			if (c_flag || a > 0x99) { a += 0x60; c_flag = 1; }
			if (h_flag || (a & 0x0f) > 0x09) { a += 0x6; }
			} else {  // after a subtraction, only adjust if (half-)carry occurred
			if (c_flag) { a -= 0x60; }
			if (h_flag) { a -= 0x6; }
			}
			// these flags are always updated
			z_flag = (a == 0); // the usual z flag
			h_flag = 0; // h flag is always cleared
		*/
		if cpu.reg.A > 0x99 || cpu.reg.F&FlagC != 0 {
			cpu.reg.A += 0x60
			cpu.reg.F |= FlagC
		}
		if (cpu.reg.A&0x0F) > 0x09 || cpu.reg.F&FlagH != 0 {
			cpu.reg.A += 0x06
		}
		cpu.reg.F &= FlagC
		if cpu.reg.A == 0 {
			cpu.reg.F |= FlagZ
		}
		cpu.reg.F &= FlagH
		cpu.reg.PC++
	}},

	// TODO:
	// DAA, SCF, AND, OR, XOR, CP - double check that these are what is remaining so nothing is missed :)

	// ...

	// TODO: 16bit arithmetic/logic instructions
	/////////////////////////////////////////////////////////////////////////////////////////
	// 16bit arithmetic/logic instructions
	// Opcode - Mnemonic - Description
	// Cycles: n
	// Bytes: n
	// Flags: Z N H C
	/////////////////////////////////////////////////////////////////////////////////////////
	// ...

	// TODO: 8bit rotation/shift instructions (CB prefix operations)
	/////////////////////////////////////////////////////////////////////////////////////////
	// 8bit rotation/shift instructions
	// Opcode - Mnemonic - Description
	// Cycles: n
	// Bytes: n
	// Flags: Z N H C
	/////////////////////////////////////////////////////////////////////////////////////////

	// 0xCB - PREFIX CB - CB prefix operation
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0xCB: {name: "PREFIX CB", cycles: 4, execute: func(cpu *CPU) {
		//TODO: Need to actually implement the CB Prefix operations lol
		fmt.Println("[Warning] CB prefix operations not implemented")
		cb_op := cpu.mem.Read(cpu.reg.PC + 1)
		fmt.Printf("\t |--> cb_op: %02X\n", cb_op)
		fmt.Printf("\t |--> This may cause Z80 UNUSED warnings\n")
		cpu.reg.PC += 2
	}},

	///...

	// TODO: Jumps/calls instructions
	/////////////////////////////////////////////////////////////////////////////////////////
	// Jumps/calls instructions
	// Opcode - Mnemonic - Description
	// Cycles: n
	// Bytes: n
	// Flags: Z N H C
	/////////////////////////////////////////////////////////////////////////////////////////
	// ...

	// TODO: Misc/control instructions
	/////////////////////////////////////////////////////////////////////////////////////////
	// Misc/control instructions
	// Opcode - Mnemonic - Description
	// Cycles: n
	// Bytes: n
	// Flags: Z N H C
	/////////////////////////////////////////////////////////////////////////////////////////

	// 0x00 - NOP - No operation
	// Cycles: 4
	// Bytes: 1
	// Flags: - - - -
	0x00: {name: "NOP", cycles: 4, execute: func(cpu *CPU) { cpu.reg.PC++ }},
	// ...

	/////////////////////////////////////////////////////////////////////////////////////////
	// Unused Z80 opcodes
	// TODO: Should I make the emulator crash or just ignore?
	//
	// These opcodes were not implemented in the Gameboy CPU,
	// and will cause the Gameboy to lock up when used.
	// https://gbdev.io/pandocs/CPU_Comparison_with_Z80.html
	// https://clrhome.org/table/ (Z80 Opcode Set)
	/////////////////////////////////////////////////////////////////////////////////////////
	0xD3: {name: "OUT (a8), A", cycles: 4, execute: func(cpu *CPU) {
		fmt.Println("[Z80 UNUSED] 0xD3 {OUT (a8), A} will cause the Gameboy to crash!")
		cpu.reg.PC++
	}},
	0xDB: {name: "IN A, (a8)", cycles: 4, execute: func(cpu *CPU) {
		fmt.Println("[Z80 UNUSED] 0xDB {IN A, (a8)} will cause the Gameboy to crash!")
		cpu.reg.PC++
	}},
	0xDD: {name: "IX PREFIX", cycles: 4, execute: func(cpu *CPU) {
		fmt.Println("[Z80 UNUSED] 0xDD {IX PREFIX} will cause the Gameboy to crash!")
		cpu.reg.PC++
	}},
	0xE3: {name: "EX (SP), HL", cycles: 4, execute: func(cpu *CPU) {
		fmt.Println("[Z80 UNUSED] 0xE3 {EX (SP), HL} will cause the Gameboy to crash!")
		cpu.reg.PC++
	}},
	0xE4: {name: "CALL PO, a16", cycles: 4, execute: func(cpu *CPU) {
		fmt.Println("[Z80 UNUSED] 0xE4 {CALL PO, a16} will cause the Gameboy to crash!")
		cpu.reg.PC++
	}},
	0xEB: {name: "EX DE, HL", cycles: 4, execute: func(cpu *CPU) {
		fmt.Println("[Z80 UNUSED] 0xEB {EX DE, HL} will cause the Gameboy to crash!")
		cpu.reg.PC++
	}},
	0xEC: {name: "CALL PE, a16", cycles: 4, execute: func(cpu *CPU) {
		fmt.Println("[Z80 UNUSED] 0xEC {CALL PE, a16} will cause the Gameboy to crash!")
		cpu.reg.PC++
	}},
	0xED: {name: "MISC PREFIX (ED)", cycles: 4, execute: func(cpu *CPU) {
		fmt.Println("[Z80 UNUSED] 0xED {MISC PREFIX (ED)} will cause the Gameboy to crash!")
		cpu.reg.PC++
	}},
	0xF4: {name: "CALL P, a16", cycles: 4, execute: func(cpu *CPU) {
		fmt.Println("[Z80 UNUSED] 0xF4 {CALL P, a16} will cause the Gameboy to crash!")
		cpu.reg.PC++
	}},
	0xFC: {name: "CALL M, a16", cycles: 4, execute: func(cpu *CPU) {
		fmt.Println("[Z80 UNUSED] 0xFC {CALL M, a16} will cause the Gameboy to crash!")
		cpu.reg.PC++
	}},
	0xFD: {name: "IY PREFIX", cycles: 4, execute: func(cpu *CPU) {
		fmt.Println("[Z80 UNUSED] 0xFD {IY PREFIX} will cause the Gameboy to crash!")
		cpu.reg.PC++
	}},
}
