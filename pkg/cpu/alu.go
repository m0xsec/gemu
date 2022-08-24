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

// Add8 - 8 bit addition and 8 bit carry addition
// ADD A, n - Add n to A
// ADC A, n - Add n to A with carry
// Flags affected:
// Z - Set if result is zero.
// N - Reset.
// H - Set if carry from bit 3.
// C - Set if carry from bit 7.
func (cpu *CPU) Add8(a *uint8, n uint8, ADC bool) {
	// Get carry flag, for ADC operations
	carry := uint16(0)
	if cpu.reg.F&FlagC != 0 {
		carry = 1
	}

	// Reset flags
	cpu.reg.F &= FlagMask

	// Add n to A
	result := uint16(0)
	if ADC {
		result = uint16(*a) + uint16(n) + carry

		// Set flags
		if uint16(*a)&0xF+uint16(n)&0xF+carry > 0xF {
			cpu.reg.F |= FlagH
		}

	} else {
		result = uint16(*a) + uint16(n)

		// Set flags
		if uint16(*a)&0xF+uint16(n)&0xF > 0xF {
			cpu.reg.F |= FlagH
		}
	}

	// Set remaining flags
	if (result & 0xFF00) != 0 {
		cpu.reg.F |= FlagC
	}
	if (result & 0x00FF) == 0 {
		cpu.reg.F |= FlagZ
	}

	// Set A
	*a = uint8(result)
}
