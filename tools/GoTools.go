package tools

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"io/fs"
	"net/mail"
	"os"
	"strconv"
	"time"
)

//v 1.2

func FileExist(pathA string) bool {
	_, err2 := os.Stat(pathA)
	if err2 != nil {
		return false
	}
	return !os.IsNotExist(err2)
}

func FileInfo(pathA string) fs.FileInfo {
	finfo, err2 := os.Stat(pathA)
	if err2 != nil {
		return nil
	}
	return finfo
}

func FormatTime(timeA time.Time) string {
	return timeA.Format("02.01.2006 15:04:05")
}

func TimeString(t *time.Time) string {
	if t != nil {
		return t.String()
	}
	return ""
}
func ReadFile(pathA string) []byte {
	if !FileExist(pathA) {
		log.Error("File not exist", pathA)
		return nil
	}
	file, errFile := os.Open(pathA)
	if errFile != nil {
		log.Error("error read file")

	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	bytes, _ := io.ReadAll(file)
	return bytes
}

func IsNumericAndLength(s string, le int) (int, bool) {
	cv, err := strconv.Atoi(s)
	if err != nil {
		return 0, false
	}
	return cv, err == nil && len(s) == le
}

func IsNumeric(s string) (int, bool) {
	cv, err := strconv.Atoi(s)
	if err != nil {
		return 0, false
	}
	return cv, err == nil
}

func WriteFile(pathB string, data []byte) error {
	file, err := os.Create(pathB)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func FromBase64(msg string) string {
	decodedBytes, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		fmt.Println("Fehler beim Dekodieren:", err)
		return ""
	}

	return string(decodedBytes)
}

func ToBool(what string) bool {
	cs, err := strconv.ParseBool(what) // convert to bool
	if err != nil {
		return false
	}
	return cs
}

func ParseJson(buffer *bytes.Buffer) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal(buffer.Bytes(), &data)
	if err != nil {
		return nil, errors.New("konnte die Antwort nicht als JSON lesen")
	}
	return data, nil
}

func JsonToBytes(daten map[string]interface{}) *bytes.Buffer {
	saJson, err := json.Marshal(daten)
	if err != nil {
		return new(bytes.Buffer)
	}
	return bytes.NewBuffer(saJson)
}
