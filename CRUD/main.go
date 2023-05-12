package main 

import ("fmt"
		"log"
		"encoding/json"
		"math/rand"
		"net/http"
		"strconv"
		"github.com/gorilla/mux"
)

type movie struct{
	ID string `json="id"`
	isbn string `json="isbn"`
	title string `json="title"`
	director * director `json="director"`
}

type director struct {
	firstName string `json="firstName"`
	lastName string `json="lastName"`
}

var movies[] movies

func main(){
r:=mux.NewRouter()

r.HandleFunc("/movies",getMovie).Methods("GET")
r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
r.HandleFunc("/movies",createMovie).Methods("POST")
r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

fmt.Printf("starting server at port 8000\n")
log.Fatal(http.ListenAndServe(":8000",r))

}