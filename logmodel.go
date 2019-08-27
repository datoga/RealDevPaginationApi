package main

import (
	"sort"
	"time"
)

type Fields struct {
	Name    string    `json:"name"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

type Model struct {
	Model  string `json:"model"`
	Pk     int    `json:"pk"`
	Fields Fields `json:"fields"`
}

type Models []Model

func (models *Models) Sort() {
	sort.Slice(*models, func(i, j int) bool {
		return (*models)[i].Pk > (*models)[j].Pk
	})
}

func (models Models) GetFields() []Fields {
	fields := []Fields{}

	for _, model := range models {
		fields = append(fields, model.Fields)
	}

	return fields
}
