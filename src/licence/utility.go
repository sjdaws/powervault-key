package licence

import (
	"encoding/binary"
)

// bigEndian returns the value as a 32-bit big-endian byte array
func bigEndian(value int) []byte {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, uint32(value))

	return bytes
}

// bigEndian16 returns the value as a 16-bit big-endian byte array
func bigEndian16(value int) []byte {
	bytes := make([]byte, 2)
	binary.BigEndian.PutUint16(bytes, uint16(value))

	return bytes
}
