package restfulencryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// --------------------------------------------------------
// ------------------ AES Encryption Logic ------------------
// --------------------------------------------------------

type EncryptedDataset struct {
	configuration AESConfiguration
	data          []byte
}

// Getters
func (e EncryptedDataset) GetConfiguration() AESConfiguration { return e.configuration }
func (e EncryptedDataset) GetData() []byte                    { return e.data }

// Setters
func (e *EncryptedDataset) SetConfiguration(cfg AESConfiguration) { e.configuration = cfg }
func (e *EncryptedDataset) SetData(d []byte)                      { e.data = d }

// When provided with valid confirguration and plain text, encrypt the datq and return an EncryptedDataset struct
func encrypt(config AESConfiguration, plaintext []byte) (EncryptedDataset, error) {
	block, err := aes.NewCipher([]byte(config.GetKey()))
	if err != nil {
		return EncryptedDataset{}, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return EncryptedDataset{}, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return EncryptedDataset{}, err
	}

	ciphertext := EncryptedDataset{configuration: config, data: gcm.Seal(nonce, nonce, plaintext, nil)}

	return ciphertext, nil
}

// When provided with an EncryptedDataset struct, decrypt the data using the AES key in the configuration
func decrypt(data EncryptedDataset) (string, error) {
	block, err := aes.NewCipher([]byte(data.GetConfiguration().GetKey()))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	ciphertext := data.GetData()
	if len(ciphertext) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce, encryptedMessage := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, encryptedMessage, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
