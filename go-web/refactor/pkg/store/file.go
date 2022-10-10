package store

import (
	"encoding/json"
	"os"
	"io/ioutil"
)

type Store interface {
	ReadUser(data interface{}) error
	WriteUser(data interface{}) error
}

const (
	FileType  Type = "file"
)

type Type string //Un alias..

type fileStore struct {
	FilePath string
}

func New(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &fileStore{fileName}
	}

	return nil
}

func (fs *fileStore) WriteUser(data interface{}) error {
	dataJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(fs.FilePath, dataJSON, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (fs *fileStore) ReadUser(data interface{}) error {
	var dataByte []byte
	jsonFile, _ := os.Open(fs.FilePath)
	dataByte, _ = ioutil.ReadAll(jsonFile)
	

	if err := json.Unmarshal(dataByte, &data); err != nil {
		return err
	}

	return nil
}