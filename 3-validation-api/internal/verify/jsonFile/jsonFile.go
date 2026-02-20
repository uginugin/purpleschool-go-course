package jsonfile

import (
	"encoding/json"
	"io"
	"os"
)

type JsonFile struct {
	Path string
	Data MailHash
}

type MailHash map[string]string

func New() (*JsonFile, error) {
	json := JsonFile{
		Path: "/Users/uginugin/Desktop/go test/3-validation-api/internal/verify/jsonFile/file.json",
		Data: make(MailHash),
	}

	if err := json.readData(); err != nil {
		return nil, err
	}

	return &json, nil

}

func (j *JsonFile) readData() error {
	jsonFile, err := os.Open(j.Path)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonData, &j.Data); err != nil {
		return err
	}

	return nil
}

func (j *JsonFile) WriteData() error {
	jsonData, err := json.Marshal(j.Data)
	if err != nil {
		return err
	}

	if err := os.WriteFile(j.Path, jsonData, 0644); err != nil {
		return err
	}

	return nil
}
