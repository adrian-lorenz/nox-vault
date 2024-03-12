package crsa

import (
	"crypto/rand"
	"crypto/rsa"
)

func GenKey() (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func Encrypt(plaintext []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plaintext)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

func Decrypt(ciphertext []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	decryptedText, err := privateKey.Decrypt(rand.Reader, ciphertext, &rsa.PKCS1v15DecryptOptions{})
	if err != nil {
		return nil, err
	}
	return decryptedText, nil
}
