package utils

import (
	"math/big"

	"github.com/google/uuid"
)

func IntToUUID(i *big.Int) uuid.UUID {
	buf := i.Bytes()
	if len(buf) < 16 {
		buf = append(make([]byte, 16-len(buf)), buf...)
	}
	return uuid.UUID(buf)
}

func UUIDToInt(u uuid.UUID) *big.Int {
	return new(big.Int).SetBytes(u[:])
}
