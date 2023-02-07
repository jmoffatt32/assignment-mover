package arguments

import (
	"errors"
	"fmt"
)

type ArgumentsController interface {
	PrintArgs()
	GetFilename() string
	GetClassname() string
}

type Arguments struct {
	Filename  string
	Classname string
}

func NewArguments(raw []string) (*Arguments, error) {

	if len(raw) < 2 {
		return &Arguments{}, errors.New("not enough arguments")
	}

	return &Arguments{
		Filename:  raw[1],
		Classname: raw[2],
	}, nil
}

func (a *Arguments) PrintArgs() {
	fmt.Printf("Filename: %s\n", a.Filename)
	fmt.Printf("Class: %s\n", a.Classname)
}

func (a *Arguments) GetFilename() string {
	return a.Filename
}

func (a *Arguments) GetClassname() string {
	return a.Classname
}
