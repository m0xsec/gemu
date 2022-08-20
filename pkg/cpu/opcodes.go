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
	0x00: {name: "NOP", cycles: 4, execute: func(cpu *CPU) {}},

	// 8-bit load instructions
	0x40: {name: "LD B, B", cycles: 4, execute: func(cpu *CPU) { /*cpu.reg.B = cpu.reg.B*/ }},
	0x41: {name: "LD B, C", cycles: 4, execute: func(cpu *CPU) { cpu.reg.B = cpu.reg.C }},
	0x42: {name: "LD B, D", cycles: 4, execute: func(cpu *CPU) { cpu.reg.B = cpu.reg.D }},
	0x43: {name: "LD B, E", cycles: 4, execute: func(cpu *CPU) { cpu.reg.B = cpu.reg.E }},
	0x44: {name: "LD B, H", cycles: 4, execute: func(cpu *CPU) { cpu.reg.B = cpu.reg.H }},
	0x45: {name: "LD B, L", cycles: 4, execute: func(cpu *CPU) { cpu.reg.B = cpu.reg.L }},
	0x46: {name: "LD B, (HL)", cycles: 8, execute: func(cpu *CPU) { cpu.reg.B = cpu.mem.Read(cpu.HL()) }},
	0x47: {name: "LD B, A", cycles: 4, execute: func(cpu *CPU) { cpu.reg.B = cpu.reg.A }},
	0x48: {name: "LD C, B", cycles: 4, execute: func(cpu *CPU) { cpu.reg.C = cpu.reg.B }},
	0x49: {name: "LD C, C", cycles: 4, execute: func(cpu *CPU) { /*cpu.reg.C = cpu.reg.C*/ }},
	0x4A: {name: "LD C, D", cycles: 4, execute: func(cpu *CPU) { cpu.reg.C = cpu.reg.D }},
	0x4B: {name: "LD C, E", cycles: 4, execute: func(cpu *CPU) { cpu.reg.C = cpu.reg.E }},
	0x4C: {name: "LD C, H", cycles: 4, execute: func(cpu *CPU) { cpu.reg.C = cpu.reg.H }},
}
