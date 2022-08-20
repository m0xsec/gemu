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

// Gameboy CPU Registers
// Most registers are 8bit, but some can be used as 16bit.
// AF, BC, DE, HL are such registers.
//
// A - Accumulator (Used for arithmetic operations)
// F - Flags
// B - B General Purpose (Can be used as 16 bit register - BC)
// C - C General Purpose (Can be used as 16 bit register - BC)
// D - D General Purpose (Can be used as 16 bit register - DE)
// E - E General Purpose (Can be used as 16 bit register - DE)
// H - H General Purpose (Can be used as 16 bit register - HL)
// L - L General Purpose (Can be used as 16 bit register - HL)
// SP - Stack Pointer
// PC - Program Counter
type Registers struct {
	A, F, B, C, D, E, H, L uint8
	SP, PC                 uint16
}

// Gameboy Flags Register
//
// F bits:
// Bit: 7 | 6 | 5 | 4 | 3 | 2 | 1 | 0 |
// Val: Z | N | H | C | 0 | 0 | 0 | 0 |
//
// Bit 0 - Unused (always 0)
// Bit 1 - Unused (always 0)
// Bit 2 - Unused (always 0)
// Bit 3 - Unused (always 0)
// Bit 4 - Carry Flag (N)
// Bit 5 - Half Carry Flag (H)
// Bit 6 - Subtract Flag (N)
// Bit 7 - Zero Flag (Z)
const (
	FlagZ      = uint8(1 << 7)
	FlagN      = uint8(1 << 6)
	FlagH      = uint8(1 << 5)
	FlagC      = uint8(1 << 4)
	FlagMask   = uint8(FlagZ | FlagN | FlagH | FlagC)
	FlagUnused = uint8(0x00)
)

// Get the value of the 16bit AF register
func (cpu *CPU) AF() uint16 {
	return uint16(cpu.reg.A)<<8 | uint16(cpu.reg.F)
}

// Get the value of the 16bit BC register
func (cpu *CPU) BC() uint16 {
	return uint16(cpu.reg.B)<<8 | uint16(cpu.reg.C)
}

// Get the value of the 16bit DE register
func (cpu *CPU) DE() uint16 {
	return uint16(cpu.reg.D)<<8 | uint16(cpu.reg.E)
}

// Get the value of the 16bit HL register
func (cpu *CPU) HL() uint16 {
	return uint16(cpu.reg.H)<<8 | uint16(cpu.reg.L)
}

// Set the value of the 16bit AF register
func (cpu *CPU) SetAF(val uint16) {
	cpu.reg.A = uint8(val >> 8)
	cpu.reg.F = uint8(val)
}

// Set the value of the 16bit BC register
func (cpu *CPU) SetBC(val uint16) {
	cpu.reg.B = uint8(val >> 8)
	cpu.reg.C = uint8(val)
}

// Set the value of the 16bit DE register
func (cpu *CPU) SetDE(val uint16) {
	cpu.reg.D = uint8(val >> 8)
	cpu.reg.E = uint8(val)
}

// Set the value of the 16bit HL register
func (cpu *CPU) SetHL(val uint16) {
	cpu.reg.H = uint8(val >> 8)
	cpu.reg.L = uint8(val)
}

// TODO: Add flag helper functions?
