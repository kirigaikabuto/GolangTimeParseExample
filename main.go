package main

import (
	"GriffonTest/postgres"
	"encoding/json"
	"fmt"
	_ "time"
)
var jsonData = `{"date_of_birth":"2020-01-19","name":"Yerassyl"}`
func main(){
	cfg := postgres.Config{
		Host: "localhost",
		User: "postgres",
		Password: "passanya",
		Port: 5432,
		Database: "griffon",
		ConnectionString: "",
		Params: "sslmode=disable",
	}
	postgreStore, err := postgres.NewPostgreStore(cfg)
	if err !=nil {
		panic(err)
	}
	obj := &postgres.Data{}
	err = json.Unmarshal([]byte(jsonData),obj)
	if err != nil{
		panic(err)
	}
	newobj, err := postgreStore.Create(obj)
	if err != nil {
		panic(err)
	}
	fmt.Println(newobj.Id,newobj.DateOfBirth)


}


