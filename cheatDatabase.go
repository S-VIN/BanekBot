package main

//import "fmt"
import "math/rand"
import "strconv"

const anekQuantity = 1023

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

func (database Database) GetAnekText(id int) string{
	if id < 0 || id >= len(database.arrayOfAneks){
		return ""
	}
	return database.arrayOfAneks[id].text
}

func (database Database)GetRandomAnek() (string, int){
	index := rand.Intn(len(database.arrayOfAneks))
	return database.GetAnekText(index), index
}

func (database Database)GetRandomLikedAnek(chatID int64) (string, int){
	var arrayLikedAnek []string
	var arrayLikedAnekIndex []int
	for i, item := range(database.arrayOfAneks){
		if item.IsLike(chatID){
			arrayLikedAnek = append(arrayLikedAnek, item.text)
			arrayLikedAnekIndex = append(arrayLikedAnekIndex, i)
		}
	}
	if len(arrayLikedAnek) == 0{
		return "", 0
	}
	index := rand.Intn(len(arrayLikedAnek))
	return arrayLikedAnek[index], arrayLikedAnekIndex[index]
}

func (database Database)GetRandomDislikedAnek(chatID int64) (string, int){
	var arrayDislikeAnek []string
	var arrayDislikedAnekIndex []int
	for i, item := range(database.arrayOfAneks){
		if item.IsDislike(chatID){
			arrayDislikeAnek = append(arrayDislikeAnek, item.text)
			arrayDislikedAnekIndex = append(arrayDislikedAnekIndex, i)
		}
	}
	if len(arrayDislikeAnek) == 0{
		return "", 0
	}
	index := rand.Intn(len(arrayDislikeAnek))
	return arrayDislikeAnek[index], arrayDislikedAnekIndex[index]
}

func (database Database)GetRandomFavouriteAnek(chatID int64) (string, int){
	var arrayFavouriteAnek []string
	var arrayFavouriteAnekIndex []int
	for i, item := range(database.arrayOfAneks){
		if item.IsFavorite(chatID){
			arrayFavouriteAnek = append(arrayFavouriteAnek, item.text)
			arrayFavouriteAnekIndex = append(arrayFavouriteAnekIndex, i)
		}
	}
	if len(arrayFavouriteAnek) == 0{
		return "", 0
	}
	index := rand.Intn(len(arrayFavouriteAnek))
	return arrayFavouriteAnek[index], arrayFavouriteAnekIndex[index]
}

func (database Database) GetStringOfFavourites(chatID int64) (string){
	var result string
	for i, item := range(database.arrayOfAneks){
		if item.IsFavorite(chatID){
			result += strconv.Itoa(i) 
			result += " "
		}
	}
	return result
}

func (database Database) IsLike(chatID int64, numberOfAnek int) bool{
	return database.arrayOfAneks[numberOfAnek].IsLike(chatID)
}

func (database Database) IsDislike(chatID int64, numberOfAnek int) bool{
	return database.arrayOfAneks[numberOfAnek].IsDislike(chatID)
}

func (database Database) IsFavourite(chatID int64, numberOfAnek int) bool{
	return database.arrayOfAneks[numberOfAnek].IsFavorite(chatID)
}

func (database *Database) Like(chatID int64, numberOfAnek int){
	database.arrayOfAneks[numberOfAnek].AddLike(chatID)
}

func (database *Database) Dislike(chatID int64, numberOfAnek int){
	database.arrayOfAneks[numberOfAnek].AddDislike(chatID)
}

func (database *Database) Favourite(chatID int64, numberOfAnek int){
	database.arrayOfAneks[numberOfAnek].AddFavorite(chatID)
}