package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type DBLoader struct {
	source string
}

func NewDBLoader(source string) *DBLoader {
	return &DBLoader{
		source: source,
	}
}

func (loader DBLoader) Load() (Models, error) {
	jsonFile, err := os.Open(loader.source)
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, err
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var models Models

	err = json.Unmarshal(byteValue, &models)

	if err != nil {
		return nil, err
	}

	return models, nil
}
