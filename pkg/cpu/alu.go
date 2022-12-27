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

// Inc8 - 8 bit increment
// INC n - Increment register n
// Flags affected:
// Z - Set if result is zero.
// N - Reset.
// H - Set if carry from bit 3.
// C - Not affected.
func (cpu *CPU) Inc8(n *uint8) {
	// Reset flags - except C
	cpu.reg.F &= ^(FlagZ | FlagN | FlagH)

	// Increment n
	result := *n + 1

	// Set flags
	if result == 0 {
		cpu.reg.F |= FlagZ
	}
	if (result & 0xF) == 0x00 {
		cpu.reg.F |= FlagH
	}

	// Set n
	*n = result
}

// Dec8 - 8 bit decrement
// DEC n - Decrement register n
// Flags affected:
// Z - Set if result is zero.
// N - Set.
// H - Set if no borrow from bit 4.
// C - Not affected.
func (cpu *CPU) Dec8(n *uint8) {
	// Reset flags - except C and N
	cpu.reg.F &= ^(FlagZ | FlagH)

	// Set Flag N
	cpu.reg.F |= FlagN

	// Decrement n
	result := *n - 1

	// Set flags
	if result == 0 {
		cpu.reg.F |= FlagZ
	} else if (result & 0xF) == 0x00 {
		cpu.reg.F |= FlagH
	}

	// Set n
	*n = result
}

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
	cpu.reg.F &= ^FlagMask

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

// Sub8 - 8 bit subtraction and 8 bit carry subtraction
// SUB A, n - Subtract n from A
// SBC A, n - Subtract n from A with carry
// Flags affected:
// Z - Set if result is zero.
// N - Set.
// H - Set if no borrow from bit 4.
// C - Set if no borrow.
func (cpu *CPU) Sub8(a *uint8, n uint8, SBC bool) {
	// Get carry flag, for SBC operations
	carry := uint16(0)
	if cpu.reg.F&FlagC != 0 {
		carry = 1
	}

	// Reset flags
	cpu.reg.F &= ^FlagMask

	// Set Flag N
	cpu.reg.F |= FlagN

	// Subtract n from A
	result := uint16(0)
	if SBC {
		result = uint16(*a) - uint16(n) - carry

		// Set flags
		if uint16(*a)&0xF < uint16(n)&0xF+carry {
			cpu.reg.F |= FlagH
		}

	} else {
		result = uint16(*a) - uint16(n)

		// Set flags
		if uint16(*a)&0xF < uint16(n)&0xF {
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

// And8 - 8 bit logical AND
// AND A, n - Logical AND n with A, result in A
// Flags affected:
// Z - Set if result is zero.
// N - Reset.
// H - Set.
// C - Reset.
func (cpu *CPU) And8(a *uint8, n uint8) {
	// Reset flags
	cpu.reg.F &= ^FlagMask

	// Set Flag H
	cpu.reg.F |= FlagH

	// Logical AND n with A
	result := *a & n

	// Set flags
	if result == 0 {
		cpu.reg.F |= FlagZ
	}

	// Set A
	*a = result
}
