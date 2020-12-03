package main

//import "fmt"
import "math/rand"
import "strconv"
var database Database

type Anek struct {
	Text string
	Theme string 
	Likes int
	Dislikes int
  }

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
	arrayOfAneks []Anek
	chats map[int64] chat
}

func NewDatabase() *Database{
	result := make(map[int64] chat)
	return &Database{result}
}

func (database * Database) Like(chatID int64, anekNumber int){
	database.chats[chatID].likes[anekNumber] = true
	database.arrayOfAneks[anekNumber].Likes++
}

func (database * Database) Dislike(chatID int64, anekNumber int){
	database.chats[chatID].dislikes[anekNumber] = true
	database.arrayOfAneks[anekNumber].Dislikes++
}

func (database * Database) AddToFavourites(chatID int64, anekNumber int){
	database.chats[chatID].favourites[anekNumber] = true
}

func (database Database) IsFavourite(chatID int64, anekNumber int) bool{
	return database.chats[chatID].favourites[anekNumber]
}

func (database Database) IsLike(chatID int64, anekNumber int) bool{
	return database.chats[chatID].likes[anekNumber]
}

func (database Database) IsDislike(chatID int64, anekNumber int) bool{
	return database.chats[chatID].dislikes[anekNumber]
}

func (database *Database) Add(inputAnek Anek) {
	database.arrayOfAneks = append(database.arrayOfAneks, inputAnek)
}

func (database Database) Get(id int) Anek{
	if id < 0 || id >= len(database.arrayOfAneks){
		return Anek{}
	}
	return database.arrayOfAneks[id]
}

func (database *Database) Setup() (err error){	
	for i := 0; i < 10; i++ {
		text, err := getAnecFromInternet(i)
		if err != nil{
			return err
		}
		database.Add(Anek{Text: text })
	}
	return err
}

func (database Database) GetChat(chatID int64) chat{
	return database.chats[chatID]
}

func (database Database) GetLikedAnek() (Anek, int){
	//i know that sometimes it does not get liked anek
	//it is feature, not bug
	for i:=0; i < len(database.arrayOfAneks) ; i++{
		temp := database.Get(rand.Intn(len(database.arrayOfAneks)))
		if temp.Likes > temp.Dislikes{
			return temp, i
		}
	}
	return Anek{}, 0
}

func (database Database) GetDislikedAnek() (Anek, int){
	//i know that sometimes it does not get liked anek
	//it is feature, not bug
	for i:=0; i < len(database.arrayOfAneks) ; i++{
		temp := database.Get(rand.Intn(len(database.arrayOfAneks)))
		if temp.Likes < temp.Dislikes{
			return temp, i
		}
	}
	return Anek{}, 0
}

func (database Database) GetStringOfFavourites(chatID int64) (string){
	var result string
	for k, _ := range database.chats[chatID].favourites{
		result += strconv.Itoa(k)
		result += " "
	}
	return result
}