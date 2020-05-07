package server

import (
	"encoding/hex"
	"strconv"
	_ "fmt"
)

func getQTYPEvalues(buf *[]byte) string {

	code, _ := strconv.ParseUint(hex.EncodeToString(*buf), 16, 16)


	switch uint16(code) {
	
		case 1: return "A"
		case 2: return "NS"
		case 3: return "MD"
		case 4: return "MF"
		case 5: return "CNAME"
		case 6: return "SOA"
		case 7: return "MB"
		case 8: return "MG"
		case 9: return "MR"
		case 10: return "NULL"
		case 11: return "WKS"
		case 12: return "PTR"
		case 13: return "HINFO"
		case 14: return "MINFO"
		case 15: return "MX"
		case 16: return "MX"
		default: return ""
		
	}
}