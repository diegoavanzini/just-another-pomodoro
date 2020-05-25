package repository

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestRead(t *testing.T) {
	repo, err := NewBoltPomodoroRepository()
	if err != nil {
		log.Fatal(err)
	}
	valueWrite := "testone"
	repo.Write("settings", "test", valueWrite)

	var valueRead string
	err = repo.Read("settings", "test", valueRead)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, valueWrite, valueRead)

}

func TestParam(t *testing.T){
	valueWrite := "testone"
	var valueRead string
	err := myTest("settings", "test", valueRead)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, valueWrite, valueRead)
}

func myTest(s string, s2 string, read interface{}) error {
	read = "testone"
	fmt.Printf("collection:resource:value %s:%s:%s\n", s, s2, read)
	return nil
}

func testFunc(f func(tx *bolt.Tx) error) error {
return nil
}
