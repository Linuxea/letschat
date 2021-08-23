package secret

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

var myKey = Key()

func Encrypt(content string) string {
	key, _ := hex.DecodeString(myKey)
	plaintext := []byte(content)

	block, _ := aes.NewCipher(key)
	aesGCM, _ := cipher.NewGCM(block)
	nonce := make([]byte, aesGCM.NonceSize())
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}

func Decrypt(encryptContext string) string {
	key, _ := hex.DecodeString(myKey)
	enc, _ := hex.DecodeString(encryptContext)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
}

func Key() string {
	bytes := []byte{116, 34, 62, 54, 227, 137, 201, 159, 246, 168, 86, 243, 239, 146, 186, 229, 162, 240, 83, 190, 250, 135, 185, 236, 142, 97, 40, 42, 223, 68, 1, 94}
	return hex.EncodeToString(bytes) //encode key in bytes to string for saving
}
