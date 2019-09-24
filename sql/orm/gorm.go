package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// migrate the schema
	db.AutoMigrate(&Product{})

	// create
	db.Create(&Product{
		Code:  "L1001",
		Price: 1000,
	})

	var product Product
	db.First(&product, 1)
	log.Printf("%+v\n", product)
	db.First(&product, "code=?", "L1001")
	log.Printf("%+v\n", product)

	// update
	//db.Model(&product).Update("Price", 2000)

}

type Product struct {
	gorm.Model
	Code string
	Price uint
}









