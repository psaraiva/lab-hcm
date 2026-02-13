package main

import "math/rand"

type Dice struct {
	FacesQtd int      `json:"faces-qtd"`
	Faces    []string `json:"faces"`
}

func NewDiceDefault() *Dice {
	return &Dice{
		FacesQtd: 6,
		Faces:    []string{"1", "2", "3", "4", "5", "6"},
	}
}

func (d *Dice) Roll() string {
	return d.Faces[rand.Intn(d.FacesQtd)]
}
