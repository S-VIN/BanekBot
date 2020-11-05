package main

import (
	"context"
	"fmt"
	"os"
	"github.com/gocolly/colly"
	"strconv"
	"github.com/jackc/pgx"
)


func getAnec(index int) string {
	var output string
	c := colly.NewCollector()
	c.OnHTML("article", func(e *colly.HTMLElement) {
		output = e.Text
	})

	c.Visit("https://baneks.ru/" + strconv.Itoa(index))
	return output
}

func main() {
	conn, err := pgx.Connect(context.Background(), "user=stepan password= host=localhost port=5432 dbname=stepan sslmode=verify-ca")
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	
	defer conn.Close(context.Background())
	//for i:=-5; i < 500; i++{
		//var text string
		//var number int
		//text = getAnec(i)
		//_, err = conn.Exec(context.Background(), "INSERT INTO aneks(number, anek)VALUES ("+ strconv.Itoa(i) + ", '" + text + "');" )
		 rows, _ := conn.Query(context.Background(), "select * from aneks")

		 for rows.Next() {
			 var number int
			 var text string
			 err := rows.Scan(&number, &text)
			 if err == nil{
				fmt.Printf("1")
			 }
			 fmt.Printf("%d. %s\n", number, text)
		 }
		// fmt.Println(number, text)
	//}
}

