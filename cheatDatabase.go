package main

//import "fmt"

var database Database

type Anek struct {
	Text string
	Theme string 
	Likes int
	Dislikes int
  }

type chat struct{
	favourite map[int]bool
	likes map[int]bool
	dislikes map[int]bool
}

type Database struct{
	arrayOfAneks []Anek
	chats map[uint64] chat
}

func (database * Database) Like(chatID uint64, anekNumber int){
	database.chats[chatID].likes[anekNumber] = true
	database.arrayOfAneks[anekNumber].Likes++
}

func (database * Database) Dislike(chatID uint64, anekNumber int){
	database.chats[chatID].dislikes[anekNumber] = true
	database.arrayOfAneks[anekNumber].Dislikes++
}

func (database * Database) AddToFavourite(chatID uint64, anekNumber int){
	database.chats[chatID].favourite[anekNumber] = true
}

func (database Database) IsFavourite(chatID uint64, anekNumber int) bool{
	return database.chats[chatID].favourite[anekNumber]
}

func (database Database) IsLike(chatID uint64, anekNumber int) bool{
	return database.chats[chatID].likes[anekNumber]
}

func (database Database) IsDislike(chatID uint64, anekNumber int) bool{
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

/*func (database Database) GetLikedAnek() Anek{
	//i know that sometimes it does not get liked anek
	//it is feature, not bug
	for i:=0; i < len(database.arrayOfAneks) ; i++{
		temp := database.Get(rand.Intn(len(database.arrayOfAneks)))
		if temp.Likes > temp.Dislikes{
			return temp
		}
	}
	return Anek{}
}*/

