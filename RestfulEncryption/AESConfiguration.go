package restfulencryption

import (
	"encoding/base64"
	"errors"
	"log"
)

// --------------------------------------------------------
// -------------------- AES Keys Logic --------------------
// --------------------------------------------------------
type AESConfiguration struct {
	byteSize int
	mode     string
	size     int
	Key      string
}

// Getters
func (c AESConfiguration) GetByteSize() int { return c.byteSize }
func (c AESConfiguration) GetMode() string  { return c.mode }
func (c AESConfiguration) GetSize() int     { return c.size }
func (c AESConfiguration) GetKey() string   { return c.Key }

// Setters (use pointer receiver to mutate)
func (c *AESConfiguration) SetByteSize(n int) { c.byteSize = n }
func (c *AESConfiguration) SetMode(m string)  { c.mode = m }
func (c *AESConfiguration) SetSize(s int)     { c.size = s }
func (c *AESConfiguration) SetKey(k string)   { c.Key = k }

// Approved sizes and modes for AES
var approvedByteSizes = []int{16, 24, 32}
var approvedAesModes = []string{"GCM", "CBC", "CFB", "OFB"} //Options for the allowance to encryption expansion in the future
var approvedAesSizes = []int{128, 192, 256}

// --------------------------------------------------------

func NewAESConfiguration(byteSize int) (*AESConfiguration, error) {
	if byteSize != 16 && byteSize != 24 && byteSize != 32 {
		return nil, errors.New("Invalid AES key size. Must be 16, 24, or 32 bytes.")
	}
	cfg := &AESConfiguration{byteSize: byteSize}
	return cfg, nil
}

// Attempt to generate a new AES key of specified byte size (16, 24, or 32 bytes)
// throw an error if there is a problem with the generation
func (config AESConfiguration) GenerateKey() AESConfiguration {

	//Check for valid key size before generating
	if !config.ValidateAESKeySize() {
		log.Fatalf("Invalid AES key size")
		return config // Return the config even if there is an error
	}

	// The output string will be longer, but the decoded key will be exactly 32 bytes.
	// Base64-encoded 32 bytes results in 44 characters (with padding).
	// Calculated: 4 * (ceil(32 / 3)) = 4 * 11 = 44.
	// We will want to chop off the characters to meet the size requirements due to this string size increase
	key := make([]byte, config.byteSize)
	config.Key = base64.StdEncoding.EncodeToString(key)[:config.byteSize] //Chop off the key strings to meet the size requirements for AES
	return config
}

// This is a validation algorithm to ensure that the byte size is valid
// The function will return true if the byte size is valid
// The function will return false if the byte size is not valid or is nil
// All valid sizes are in the approvedByteSizes array constants
func (config AESConfiguration) ValidateAESKeySize() bool {
	//never go forward if there is nothing to compare
	if config.byteSize != 0 {
		// Validate the byte size that is set in the configuration
		for _, size := range approvedByteSizes {
			if config.byteSize == size {
				return true
			}
		}
	}
	//If nothing hits then this will be false
	return false
}
