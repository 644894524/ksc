package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Uid  	 int `json:"uid"`
	Name 	 string `json:"name"`
	Age  	 int
	ClassId  int
	SchoolId int
	Phone    string
}

func main(){
	info := User{Uid: 12343, Name: "zhangsan", Age: 18, ClassId: 20, SchoolId: 20, Phone: "13621075519"}
	jsonInfo, _ := json.Marshal(info)
	fmt.Println(string(jsonInfo))
}

