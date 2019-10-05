package dataconvert

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestJSON(t *testing.T) {
	type FruitBasket struct {
		Name    string
		Fruit   []string
		Id      int64 `json:"ref"`
		Created time.Time
	}

	jsonData := []byte(`
	 {
        "Name": "Standard",
        "Fruit": [
             "Apple",
            "Banana",
            "Orange"
        ],
        "ref": 999,
        "Created": "2018-04-09T23:00:00Z"
    }
	`)
	var basket FruitBasket
	err := json.Unmarshal(jsonData, &basket)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", basket)

}
