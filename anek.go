package main

func IsInSlice(a []int64, x int64) bool {
    for _, n := range a {
        if x == n {
            return true
        }
    }
    return false
}

type anek struct{
	text string
	like []int64
	dislike []int64
	favourite []int64
	allLikes int
	allDislikes int
}

func AnekInit(numOfAnek int) *anek{
	l := make([]int64, 0)
	d := make([]int64, 0)
	f := make([]int64, 0)
	text, _ := getAnecFromInternet(numOfAnek)
	return &anek{text, l, d, f, 0, 0}
}

func (a *anek) AddFavorite(chatID int64){
	if IsInSlice(a.favourite, chatID){
		return
	} 
	a.favourite = append(a.favourite, chatID)
}

func (a *anek) AddLike(chatID int64){
	if IsInSlice(a.like, chatID){
		return
	} 
	a.like = append(a.like, chatID)
	a.allLikes++
}

func (a *anek) AddDislike(chatID int64){
	if IsInSlice(a.dislike, chatID){
		return
	} 
	a.dislike = append(a.dislike, chatID)
	a.allDislikes++
}

func (a anek) IsFavorite(chatID int64) bool{
	return IsInSlice(a.favourite, chatID)
}

func (a anek) IsLike(chatID int64) bool{
	return IsInSlice(a.like, chatID)
}

func (a anek) IsDislike(chatID int64) bool{
	return IsInSlice(a.dislike, chatID)
}

func (a anek) GetLikes() int{
	return a.allLikes
}

func (a anek) GetDislikes() int{
	return a.allDislikes
}

func (a anek) GetText() string{
	return a.text
}