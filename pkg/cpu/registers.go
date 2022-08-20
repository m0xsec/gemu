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

// Some 8-bit registers can be used as 16-bit registers.
// AF, BC, DE, HL are such registers.

func (cpu *CPU) AF() uint16 {
	return uint16(cpu.reg.A)<<8 | uint16(cpu.reg.F)
}

func (cpu *CPU) BC() uint16 {
	return uint16(cpu.reg.B)<<8 | uint16(cpu.reg.C)
}

func (cpu *CPU) DE() uint16 {
	return uint16(cpu.reg.D)<<8 | uint16(cpu.reg.E)
}

func (cpu *CPU) HL() uint16 {
	return uint16(cpu.reg.H)<<8 | uint16(cpu.reg.L)
}

func (cpu *CPU) SetAF(val uint16) {
	cpu.reg.A = uint8(val >> 8)
	cpu.reg.F = uint8(val)
}

func (cpu *CPU) SetBC(val uint16) {
	cpu.reg.B = uint8(val >> 8)
	cpu.reg.C = uint8(val)
}

func (cpu *CPU) SetDE(val uint16) {
	cpu.reg.D = uint8(val >> 8)
	cpu.reg.E = uint8(val)
}

func (cpu *CPU) SetHL(val uint16) {
	cpu.reg.H = uint8(val >> 8)
	cpu.reg.L = uint8(val)
}

// TODO: Add flag helper functions?
