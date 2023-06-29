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

func TestCutExtension(t *testing.T) {
	inName := "demo.pdf"
	expected := "demo"

	get := CutExtension(inName)

	if get != expected {
		t.Errorf("%s != %s", get, expected)
	}
}

func TestCreateDecFileName(t *testing.T) {
	expected := "demo_clean.pdf"
	inName := "demo.pdf"

	get := CreateDecFileName(inName)
	if get != expected {
		t.Errorf("%s != %s", get, expected)
	}
}
