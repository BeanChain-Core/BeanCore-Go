package transaction

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

// TX represents a transaction in the blockchain.
type TX struct {
	From         string  `json:"from"`
	To           string  `json:"to"`
	Amount       float64 `json:"amount"`
	Timestamp    string  `json:"timestamp"`
	Nonce        int     `json:"nonce"`
	TXHash       string  `json:"tx_hash"`
	GasFee       int64   `json:"gas_fee"`
	PublicKeyHex string  `json:"public_key"`
}

// GenerateHash generates a SHA-256 hash for the transaction.
func (t *TX) GenerateHash() string {
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

// MarshalJSON marshals the transaction to JSON format.
func (t *TX) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(*t)
	if err != nil {
		return nil, fmt.Errorf("error marshaling transaction: %w", err)
	}
	return b, nil
}

// UnmarshalJSON unmarshals the transaction from JSON format.
func (t *TX) UnmarshalJSON(data []byte) error {
	type TmpJson TX
	var tmpJson TmpJson

	err := json.Unmarshal(data, &tmpJson)
	if err != nil {
		return err
	}

	// TODO: validate
	*t = TX(tmpJson)
	return nil
}
