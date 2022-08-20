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

var Opcodes = map[byte]struct {
	Name    string
	Cycles  uint32
	Execute func(cpu *CPU)
}{
	0x00: {Name: "NOP", Cycles: 4, Execute: func(cpu *CPU) {}},
}
