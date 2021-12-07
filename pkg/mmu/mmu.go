/*
	 ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄       ▄▄  ▄         ▄
	▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░▌     ▐░░▌▐░▌       ▐░▌
	▐░█▀▀▀▀▀▀▀▀▀ ▐░█▀▀▀▀▀▀▀▀▀ ▐░▌░▌   ▐░▐░▌▐░▌       ▐░▌
	▐░▌          ▐░▌          ▐░▌▐░▌ ▐░▌▐░▌▐░▌       ▐░▌
	▐░▌ ▄▄▄▄▄▄▄▄ ▐░█▄▄▄▄▄▄▄▄▄ ▐░▌ ▐░▐░▌ ▐░▌▐░▌       ▐░▌
	▐░▌▐░░░░░░░░▌▐░░░░░░░░░░░▌▐░▌  ▐░▌  ▐░▌▐░▌       ▐░▌
	▐░▌ ▀▀▀▀▀▀█░▌▐░█▀▀▀▀▀▀▀▀▀ ▐░▌   ▀   ▐░▌▐░▌       ▐░▌
	▐░▌       ▐░▌▐░▌          ▐░▌       ▐░▌▐░▌       ▐░▌
	▐░█▄▄▄▄▄▄▄█░▌▐░█▄▄▄▄▄▄▄▄▄ ▐░▌       ▐░▌▐░█▄▄▄▄▄▄▄█░▌
	▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░▌       ▐░▌▐░░░░░░░░░░░▌
	▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀         ▀  ▀▀▀▀▀▀▀▀▀▀▀
					the GameBoy Emulator
							m0x <3
*/
package mmu

/* https://gbdev.io/pandocs/Memory_Map.html

The Game Boy has a 16-bit address bus, which is used to address ROM, RAM, and I/O.

Start	End		Description						Notes
0000	3FFF	16 KiB ROM bank 00				From cartridge, usually a fixed bank
4000	7FFF	16 KiB ROM Bank 01~NN			From cartridge, switchable bank via mapper (if any)
8000	9FFF	8 KiB Video RAM (VRAM)			In CGB mode, switchable bank 0/1
A000	BFFF	8 KiB External RAM				From cartridge, switchable bank if any
C000	CFFF	4 KiB Work RAM (WRAM)
D000	DFFF	4 KiB Work RAM (WRAM)			In CGB mode, switchable bank 1~7
E000	FDFF	Mirror of C000~DDFF (ECHO RAM)	Nintendo says use of this area is prohibited.
FE00	FE9F	Sprite attribute table (OAM)
FEA0	FEFF	Not Usable						Nintendo says use of this area is prohibited
FF00	FF7F	I/O Registers
FF80	FFFE	High RAM (HRAM)
FFFF	FFFF	Interrupt Enable register (IE)

*/

// MMU is the Memory Management Unit. While the GameBoy did not have an actual
// MMU, it makes sense for our emulator. The GameBoy uses Memory Mapping to talk to
// various subsystems. The MMU will be responsible for handling that mapping and will
// be the only thing to actually access the memory directly.
type MMU struct {
	// The GameBoy has a memory map from 0x0000 - 0xFFFF
	memory [0xFFFF]byte

	// TODO: Have different mapped sections of memory defined here?
	// HighRAM, OAM, ROM Banks, etc?
}

// Initializes the MMU
func (*MMU) init() {

}

// TODO: Write functions for reading and writing to memory, while handling and respecting
// memory mapping rules
