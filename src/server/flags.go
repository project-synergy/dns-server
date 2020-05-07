package server

import(
	"strconv"
	"fmt"
)

type Flags struct{
	/*
	*	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
	*	|QR|   Opcode  |AA|TC|RD|RA|   Z    |   RCODE   |
	*	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
	*/
	QR uint8
	OPCODE [4]uint8
	AA uint8
	TC uint8
	RD uint8
	
	RA uint8
	Z [3]uint8
	RCODE [4]uint8
}

func (flags *Flags) setQR(qr uint8) {
	flags.QR = qr
}

func (flags *Flags) setOPCODE(opcode [4]uint8) {
	flags.OPCODE = opcode
}

func (flags *Flags) setAA(aa uint8) {
	flags.AA = aa
}

func (flags *Flags) setTC(tc uint8) {
	flags.TC = tc
}

func (flags *Flags) setRD(rd uint8) {
	flags.RD = rd
}

func (flags *Flags) setRA(ra uint8) {
	flags.RA = ra
}

func (flags *Flags) setZ(z [3]uint8) {
	flags.Z = z
}

func (flags *Flags) setRCODE(rcode [4]uint8) {
	flags.RCODE = rcode
}

func (flags *Flags) inBytes() *[2]byte {

	byte1_str := fmt.Sprintf("%d%d%d%d%d%d%d%d",
				flags.QR,
				flags.OPCODE[0],
				flags.OPCODE[1],
				flags.OPCODE[2],
				flags.OPCODE[3],
				flags.AA,
				flags.TC,
				flags.RD,
			)
	byte2_str := fmt.Sprintf("%d%d%d%d%d%d%d%d",
		flags.RA,
		flags.Z[0],
		flags.Z[1],
		flags.Z[2],
		flags.RCODE[0],
		flags.RCODE[1],
		flags.RCODE[2],
		flags.RCODE[3],
	)

	byte1, _ := strconv.ParseInt(byte1_str, 2, 64)
	byte2, _ := strconv.ParseInt(byte2_str, 2, 64)

	return &[2]byte{ byte(uint8(byte1)), byte(uint8(byte2))}
}