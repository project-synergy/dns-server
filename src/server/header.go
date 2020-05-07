package server


type Header struct {
	TransactionID uint16
	Flags Flags
	QDCOUNT uint16
	ANCOUNT uint16
	NSCOUNT uint16
	ARCOUNT uint16

	byteRange posRange
}


func (header *Header) setTransactionID() {

	if header.TransactionID == 0 {
		res.Header.TransactionID = req.Header.TransactionID
	}
}

func (header *Header) setQDCount(queryCount uint16) {
	header.QDCOUNT = queryCount
}

func (header *Header) setANCount(queryCount uint16) {
	header.ANCOUNT = queryCount
}

func (header *Header) setNSCount(queryCount uint16) {
	header.NSCOUNT = queryCount
}

func (header *Header) setARCount(queryCount uint16) {
	header.ARCOUNT = queryCount
}

func (header *Header) toBytes() [12]byte {
	
	header_bytes := [12]byte{}

	transactionID := Convertto2Bytes(res.Header.TransactionID)

	header_bytes[0] = transactionID[0]
	header_bytes[1] = transactionID[1]

	flags := res.Header.Flags.inBytes()

	header_bytes[2] = flags[0]
	header_bytes[3] = flags[1]

	QDCOUNT := Convertto2Bytes(res.Header.QDCOUNT)

	header_bytes[4] = QDCOUNT[0]
	header_bytes[5] = QDCOUNT[1]

	ANCOUNT := Convertto2Bytes(res.Header.ANCOUNT)

	header_bytes[6] = ANCOUNT[0]
	header_bytes[7] = ANCOUNT[1]

	NSCOUNT := Convertto2Bytes(res.Header.NSCOUNT)

	header_bytes[8] = NSCOUNT[0]
	header_bytes[9] = NSCOUNT[1]

	ARCOUNT := Convertto2Bytes(res.Header.ARCOUNT)

	header_bytes[10] = ARCOUNT[0]
	header_bytes[11] = ARCOUNT[1]

	return header_bytes
}

func Convertto2Bytes(data uint16) *[2]byte {
	return &[2]byte{ byte(data>>8) & 0xFF, byte(data) & 0xFF}
}