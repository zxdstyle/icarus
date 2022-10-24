package helpers

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

func InetAddr(ipaddr string) uint32 {
	var ret uint32

	ip := net.ParseIP(ipaddr)
	if ip == nil {
		return 0
	}

	if err := binary.Read(bytes.NewBuffer(ip.To4()), binary.BigEndian, &ret); err != nil {
		return 0
	}

	return ret
}

func IPv4Uint32ToString(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}
