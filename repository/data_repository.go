package repository

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	packr "github.com/gobuffalo/packr/v2"
	scribble "github.com/nanobox-io/golang-scribble"
	"log"
	"os"
)

type IPomodoroRepository interface {
	Write(collection, resource string, v interface{}) error
	Read(collection, resource string, v interface{}) error
	ReadAll(s string) ([]string, error)
	Close()
}

type pomodoroRepository struct {
	data *scribble.Driver
}

func (p pomodoroRepository) Close() {
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

type boltRepository struct {
	data *bolt.DB
}

func (br boltRepository) Close() {
	br.data.Close()
}

func (br boltRepository) Write(collection, resource string, v interface{}) error {
	confBytes, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("could not marshal %s json: %v", collection, err)
	}
	err = br.data.Update(func(tx *bolt.Tx) error {
		err = tx.Bucket([]byte("pomodoro")).Bucket([]byte(collection)).Put([]byte(resource), confBytes)
		if err != nil {
			return fmt.Errorf("could not set %s: %v", collection, err)
		}
		return nil
	})
	//fmt.Println("Set Config")
	return err
}

func (br boltRepository) Read(collection, resource string, v interface{}) error {
	err := br.data.View(func(tx *bolt.Tx) error {
		v = tx.Bucket([]byte("pomodoro")).Bucket([]byte(collection)).Get([]byte(resource))
		//fmt.Printf("Config: %s\n", v)
		return nil
	})
	return err
}

func (br boltRepository) ReadAll(collection string) ([]string, error) {
	ret := []string{}
	err := br.data.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("pomodoro")).Bucket([]byte(collection))
		b.ForEach(func(k, v []byte) error {
			//fmt.Println(string(k), string(v))
			ret = append(ret, string(v))
			return nil
		})
		return nil
	})
	return ret, err
}

func NewBoltPomodoroRepository() (IPomodoroRepository, error) {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	dataFolder := packr.New("data", "../data")
	db, err := bolt.Open(dataFolder.ResolutionDir+string(os.PathSeparator)+"pomodoro.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		root, err := tx.CreateBucketIfNotExists([]byte("pomodoro"))
		if err != nil {
			return fmt.Errorf("could not create pomodoro bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte("settings"))
		if err != nil {
			return fmt.Errorf("could not create current bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte("current"))
		if err != nil {
			return fmt.Errorf("could not create current bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte("workdone"))
		if err != nil {
			return fmt.Errorf("could not create workdone bucket: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not set up buckets, %v", err)
	}
	fmt.Println("DB Setup Done")

	return &boltRepository{
		data: db,
	}, nil
}


