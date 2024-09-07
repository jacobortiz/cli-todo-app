package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	FileName string
}

func NewStorage[T any](filename string) *Storage[T] {
	return &Storage[T]{FileName: filename}
}

func (s *Storage[T]) Save(data T) error {
	filedata, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		return err
	}

	return os.WriteFile(s.FileName, filedata, 0644)
}

func (s *Storage[T]) Load(data *T) error {
	filedata, err := os.ReadFile(s.FileName)

	if err != nil {
		return err
	}

	return json.Unmarshal(filedata, data)
}
