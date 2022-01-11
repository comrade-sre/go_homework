package main
import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "POST HANDLER")
	}).Methods(http.MethodPost)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "GET HANDLER")
	}).Methods(http.MethodGet)
	router.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintln(w, "GET BY ID HANDLER, RESOURSE IS IS ", vars["id"])
		return
	}).Methods(http.MethodGet)
	router.HandleFunc("/{id}/name/{name}", func(w http.ResponseWriter, r *http.Request){
		vars := mux.Vars(r)
		fmt.Fprintf(w, "GET BY ID HANDLER WITH NAME. RESOURCE ID IS %s AND NAME IS %s\n", vars["id"], vars["name"])
		return
	}).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8020", router))

}
