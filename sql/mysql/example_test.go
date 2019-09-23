package mysql

import (
	"fmt"
	"testing"
)

func TestDBPing(t *testing.T) {
	db, err := GetDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func TestGetSingleResult(t *testing.T) {
	result, err := GetSingleResult(1)
	if err != nil {
		panic(err)
	}
	fmt.Println("result:",result)
}

