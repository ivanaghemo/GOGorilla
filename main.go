package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

/*
Las Cloud APIs son interfaces de programación de aplicaciones (API)
que se utilizan para construir aplicaciones en el mercado de la computación en la nube. ​
Las Cloud APIs permiten al software solicitar los datos y los cálculos de uno o más
servicios a través de una interfaz directa o indirecta.
*/
func GETUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde método GET")
}
func POSTUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde método POST")
}
func PUTUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde método PUT")
}
func DELETEUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde método DELETE")
}

func main() {

	port := os.Getenv("PORT")
	r := mux.NewRouter().StrictSlash(false)
	/*
		Cuando StrictSlash(false), las dos rutas siguientes son distintas
		/api/user
		/api/user/
	*/
	r.HandleFunc("/users", GETUsers).Methods("GET")
	r.HandleFunc("/users", POSTUsers).Methods("POST")
	r.HandleFunc("/users", PUTUsers).Methods("PUT")
	r.HandleFunc("/users", DELETEUsers).Methods("DELETE")

	server := &http.Server{
		Addr:           port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening...")
	server.ListenAndServe()

}
