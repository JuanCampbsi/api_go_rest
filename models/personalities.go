package models

type Personalitie struct {
	Name    string `json:"name"`
	History string `json:"history"`
}

var Personalities []Personalitie
