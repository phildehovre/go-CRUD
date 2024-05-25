package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json: "id"`
	Isbn     string    `json: "isbn"`
	title    string    `json: "title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "438227", title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "437459", title: "Movie Two", Director: &Director{Firstname: "Jane", Lastname: "Doe"}})

	// func getMovies() {
	// 	fmt.Printf("Getting movies %s", movies)
	// }

	// func getMovie(id) {
	// 	fmt.Printf("Getting mmovie: %s", movies[ID: id])
	// }

	// func createMovie(data) {
	// 	fmt.Printf("Creating movie: %s", data)
	// }

	// r.HandleFunc("/movies", getMovies).Methods("GET")
	// r.HandleFunc("/movies/id", getMovie).Methods("GET")
	// r.HandleFunc("/movies", createMovie).Methods("POST")
	// r.HandleFunc("/movies", updateMovie).Methods("PUT")
	// r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	// fmt.Printf("Starting server at port 8000\n")
	// log.Fatal(http.ListenAndServe(":8000", r))

	for v := 0; v < len(movies); v++ {
		fmt.Println(movies[v])
		if movies[v].title == "Movie One" {
			fmt.Printf("5")
		}
	}
}
