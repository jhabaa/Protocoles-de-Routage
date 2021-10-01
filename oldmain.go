package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write(([]byte("Hi")))
}
func main() {
	fmt.Println("Lancement du Serveur...")
	db, err := sql.Open("mysql", "root:hean2000@tcp(127.0.0.1:3306)/GoDB")

	if err != nil {
		panic(err)
	}
	db.Begin()
	print("Connecté à la Base de donnée :") //Check code
	//Here we try to insert data to our database. To check connection during development
	db.Query("INSERT INTO `GoDB`.`Recette` (`Name`, `Ingredients`) VALUES ('Bolognese', 'Tomate, pates, Viande');") //okk
	db.SetConnMaxIdleTime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	//Mise en place du serveur HTTP
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8001", nil)
	print("Démarage du serveur WEB")

}
