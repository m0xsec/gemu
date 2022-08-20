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

// stackPush pushes a value onto the stack.
func (cpu *CPU) stackPush(b uint8) {
	cpu.reg.SP--
	cpu.mem.Write(cpu.reg.SP, b)
	fmt.Printf("[Stack] Write: %02x to %x\n", b, cpu.reg.SP)
}

// stackPop pops a value from the stack.
func (cpu *CPU) stackPop() uint8 {
	b := cpu.mem.Read(cpu.reg.SP)
	cpu.reg.SP++
	fmt.Printf("[Stack] Read: %02x from %x\n", b, cpu.reg.SP)
	return b
}
