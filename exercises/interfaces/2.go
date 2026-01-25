package main

import (
	"fmt"
	"os"
)

type Storage interface {
    Save(data string) error
    Load() (string, error)
}

type FileStorage struct {
	path string
}

type InMemoryStorage struct {
	data string
}

func NewFileStorage(path string) *FileStorage {
	return &FileStorage{path: path}
} 

func (fs *FileStorage) Save(data string) error {
	return os.WriteFile(fs.path, []byte(data), 0644)
}

func (fs *FileStorage) Load() (string, error) {
	data, err := os.ReadFile(fs.path)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{}
}

func (s *InMemoryStorage) Save(data string) error {
	s.data = data
	return nil
}

func (s *InMemoryStorage) Load() (string, error) {
	return s.data, nil
}


func RunApp() {
	mem := NewInMemoryStorage()
	file := NewFileStorage("data.txt")
	
	mem.Save("Hello World")
	data,_ := mem.Load()
	fmt.Println(data)

	file.Save("Hello World")
	data,_ = file.Load()
	fmt.Println(data)
}


