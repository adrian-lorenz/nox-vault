package cfernet

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/fernet/fernet-go"
	log "github.com/sirupsen/logrus"
	"time"
)

type Encryptor struct {
	Ckey *fernet.Key
}

func NewEncryptor(keyA string) *Encryptor {
	key, err := fernet.DecodeKey(keyA)
	if err != nil {
		log.Errorln("Key Mismatch")
	}
	return &Encryptor{Ckey: key}
}

func (e *Encryptor) Encrypt(message string) string {
	if message == "" {
		return ""
	}
	token, err := fernet.EncryptAndSign([]byte(message), e.Ckey)
	if err != nil {
		log.Errorln(err.Error())
		return ""
	}
	return string(token)
}

func (e *Encryptor) Decrypt(cipherText string) string {
	tenYears := 10 * 365 * 24 * time.Hour
	plainText := fernet.VerifyAndDecrypt([]byte(cipherText), tenYears, []*fernet.Key{e.Ckey})
	return string(plainText)
}

func CreateFernetKey(len int) (string, error) {
	keyLength := len
	key := make([]byte, keyLength)
	_, err := rand.Read(key)
	if err != nil {
		log.Error("Fehler beim Generieren des Schlüssels:", err)
		return "", err
	}
	encodedKey := base64.URLEncoding.EncodeToString(key)
	return encodedKey, nil
}
