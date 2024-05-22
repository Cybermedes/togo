package main

import (
	"encoding/json"
	"errors"
	"os"
)

// Load the TODO list from a json file
func (t *Todos) LoadFromFile(filename string) error {

	data, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(data) == 0 {
		return err
	}

	err = json.Unmarshal(data, t)
	if err != nil {
		return err
	}
	return nil
}

// Save the updated TODO list to the json file
func (t *Todos) SaveToFile(filename string) error {

	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
