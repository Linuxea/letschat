package secret

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	encrypt := Encrypt("abcdefg12345")
	fmt.Println(encrypt)
	decrypt := Decrypt(encrypt)
	fmt.Println(decrypt)
}
