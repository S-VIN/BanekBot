package main

//import "fmt"

var database Database


type Anek struct {
	Text string
	Theme string 
	Likes int
	Dislikes int
  }

type Database struct{
	arrayOfAneks []Anek
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

func (database *Database) Like(id int) {
	database.arrayOfAneks[id].Likes++
}

func (database *Database) Dislike(id int) {
	database.arrayOfAneks[id].Dislikes++
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

