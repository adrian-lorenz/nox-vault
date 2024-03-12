package crsa

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func SavePrivateKeyToPEM(fileName string, key *rsa.PrivateKey) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	err = pem.Encode(file, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	if err != nil {
		return err
	}
	return nil
}

func SavePublicKeyToPEM(fileName string, key *rsa.PublicKey) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	pubBytes, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		return err
	}
	err = pem.Encode(file, &pem.Block{Type: "PUBLIC KEY", Bytes: pubBytes})
	if err != nil {
		return err
	}
	return nil
}

func LoadPrivateKeyFromPEM(fileName string) (*rsa.PrivateKey, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	info, _ := file.Stat()
	key := make([]byte, info.Size())
	_, err = file.Read(key)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(key)
	parsedKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return parsedKey, nil
}

func LoadPublicKeyFromPEM(fileName string) (*rsa.PublicKey, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	info, _ := file.Stat()
	key := make([]byte, info.Size())
	_, err = file.Read(key)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(key)
	parsedKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	rsaKey, ok := parsedKey.(*rsa.PublicKey)
	if !ok {
		return nil, err
	}
	return rsaKey, nil
}
func PrivateKeyToPEMString(key *rsa.PrivateKey) (string, error) {
	// Private Key im PEM-Format codieren
	privBytes := x509.MarshalPKCS1PrivateKey(key)
	privPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privBytes,
	})
	return string(privPEM), nil
}

func PublicKeyToPEMString(key *rsa.PublicKey) (string, error) {
	// Public Key im PEM-Format codieren
	pubBytes, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		return "", err
	}
	pubPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubBytes,
	})
	return string(pubPEM), nil
}

func PrivateKeyFromPEMString(keyString string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(keyString))
	if block == nil {
		return nil, nil
	}
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privKey, nil
}

func PublicKeyFromPEMString(keyString string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(keyString))
	if block == nil {
		return nil, nil
	}
	ifc, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	rsaKey, ok := ifc.(*rsa.PublicKey)
	if !ok {
		return nil, nil
	}
	return rsaKey, nil
}