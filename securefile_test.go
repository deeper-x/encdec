package securefile

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/h2non/filetype"
)

func TestEncryption(t *testing.T) {
	key := "test"
	_, err := Encrypt("./assets/sample.pdf", []byte(key))
	if err != nil {
		t.Error(err)
	}
}

func TestDecryption(t *testing.T) {
	key := "test"
	_, err := Decrypt("./assets/sample.pdf", []byte(key))
	if err != nil {
		t.Error(err)
	}
}

func TestMultipleDecEnc(t *testing.T) {
	for i := 0; i <= 1000; i++ {
		log.Println("Encryption Nr.", i)
		_, err := Encrypt("./assets/sample.pdf", []byte(""))
		if err != nil {
			t.Error(err)
		}

		ftype, err := checkFileType("./assets/sample.pdf")
		if err != nil {
			t.Error(err)
		}
		log.Println("Encrypted file type is", ftype)

		log.Println("Decryption Nr.", i)
		_, err = Decrypt("./assets/sample.pdf", []byte(""))
		if err != nil {
			t.Error(err)
		}
		ftype, err = checkFileType("./assets/sample.pdf")
		if err != nil {
			t.Error(err)
		}
		log.Println("Decrypted file type is", ftype)

	}
}

func checkFileType(inFile string) (string, error) {
	buf, err := ioutil.ReadFile("./assets/sample.pdf")
	res := "Unknown"

	if err != nil {
		return "", err
	}

	ftype, err := filetype.Match(buf)
	if err != nil {
		return "", err
	}

	if ftype.MIME.Value != "" {
		res = ftype.MIME.Value
	}

	return res, nil
}
