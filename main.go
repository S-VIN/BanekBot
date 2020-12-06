package main

import "fmt"

func main() {

database.Setup()

	err := telegram.CreateBot()
	if err != nil{
		fmt.Println(err)
	}
	telegram.CheckUpdates()


d := NewDatabase()
d.chats[4] = NewChat()
d.chats[4].favourites[3] = true
fmt.Println(d.chats[4].favourites[3])


}