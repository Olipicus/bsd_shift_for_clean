package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"code.olipicus.com/bsd_shift_for_clean/api/member"
	"code.olipicus.com/go_rest_api/api/person"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", index)

	//REST API For Person
	router.HandleFunc("/person/{id}", person.Handler.GetDataByID).Methods("GET")
	router.HandleFunc("/person", person.Handler.InsertData).Methods("POST")
	router.HandleFunc("/person/{id}", person.Handler.UpdateByID).Methods("PUT")
	router.HandleFunc("/person/{id}", person.Handler.RemoveByID).Methods("DELETE")

	router.HandleFunc("/member/random/{id}", member.Handler.Random).Methods("GET")
	router.HandleFunc("/member/{id}", member.Handler.GetDataByID).Methods("GET")

	log.Println("Server Start ...")
	log.Fatal(http.ListenAndServe(":8081", router))

}

func index(w http.ResponseWriter, r *http.Request) {
	addrs, _ := net.InterfaceAddrs()
	fmt.Fprintf(w, "Hello, %v", addrs)
}
