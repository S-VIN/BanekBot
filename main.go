/*package main

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


}*/

package main

import (
	"fmt"
	
)

type chat struct{
	favourites map[int]bool
	likes map[int]bool
	dislikes map[int]bool
}

func NewChat() *chat{
	f := make(map[int]bool)
	l := make(map[int]bool)
	d := make(map[int]bool)
	return &chat{f, l, d}
}

type Database struct{
	chats map[int64]chat
}

func NewDatabase() *Database{
	result := make(map[int64] chat)
	return &Database{result}
}

func (database *Database)AddToFavourites(chatID int64, anekID int){
	database = NewDatabase()
	database.chats[chatID] = *NewChat()
	database.chats[chatID].favourites[anekID] = true
}

func (database Database)IsFavourite(chatID int64, anekID int) bool{
	return database.chats[chatID].favourites[anekID]
}

func main() {
	/*d := NewDatabase()
	d.chats[4] = *NewChat()
	d.chats[4].favourites[3] = true*/

	var d Database
	d.AddToFavourites(4, 5)
	d.AddToFavourites(5, 6)
	
	fmt.Println(d.IsFavourite(4, 5))


}