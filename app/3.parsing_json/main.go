package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
)

var exampleResponse1 = `{
  "id": 10,
  "petId": 198772,
  "quantity": 7,
  "shipDate": "2026-02-11T12:34:19.923Z",
  "status": "approved",
  "complete": true
}`

type ExampleResponse1Struct struct {
	ID       int       `json:"id"`
	PetID    int       `json:"petId"`
	Quantity int       `json:"quantity"`
	ShipDate time.Time `json:"shipDate"`
	Status   string    `json:"status"`
	Complete bool      `json:"complete"`
}

func main() {
	var data ExampleResponse1Struct
	err := json.Unmarshal([]byte(exampleResponse1), &data)
	if err != nil {
		spew.Dump(data)
		spew.Dump(err)
		return
	}
	spew.Dump(data)
	spew.Dump(err)

	var data2 ExampleResponse1Struct
	data2.ID = data.ID
	data2.PetID = data.PetID
	data2.Quantity = data.Quantity
	data2.ShipDate = data.ShipDate
	data2.Complete = data.Complete

	result, _ := json.Marshal(data2)
	fmt.Println(string(result))
}
