package repository

import (
	"github.com/gobuffalo/packr/v2"
	scribble "github.com/nanobox-io/golang-scribble"
)

type ISettingsRepository interface {
	Write(collection, resource string, v interface{}) error
	Read(collection, resource string, v interface{}) error
	ReadAll(s string) ([]string, error)
}

type settingRepository struct {
	settingsData *scribble.Driver
}

func (s3 settingRepository) ReadAll(s string) ([]string, error) {
	return s3.settingsData.ReadAll(s)
}

func (s3 settingRepository) Write(collection, resource string, v interface{}) error  {
	return s3.settingsData.Write(collection, resource, v)
}

func (s3 settingRepository) Read(collection, resource string, v interface{}) error {
	return s3.settingsData.Read(collection, resource, v)
}

func NewSettingsRepository() (ISettingsRepository, error) {
	dataFolder := packr.New("data", "../data")
	var err error
	data, err := scribble.New(dataFolder.ResolutionDir, nil)
	if err != nil {
		return nil, err
	}
	return &settingRepository{
		settingsData:  data,
	}, nil
}
