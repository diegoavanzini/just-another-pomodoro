package repository

import (
	models "bitbucket.org/avanz/anotherPomodoro/model"
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	packr "github.com/gobuffalo/packr/v2"
	"log"
	"os"
)

type IPomodoroRepository interface {
	Write(collection, key string, value interface{}) error
	Read(collection, key string, value interface{}) error
	ReadAll(collection, filter string) ([]string, error)
	Close()
}

const DB = "pomodoro"

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
		err = tx.Bucket([]byte(DB)).Bucket([]byte(collection)).Put([]byte(resource), confBytes)
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
		value := tx.Bucket([]byte(DB)).Bucket([]byte(collection)).Get([]byte(resource))
		err := json.Unmarshal(value, v)
		return err
	})
	fmt.Printf("collection:resource:value %s:%s:%s\n", collection, resource, v)
	return err
}

func (br boltRepository) ReadAll(collection, filter string) ([]string, error) {
	ret := []string{}
	err := br.data.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(DB)).Bucket([]byte(collection))
		b.ForEach(func(k, v []byte) error {
			pomodoro := models.PomodoroPosition{}
			if err := json.Unmarshal(v, &pomodoro); err != nil {
				return err
			}
			if pomodoro.TimeStarted.Format("20060102") == filter {
				//fmt.Println(string(k), string(v))
				ret = append(ret, string(v))
			}
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
