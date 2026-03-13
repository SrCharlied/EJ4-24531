package storage

import (
	"encoding/json"
	"os"

	"go-http/interno/modelos"
)

type JSONStore struct {
	FilePath string
	Games    []modelos.Game
}

func NewJSONStore(path string) (*JSONStore, error) {
	store := &JSONStore{
		FilePath: path,
	}

	if err := store.Load(); err != nil {
		return nil, err
	}

	return store, nil
}

func (s *JSONStore) Load() error {
	data, err := os.ReadFile(s.FilePath)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &s.Games)
}

// con esta función se guarda toda la data actualizada que se tenga en memoria, en el archivo json

func (s *JSONStore) Save() error {

	data, err := json.MarshalIndent(s.Games, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.FilePath, data, 0644)
}
