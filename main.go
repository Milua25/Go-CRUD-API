/*
C - Create
R - Read
U - Update
D - Delete
*/

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {

	// adding new movies
	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "438337",
		Title: "Movie One",
		Director: &Director{
			Firstname: "John",
			Lastname:  "Snow",
		},
	})

	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "457897",
		Title: "Movie Two",
		Director: &Director{
			Firstname: "Steven",
			Lastname:  "Smith",
		},
	})

	// Create New Router
	r := mux.NewRouter()

	// Router Functions
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server on port 8080")

	// starting server on port 8080
	log.Fatal(http.ListenAndServe(":8080", r))
}

// Get All Movies Function
func getMovies(w http.ResponseWriter, req *http.Request) {

	// Setting the content-type header
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(movies)
}

// Delete Movie
func deleteMovie(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// Get Movie
func getMovie(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// Create Movie
func createMovie(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var movie Movie
	err := json.NewDecoder(req.Body).Decode(&movie)

	if err != nil {
		log.Fatal(err)
	}

	movie.ID = strconv.Itoa(rand.Intn(100000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

// Update Movie
func updateMovie(w http.ResponseWriter, req *http.Request) {
	// set json content type
	w.Header().Set("Content-Type", "application/json")

	// params
	params := mux.Vars(req)

	// loop over the movies, range
	for index, item := range movies {
		// delete the movie
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			// add a new movie
			var movie Movie
			_ = json.NewDecoder(req.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}
