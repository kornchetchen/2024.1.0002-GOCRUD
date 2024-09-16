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

func main(){
	r :=mux.NewRouter() //function in side mux 

	r.HandleFunc 
}