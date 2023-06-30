package securefile

import (
	"log"
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

func TestMultipleDecEnc(t *testing.T) {
	for i := 0; i <= 100; i++ {
		log.Println("Encryption: ", i)
		_, err := Encrypt("./assets/sample.pdf", []byte(""))
		if err != nil {
			t.Error(err)
		}
		log.Println("Descryption: ", i)
		_, err = Decrypt("./assets/sample.pdf", []byte(""))
		if err != nil {
			t.Error(err)
		}
	}
}
