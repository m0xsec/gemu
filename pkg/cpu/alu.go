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

// ALU - Arithmetic Logic Unit

// Add8 - 8 bit addition (ADD A,n - Add n to A)
// ADD A, n - Add n to A
// Flags affected:
// Z - Set if result is zero.
// N - Reset.
// H - Set if carry from bit 3.
// C - Set if carry from bit 7.
func (cpu *CPU) Add8(a *uint8, n uint8) {
	// Reset flags
	cpu.reg.F &= FlagMask

	// Add n to A
	result := uint16(*a) + uint16(n)

	// Set flags
	if (result & 0xFF00) != 0 {
		cpu.reg.F |= FlagC
	}
	if (result & 0x00FF) == 0 {
		cpu.reg.F |= FlagZ
	}
	if uint16(*a)&0xF+uint16(n)&0xF > 0xF {
		cpu.reg.F |= FlagH
	}

	// Set A
	*a = uint8(result)
}
