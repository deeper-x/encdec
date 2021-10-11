package securefile

import (
	"testing"
)

func TestEncryption(t *testing.T) {
	key := "test"
	_, err := Encrypt("./assets/demo.png", []byte(key))
	if err != nil {
		t.Error(err)
	}
}

func TestDecryption(t *testing.T) {
	key := "test"
	_, err := Decrypt("./assets/demo.png", []byte(key))
	if err != nil {
		t.Error(err)
	}
}
