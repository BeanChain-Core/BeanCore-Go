package transaction

import (
	"testing"
)

func TestGenerateHash(t *testing.T) {
	// positive test case
	tx := &TX{
		From:      "Alice",
		To:        "Bob",
		Amount:    10.5,
		Timestamp: "2024-06-10T12:00:00Z",
		Nonce:     1,
		GasFee:    1000,
	}
	expected := tx.generateHash()
	// Calling again should produce the same hash for the same data
	actual := tx.generateHash()
	if expected != actual {
		t.Errorf("Expected hash %s, got %s", expected, actual)
	}

	// negative test case
	tx1 := &TX{
		From:      "Alice",
		To:        "Bob",
		Amount:    10.5,
		Timestamp: "2024-06-10T12:00:00Z",
		Nonce:     1,
		GasFee:    1000,
	}
	tx2 := &TX{
		From:      "Alice",
		To:        "Charlie", // Different recipient
		Amount:    10.5,
		Timestamp: "2024-06-10T12:00:00Z",
		Nonce:     1,
		GasFee:    1000,
	}
	hash1 := tx1.generateHash()
	hash2 := tx2.generateHash()
	if hash1 == hash2 {
		t.Errorf("Expected different hashes for different transactions, got %s and %s", hash1, hash2)
	}
}
