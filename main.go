package main

import(
	"fmt"
	"log"
	"endcoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Tittle string `json"tittle"`
	Director *Director `json:director"` //* is a pointer 
}
 
type Director {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`

}

var movie []Movie

fubc getMovie( w http.ResponseWritter. r *http.Request)
	w.Header().Set("Content-Type ". "application/json"
	json.NewEncoder(w).Encode(movie)
)


func main(){
	r :=mux.NewRouter() //function in side mux 

	movie = append(movie , Movie{ID: "1", Isbn:"438227", Title: "Movie one" , Director : &Director{Firstname:" John ", Lastname:"Doe"}})
	movie = append (movies, Movie{ ID:"2", Isbn: "45455",Title: "Movie Two" , Director: &Director{Firstname:"Steve",Lastname:"smith"}})
	r.HandleFunc("/movie",getMovie).methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies/{id}",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Print("Starting server at port 8080\n")
	log.Fatal(http.ListenAndsrve(":8000",r))
}