package main

import(
	"endcoding/json"
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
	Tittle string `json:"tittle"`
	Director *Director `json:"director"` //* is a pointer 
}
 
type Director {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`

}

var movies []Movie

fubc getMovie( w http.ResponseWritter. r *http.Request)
	w.Header().Set("Content-Type ". "application/json"
	json.NewEncoder(w).Encode(movie)
)

func deleteMovie( w http.ResponseWriter , r *http.Request)
	w.Header().set("Content-type","application/json")
	parms := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"]
			movies = append(movies[:index], movies[index+1:]...)
			break
	}
	json.NewEncoder(w).Encode(movies)

func getMovie(w http.ResponseWriter , r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return 
		}
	}
	http.NotFound(w.r)
}

func createMovie(w http.ResponseWriter , r *http.Request){
	w.Header().Set("content-Type","application/json")
	var movie Movie
	if err = json.NewDecoder(r.Body).Decode(&movie); err != nil{
		http.Error(w,err,Error() , http.StatusBadRequest)
		return
	}
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies,movie)
	json.NewEncoder(w).Encode(movie)
}
func updateMovie(w http.ResponseWriter, *http.Request) {
	//json content type
	w.Header().Set("content-Type","application/json")
	//params
	params := mux.Vars(r)
	//loop over the movies.range
	//delete the movie with the i.d that you've sent
	//add a new movie -the movie that we send in the body of postman
	for index,item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index]. movie[index+1:]...)
			var movie Movie 
			if err := json.NewDecoder(r.Body).Decode(&movie); err != mil{
				http.Error(w,err.Error() , http.StatusBadRequest)
				return
			}
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies,movie)
			json.NewEncoder(w).Encode(movie)
			return

}

	}http.NotFound(w,r)
}

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
	log.Fatal(http.ListenAndsrve(":8080",r))
}