package main

import "fmt"



func main() {

database.Setup()

	err := telegram.CreateBot()
	if err != nil{
		fmt.Println(err)
	}
	telegram.CheckUpdates()



}