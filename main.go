package main

import "fmt"

var database Database

func main() {

	database = *NewDatabase()

	err := telegram.CreateBot()
	if err != nil{
		fmt.Println(err)
	}
	telegram.CheckUpdates()
}