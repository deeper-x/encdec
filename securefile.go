package securefile

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

// Encrypt encrypts a file using AES-GCM.
func Encrypt(source string, password []byte) (bool, error) {
	// we need file to exist
	if _, err := os.Stat(source); os.IsNotExist(err) {
		return false, err
	}

	// read file
	plaintext, err := os.ReadFile(source)

	if err != nil {
		return false, err
	}

	// generate salt
	key := password
	nonce := make([]byte, 12)

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return false, err
	}

	// generate key
	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		return false, err
	}

	// generate gcm cipher
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return false, err
	}

	// generate ciphertext
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	// Append the nonce to the end of file
	ciphertext = append(ciphertext, nonce...)

	// build encrypted file
	f, err := os.Create(source)
	if err != nil {
		return false, err
	}

	// write to file encrypted content
	_, err = io.Copy(f, bytes.NewReader(ciphertext))
	if err != nil {
		return false, err
	}

	return true, nil
}

// Decrypt decrypts a file using AES-GCM.
func Decrypt(source string, password []byte) (bool, error) {
	// get file information, we need to know if file exists
	if _, err := os.Stat(source); os.IsNotExist(err) {
		return false, err
	}

	// read file
	ciphertext, err := os.ReadFile(source)

	if err != nil {
		return false, err
	}

	// generate salt
	key := password
	salt := ciphertext[len(ciphertext)-12:]
	str := hex.EncodeToString(salt)

	// get decoded string
	nonce, err := hex.DecodeString(str)

	if err != nil {
		return false, err
	}

	// generate key
	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)

	// generate block
	block, err := aes.NewCipher(dk)
	if err != nil {
		return false, err
	}

	// generate gcm cipher
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return false, err
	}

	// generate plaintext
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext[:len(ciphertext)-12], nil)
	if err != nil {
		return false, err
	}

	// build decrypted file
	newName := CreateDecFileName(source)
	f, err := os.Create(newName)
	if err != nil {
		return false, err
	}

	// write to file decrypted content
	_, err = io.Copy(f, bytes.NewReader(plaintext))
	if err != nil {
		return false, err
	}

	return true, nil
}

// CutExtension remove extension from file name
func CutExtension(fname string) string {
	return strings.TrimSuffix(fname, filepath.Ext(fname))
}

// CreateDecFileName write filename with _clean.extension format
func CreateDecFileName(fname string) string {
	base := CutExtension(fname)
	ext := filepath.Ext(fname)

	return fmt.Sprintf("%s_clean%s", base, ext)

}
