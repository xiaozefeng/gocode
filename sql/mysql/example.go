package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	driverName = "mysql"
	username   = "root"
	password   = "root"
	database   = "goserver"
	host       = "localhost:3306"
)

func GetDB() (*sql.DB, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=skip-verify&autocommit=true", username, password, host, database)
	log.Println(url)
	db, err := sql.Open(driverName, url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetSingleResult(id int) (interface{}, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.Prepare("select username from tb_user where id =?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var r string
	err = stmt.QueryRow(id).Scan(&r)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return nil, err
	}
	return r, nil
}
