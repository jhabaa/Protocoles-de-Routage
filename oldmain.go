package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Recette struct {
	idRecette   int
	Name        string
	Ingredients string
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write(([]byte("Hi")))
}
func connection() {
	fmt.Println("Lancement du Serveur...")
	db, err := sql.Open("mysql", "root:hean2000@tcp(127.0.0.1:3306)/GoDB")

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connecté à la BD")
	}
	db.Begin()
}

func insert(w http.ResponseWriter, r *http.Request) {
	db.Exec("INSERT INTO `GoDB`.`Recette` (`Name`, `Ingredients`) VALUES ('Pizza', 'Creme');")
	fmt.Println("Changement effectué")
}

func main() {

	//db.Begin()
	//print("Connecté à la Base de donnée :") //Check code
	//Here we try to insert data to our database. To check connection during development
	//db.Query("INSERT INTO `GoDB`.`Recette` (`Name`, `Ingredients`) VALUES ('Pizza', 'Crême');") //okk
	//	db.SetConnMaxIdleTime(time.Minute * 3)
	//db.SetMaxOpenConns(10)
	//db.SetMaxIdleConns(10)
	connection()
	http.HandleFunc("/insert", insert)
	print("Démarage du serveur WEB")
	//Mise en place du serveur HTTP
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8001", nil)

}
