package restfulencryption

import "fmt"

//Note: This file is intended for the purpose to allow a user to encrypt data
//With a Rest Api Interface using Basic Encryption methods
//This is a simple implementation of a web server that will allow a user to
//send a request to encrypt or decrypt data using a simple algorithm
//This is not intended to be used in production as it is not secure
//This is for educational purposes only

//The web server will listen on port 8080 and will have the following endpoints:
//GET / - returns a simple message indicating the server is running
//POST /encrypt - accepts a JSON payload with the data to encrypt and returns the encrypted data
//POST /decrypt - accepts a JSON payload with the data to decrypt and returns the decrypted data

//The JSON payload for the encrypt and decrypt endpoints will be in the following format:
//{
//    "data": "string to encrypt or decrypt"
//}

//The response will be in the following format:
//{
//    "result": "encrypted or decrypted string"
//}

//Example usage:
//To encrypt data, send a POST request to /encrypt with the JSON payload
//To decrypt data, send a POST request to /decrypt with the JSON payload

//Note: This is a very basic implementation and does not include error handling or validation
//It is recommended to add these features for a more robust application

//This file is a placeholder for the web controller implementation
//The actual implementation will be done in a separate file

//This file is part of the RestfulEncryption package
//It is intended to be used in conjunction with other files in the package
//to provide a complete encryption solution

func ExecuteBasicEncryptionApplicationDemo() {
	fmt.Println("Starting Basic Encryption  Application...")
	fmt.Println("Generating AES Key...")
	//Generate a new AES key of specified byte size (16, 24, or 32 bytes)
	aesConfig := AESConfiguration{byteSize: 32, mode: "GCM", size: 256}

	//Proper error handling for key generation, although this should always skip
	//unless overrighting the aboce value with an invalid size
	aesConfig = aesConfig.GenerateKey()

	fmt.Println("AES Key Generated Successfully.")
	fmt.Println("Key (Base64): " + aesConfig.Key)

	sampleData := "Hello, World! This is a test of AES encryption."
	fmt.Println("Encrypting Sample Data: " + sampleData)

	encryptedData, err := encrypt(aesConfig, []byte(sampleData))
	if err != nil {
		fmt.Println("Error encrypting data:", err)
		return
	}
	fmt.Println("Data Encrypted Successfully.")
	fmt.Println("Encrypted Data (Base64):", encryptedData.data)

	fmt.Println("\nDecrypting Sample Data...")
	decryptedData, err := decrypt(encryptedData)
	if err != nil {
		fmt.Println("Error decrypting data:", err)
		return
	}
	fmt.Println("Data Decrypted Successfully.")
	fmt.Println("Decrypted Data:", decryptedData)

	if decryptedData == sampleData {
		fmt.Println("Success: Decrypted data matches original!")
	} else {
		fmt.Println("Error: Decrypted data does not match original.")
	}

	fmt.Println("Basic Encryption Application Demo Completed.")
}
