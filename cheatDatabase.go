package main

//import "fmt"
import "math/rand"
import "strconv"

const anekQuantity = 10

type Database struct{
	arrayOfAneks [anekQuantity]anek
}

func NewDatabase() *Database{
	var d Database
	for i := 0; i < anekQuantity; i++{
		d.arrayOfAneks[i] = *AnekInit(i)
	}
	return &d
}

func (database Database) GetText(id int) anek{
	if id < 0 || id >= len(database.arrayOfAneks){
		return anek{}
	}
	return database.arrayOfAneks[id]
}

func (database Database) GetStringOfFavourites(chatID int64) (string){
	var result string
	for k, _ := range database.chats[chatID].favourites{
		result += strconv.Itoa(k)
		result += " "
	}
	return result
}