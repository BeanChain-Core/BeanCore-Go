package transaction

import (
	"strings"
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
	expected := tx.GenerateHash()
	// Calling again should produce the same hash for the same data
	actual := tx.GenerateHash()
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
	hash1 := tx1.GenerateHash()
	hash2 := tx2.GenerateHash()
	if hash1 == hash2 {
		t.Errorf("Expected different hashes for different transactions, got %s and %s", hash1, hash2)
	}
}

func TestMarshalJSON(t *testing.T) {
	tx := &TX{
		From:         "Alice",
		To:           "Bob",
		Amount:       42.0,
		Timestamp:    "2024-06-10T12:00:00Z",
		Nonce:        7,
		TXHash:       "somehash",
		GasFee:       1000,
		PublicKeyHex: "abcdef123456",
	}
	data, err := tx.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON returned error: %v", err)
	}
	// Check that the output contains expected fields
	jsonStr := string(data)
	if !containsAll(jsonStr, []string{
		`"from":"Alice"`,
		`"to":"Bob"`,
		`"amount":42`,
		`"timestamp":"2024-06-10T12:00:00Z"`,
		`"nonce":7`,
		`"tx_hash":"somehash"`,
		`"gas_fee":1000`,
		`"public_key":"abcdef123456"`,
	}) {
		t.Errorf("MarshalJSON output missing expected fields: %s", jsonStr)
	}
}

func TestUnmarshalJSON(t *testing.T) {
	// Positive test case: valid JSON
	jsonData := `{
		"from": "Alice",
		"to": "Bob",
		"amount": 99.99,
		"timestamp": "2024-06-10T12:00:00Z",
		"nonce": 42,
		"tx_hash": "hashvalue",
		"gas_fee": 1234,
		"public_key": "deadbeef"
	}`
	tx := &TX{}
	err := tx.UnmarshalJSON([]byte(jsonData))
	// err := json.Unmarshal([]byte(jsonData), &tx)
	if err != nil {
		t.Fatalf("UnmarshalJSON failed on valid input: %v", err)
	}
	if tx.From != "Alice" || tx.To != "Bob" || tx.Amount != 99.99 ||
		tx.Timestamp != "2024-06-10T12:00:00Z" || tx.Nonce != 42 ||
		tx.TXHash != "hashvalue" || tx.GasFee != 1234 || tx.PublicKeyHex != "deadbeef" {
		t.Errorf("UnmarshalJSON did not populate struct correctly: %+v", tx)
	}

	// Negative test case: invalid JSON
	invalidJSON := `{"from": "Alice", "amount": "not-a-number"}`
	var tx2 TX
	err = tx2.UnmarshalJSON([]byte(invalidJSON))
	if err == nil {
		t.Errorf("UnmarshalJSON should fail on invalid input, but got no error")
	}

	// Negative test case: completely malformed JSON
	malformedJSON := `{this is not valid json}`
	var tx3 TX
	err = tx3.UnmarshalJSON([]byte(malformedJSON))
	if err == nil {
		t.Errorf("UnmarshalJSON should fail on malformed JSON, but got no error")
	}
}

// containsAll checks if all substrings are present in s.
func containsAll(s string, subs []string) bool {
	for _, sub := range subs {
		if !strings.Contains(s, sub) {
			return false
		}
	}
	return true
}
