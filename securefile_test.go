package securefile

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/h2non/filetype"
)

func TestEncryption(t *testing.T) {
	key := "test"
	_, err := Encrypt("./assets/DELEGA_DA_FIRMARE.pdf", []byte(key))
	if err != nil {
		t.Error(err)
	}
}

func TestDecryption(t *testing.T) {
	key := "test"
	_, err := Decrypt("./assets/DELEGA_DA_FIRMARE.pdf", []byte(key))
	if err != nil {
		t.Error(err)
	}
}

func TestMultipleDecEnc(t *testing.T) {
	for i := 0; i <= 1000; i++ {
		log.Println("Encryption: ", i)
		_, err := Encrypt("./assets/DELEGA_DA_FIRMARE.pdf", []byte(""))
		if err != nil {
			t.Error(err)
		}

		ftype, err := checkFileType("./assets/DELEGA_DA_FIRMARE.pdf")
		if err != nil {
			t.Error(err)
		}
		log.Println("Encrypted type is", ftype)

		log.Println("Decryption: ", i)
		_, err = Decrypt("./assets/DELEGA_DA_FIRMARE.pdf", []byte(""))
		if err != nil {
			t.Error(err)
		}
		ftype, err = checkFileType("./assets/DELEGA_DA_FIRMARE.pdf")
		if err != nil {
			t.Error(err)
		}
		log.Println("Decrypted type is", ftype)

	}
}

func checkFileType(inFile string) (string, error) {
	buf, err := ioutil.ReadFile("./assets/DELEGA_DA_FIRMARE.pdf")
	if err != nil {
		return "", err
	}

	ftype, err := filetype.Match(buf)
	if err != nil {
		return "", err
	}

	return ftype.MIME.Value, nil
}
