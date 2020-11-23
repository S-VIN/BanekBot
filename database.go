package main

import "github.com/jackc/pgx"
import "strconv"


var Conn *pgx.Conn 


func ConnectToDatabase() error{
	Conn, err := pgx.Connect(context.Background(), "user=stepan password= host=localhost port=5432 dbname=stepan sslmode=verify-ca")
	if err != nil {
		return err
	}
	defer Conn.Close(context.Background())
	return nil
}


func GetAnecFromDatabase(index int) (string, error){
	var number int
	var anek string
	//_, err = conn.Exec(context.Background(), "INSERT INTO aneks(number, anek)VALUES ("+ strconv.Itoa(i) + ", '" + text + "');" )
	err := Conn.QueryRow(context.Background(), "select number, anek from aneks where number="+strconv.Itoa(index)+";").Scan(&number, &anek)
	if err != nil {
		return "", err
	}
	return anek, nil
}