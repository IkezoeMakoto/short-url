package lib

import (
	"crypto/rand"
	"encoding/binary"
	"strconv"
)

func GetRand() string {
	var n uint64
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return strconv.FormatUint(n, 36)
}
