package transaction

import (
	"crypto/sha256"
	"fmt"
)

type TX struct {
	From       string
	blicKeyHex string
	To         string
	Amount     float64
	Timestamp  string
	Nonce      int
	TXHash     string
	GasFee     int64
}

func (t *TX) generateHash() string {
	data := fmt.Sprintf("%s%s%.8f%s%d%d",
		t.From,
		t.To,
		t.Amount,
		t.Timestamp,
		t.Nonce,
		t.GasFee)
	sum := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", sum)
}
