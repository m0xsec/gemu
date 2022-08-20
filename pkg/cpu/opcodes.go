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

// https://gbdev.io/gb-opcodes/optables/
// https://gbdev.io/gb-opcodes/Opcodes.json
// TODO: Could build the opcode table from the Opcodes.json file?

var opcodes = map[byte]struct {
	name    string
	cycles  uint32
	execute func(cpu *CPU)
}{
	0x00: {name: "NOP", cycles: 4, execute: func(cpu *CPU) { cpu.reg.PC++ }},

	// 8-bit load/store/move instructions
	0x02: {name: "LD (BC),A", cycles: 8, execute: func(cpu *CPU) {
		cpu.mem.Write(cpu.BC(), cpu.reg.A)
		cpu.reg.PC++
	}},
	0x06: {name: "LD B,d8", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.B = cpu.mem.Read(cpu.reg.PC + 1)
		cpu.reg.PC += 2
	}},
	0x0A: {name: "LD A,(BC)", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.A = cpu.mem.Read(cpu.BC())
		cpu.reg.PC++
	}},
	0x0E: {name: "LD C,d8", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.C = cpu.mem.Read(cpu.reg.PC + 1)
		cpu.reg.PC += 2
	}},
	0x12: {name: "LD (DE),A", cycles: 8, execute: func(cpu *CPU) {
		cpu.mem.Write(cpu.DE(), cpu.reg.A)
		cpu.reg.PC++
	}},
	0x16: {name: "LD D,d8", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.D = cpu.mem.Read(cpu.reg.PC + 1)
		cpu.reg.PC += 2
	}},
	0x1A: {name: "LD A,(DE)", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.A = cpu.mem.Read(cpu.DE())
		cpu.reg.PC++
	}},
	0x1E: {name: "LD E,d8", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.E = cpu.mem.Read(cpu.reg.PC + 1)
		cpu.reg.PC += 2
	}},
	0x22: {name: "LD (HL+),A", cycles: 8, execute: func(cpu *CPU) {
		cpu.mem.Write(cpu.HL(), cpu.reg.A)
		cpu.reg.PC++
	}},
	0x26: {name: "LD H,d8", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.H = cpu.mem.Read(cpu.reg.PC + 1)
		cpu.reg.PC += 2
	}},
	0x2A: {name: "LD A,(HL+)", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.A = cpu.mem.Read(cpu.HL())
		cpu.reg.PC++
	}},
	0x2E: {name: "LD L,d8", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.L = cpu.mem.Read(cpu.reg.PC + 1)
		cpu.reg.PC += 2
	}},
	0x32: {name: "LD (HL-),A", cycles: 8, execute: func(cpu *CPU) {
		cpu.mem.Write(cpu.HL(), cpu.reg.A)
		cpu.reg.PC++
	}},
	0x36: {name: "LD (HL),d8", cycles: 8, execute: func(cpu *CPU) {
		cpu.mem.Write(cpu.HL(), cpu.mem.Read(cpu.reg.PC+1))
		cpu.reg.PC += 2
	}},
	0x3A: {name: "LD A,(HL-)", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.A = cpu.mem.Read(cpu.HL())
		cpu.reg.PC++
	}},
	0x3E: {name: "LD A,d8", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.A = cpu.mem.Read(cpu.reg.PC + 1)
		cpu.reg.PC += 2
	}},
	0x40: {name: "LD B, B", cycles: 4, execute: func(cpu *CPU) { /*cpu.reg.B = cpu.reg.B*/ cpu.reg.PC++ }},
	0x41: {name: "LD B, C", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.B = cpu.reg.C
		cpu.reg.PC++
	}},
	0x42: {name: "LD B, D", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.B = cpu.reg.D
		cpu.reg.PC++
	}},
	0x43: {name: "LD B, E", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.B = cpu.reg.E
		cpu.reg.PC++
	}},
	0x44: {name: "LD B, H", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.B = cpu.reg.H
		cpu.reg.PC++
	}},
	0x45: {name: "LD B, L", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.B = cpu.reg.L
		cpu.reg.PC++
	}},
	0x46: {name: "LD B, (HL)", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.B = cpu.mem.Read(cpu.HL())
		cpu.reg.PC++
	}},
	0x47: {name: "LD B, A", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.B = cpu.reg.A
		cpu.reg.PC++
	}},
	0x48: {name: "LD C, B", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.C = cpu.reg.B
		cpu.reg.PC++
	}},
	0x49: {name: "LD C, C", cycles: 4, execute: func(cpu *CPU) { /*cpu.reg.C = cpu.reg.C*/ cpu.reg.PC++ }},
	0x4A: {name: "LD C, D", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.C = cpu.reg.D
		cpu.reg.PC++
	}},
	0x4B: {name: "LD C, E", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.C = cpu.reg.E
		cpu.reg.PC++
	}},
	0x4C: {name: "LD C, H", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.C = cpu.reg.H
		cpu.reg.PC++
	}},
	0x4D: {name: "LD C, L", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.C = cpu.reg.L
		cpu.reg.PC++
	}},
	0x4E: {name: "LD C, (HL)", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.C = cpu.mem.Read(cpu.HL())
		cpu.reg.PC++
	}},
	0x4F: {name: "LD C, A", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.C = cpu.reg.A
		cpu.reg.PC++
	}},
	0x50: {name: "LD D, B", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.D = cpu.reg.B
		cpu.reg.PC++
	}},
	0x51: {name: "LD D, C", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.D = cpu.reg.C
		cpu.reg.PC++
	}},
	0x52: {name: "LD D, D", cycles: 4, execute: func(cpu *CPU) { /*cpu.reg.D = cpu.reg.D*/ cpu.reg.PC++ }},
	0x53: {name: "LD D, E", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.D = cpu.reg.E
		cpu.reg.PC++
	}},
	0x54: {name: "LD D, H", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.D = cpu.reg.H
		cpu.reg.PC++
	}},
	0x55: {name: "LD D, L", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.D = cpu.reg.L
		cpu.reg.PC++
	}},
	0x56: {name: "LD D, (HL)", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.D = cpu.mem.Read(cpu.HL())
		cpu.reg.PC++
	}},
	0x57: {name: "LD D, A", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.D = cpu.reg.A
		cpu.reg.PC++
	}},
	0x58: {name: "LD E, B", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.E = cpu.reg.B
		cpu.reg.PC++
	}},
	0x59: {name: "LD E, C", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.E = cpu.reg.C
		cpu.reg.PC++
	}},
	0x5A: {name: "LD E, D", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.E = cpu.reg.D
		cpu.reg.PC++
	}},
	0x5B: {name: "LD E, E", cycles: 4, execute: func(cpu *CPU) { /*cpu.reg.E = cpu.reg.E*/ cpu.reg.PC++ }},
	0x5C: {name: "LD E, H", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.E = cpu.reg.H
		cpu.reg.PC++
	}},
	0x5D: {name: "LD E, L", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.E = cpu.reg.L
		cpu.reg.PC++
	}},
	0x5E: {name: "LD E, (HL)", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.E = cpu.mem.Read(cpu.HL())
		cpu.reg.PC++
	}},
	0x5F: {name: "LD E, A", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.E = cpu.reg.A
		cpu.reg.PC++
	}},
	0x60: {name: "LD H, B", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.H = cpu.reg.B
		cpu.reg.PC++
	}},
	0x61: {name: "LD H, C", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.H = cpu.reg.C
		cpu.reg.PC++
	}},
	0x62: {name: "LD H, D", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.H = cpu.reg.D
		cpu.reg.PC++
	}},
	0x63: {name: "LD H, E", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.H = cpu.reg.E
		cpu.reg.PC++
	}},
	0x64: {name: "LD H, H", cycles: 4, execute: func(cpu *CPU) { /*cpu.reg.H = cpu.reg.H*/ cpu.reg.PC++ }},
	0x65: {name: "LD H, L", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.H = cpu.reg.L
		cpu.reg.PC++
	}},
	0x66: {name: "LD H, (HL)", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.H = cpu.mem.Read(cpu.HL())
		cpu.reg.PC++
	}},
	0x67: {name: "LD H, A", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.H = cpu.reg.A
		cpu.reg.PC++
	}},
	0x68: {name: "LD L, B", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.L = cpu.reg.B
		cpu.reg.PC++
	}},
	0x69: {name: "LD L, C", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.L = cpu.reg.C
		cpu.reg.PC++
	}},
	0x6A: {name: "LD L, D", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.L = cpu.reg.D
		cpu.reg.PC++
	}},
	0x6B: {name: "LD L, E", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.L = cpu.reg.E
		cpu.reg.PC++
	}},
	0x6C: {name: "LD L, H", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.L = cpu.reg.H
		cpu.reg.PC++
	}},
	0x6D: {name: "LD L, L", cycles: 4, execute: func(cpu *CPU) { /*cpu.reg.L = cpu.reg.L*/ cpu.reg.PC++ }},
	0x6E: {name: "LD L, (HL)", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.L = cpu.mem.Read(cpu.HL())
		cpu.reg.PC++
	}},
	0x6F: {name: "LD L, A", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.L = cpu.reg.A
		cpu.reg.PC++
	}},
	0x70: {name: "LD (HL), B", cycles: 8, execute: func(cpu *CPU) {
		cpu.mem.Write(cpu.HL(), cpu.reg.B)
		cpu.reg.PC++
	}},
	0x71: {name: "LD (HL), C", cycles: 8, execute: func(cpu *CPU) {
		cpu.mem.Write(cpu.HL(), cpu.reg.C)
		cpu.reg.PC++
	}},
	0x72: {name: "LD (HL), D", cycles: 8, execute: func(cpu *CPU) {
		cpu.mem.Write(cpu.HL(), cpu.reg.D)
		cpu.reg.PC++
	}},
	0x73: {name: "LD (HL), E", cycles: 8, execute: func(cpu *CPU) {
		cpu.mem.Write(cpu.HL(), cpu.reg.E)
		cpu.reg.PC++
	}},
	0x74: {name: "LD (HL), H", cycles: 8, execute: func(cpu *CPU) {
		cpu.mem.Write(cpu.HL(), cpu.reg.H)
		cpu.reg.PC++
	}},
	0x75: {name: "LD (HL), L", cycles: 8, execute: func(cpu *CPU) {
		cpu.mem.Write(cpu.HL(), cpu.reg.L)
		cpu.reg.PC++
	}},
	0x77: {name: "LD (HL), A", cycles: 8, execute: func(cpu *CPU) {
		cpu.mem.Write(cpu.HL(), cpu.reg.A)
		cpu.reg.PC++
	}},
	0x78: {name: "LD A, B", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.A = cpu.reg.B
		cpu.reg.PC++
	}},
	0x79: {name: "LD A, C", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.A = cpu.reg.C
		cpu.reg.PC++
	}},
	0x7A: {name: "LD A, D", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.A = cpu.reg.D
		cpu.reg.PC++
	}},
	0x7B: {name: "LD A, E", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.A = cpu.reg.E
		cpu.reg.PC++
	}},
	0x7C: {name: "LD A, H", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.A = cpu.reg.H
		cpu.reg.PC++
	}},
	0x7D: {name: "LD A, L", cycles: 4, execute: func(cpu *CPU) {
		cpu.reg.A = cpu.reg.L
		cpu.reg.PC++
	}},
	0x7E: {name: "LD A, (HL)", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.A = cpu.mem.Read(cpu.HL())
		cpu.reg.PC++
	}},
	0x7F: {name: "LD A, A", cycles: 4, execute: func(cpu *CPU) { /*cpu.reg.A = cpu.reg.A*/ cpu.reg.PC++ }},
	0xE0: {name: "LDH (n), A", cycles: 12, execute: func(cpu *CPU) {
		cpu.mem.Write(0xFF00+uint16(cpu.mem.Read(cpu.reg.PC+1)), cpu.reg.A)
		cpu.reg.PC += 2
	}},
	0xE2: {name: "LD (C), A", cycles: 8, execute: func(cpu *CPU) {
		cpu.mem.Write(0xFF00+uint16(cpu.reg.C), cpu.reg.A)
		cpu.reg.PC += 2
	}},
	0xEA: {name: "LD (a16), A", cycles: 16, execute: func(cpu *CPU) {
		addr := uint16(cpu.mem.Read(cpu.reg.PC+1)) | uint16(cpu.mem.Read(cpu.reg.PC+2))<<8
		cpu.mem.Write(addr, cpu.reg.A)
		cpu.reg.PC += 3
	}},
	0xF0: {name: "LDH A, (a8)", cycles: 12, execute: func(cpu *CPU) {
		addr := uint16(cpu.mem.Read(cpu.reg.PC + 1))
		cpu.reg.A = cpu.mem.Read(0xFF00 + addr)
		cpu.reg.PC += 2
	}},
	0xF2: {name: "LD A, (C)", cycles: 8, execute: func(cpu *CPU) {
		cpu.reg.A = cpu.mem.Read(0xFF00 + uint16(cpu.reg.C))
		cpu.reg.PC += 2
	}},
	0xFA: {name: "LD A, (a16)", cycles: 16, execute: func(cpu *CPU) {
		addr := uint16(cpu.mem.Read(cpu.reg.PC+1)) | uint16(cpu.mem.Read(cpu.reg.PC+2))<<8
		cpu.reg.A = cpu.mem.Read(addr)
		cpu.reg.PC += 3
	}},

	// TODO: 16bit load/store/move instructions
	// ...

	// TODO: 8bit arithmetic/logic instructions
	// ...

	// TODO: 16bit arithmetic/logic instructions
	// ...

	// TODO: 8bit rotation/shift instructions
	// ...

	// TODO: Jumps/calls instructions
	// ...

	// TODO: Misc/control instructions
	// ...
}
