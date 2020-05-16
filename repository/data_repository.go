package repository

import (
	"github.com/gobuffalo/packr/v2"
	scribble "github.com/nanobox-io/golang-scribble"
)

type IPomodoroRepository interface {
	Write(collection, resource string, v interface{}) error
	Read(collection, resource string, v interface{}) error
	ReadAll(s string) ([]string, error)
}

type pomodoroRepository struct {
	data *scribble.Driver
}

func (p pomodoroRepository) ReadAll(s string) ([]string, error) {
	return p.data.ReadAll(s)
}

func (p pomodoroRepository) Write(collection, resource string, v interface{}) error {
	return p.data.Write(collection, resource, v)
}

func (p pomodoroRepository) Read(collection, resource string, v interface{}) error {
	return p.data.Read(collection, resource, v)
}

func NewPomodoroRepository() (IPomodoroRepository, error) {
	dataFolder := packr.New("data", "../data")
	var err error
	data, err := scribble.New(dataFolder.ResolutionDir, nil)
	if err != nil {
		return nil, err
	}
	return &pomodoroRepository{
		data: data,
	}, nil
}
