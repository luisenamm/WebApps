package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func vuejs(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "vue01.html")
}

// Film is the struct containing film data
type movie struct {
	Titulo   string
	Director string
	Fecha    int
}

func sendMovies(w http.ResponseWriter, r *http.Request) {
	movies := [3]movie{
		movie{
			Titulo:   "Avengers: Infinity War",
			Director: "Joe Russo, Anthony Russo",
			Fecha:    2018,
		},
		movie{
			Titulo:   "Bastardos Sin Gloria",
			Director: "Quentin Tarantino",
			Fecha:    2009,
		},
		movie{
			Titulo:   "Star Wars: La venganza de los Sith",
			Director: "George Lucas",
			Fecha:    2005,
		},
	}
	json.NewEncoder(w).Encode(movies)

}

func main() {
	http.HandleFunc("/vue", vuejs)
	http.HandleFunc("/movies", sendMovies)
	fmt.Println(time.Now().Format("02-01-2006 15:04:05"))
	err := http.ListenAndServe("localhost"+":"+"8080", nil)
	if err != nil {
		return
	}
}
