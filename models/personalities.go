package models

type Personalitie struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	History string `json:"history"`
}

var Personalities []Personalitie
