
package main

import (
    "encoding/json"
	"fmt"
	// "io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"status\": \"OK\"}")
}

//GetBattlefield gets a Battlefield resource by name  
func GetBattlefield(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	json.NewEncoder(w).Encode("GetBattlefield: "+name)
}

//StartBattlefield is
func StartBattlefield(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	json.NewEncoder(w).Encode("StartBattlefield: "+name)
}

//StartNonameBattlefield is
func StartNonameBattlefield(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("StartNonameBattlefield")
}

//ShieldHandler is
func ShieldHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	player := vars["player"]
	status := vars["status"]
	json.NewEncoder(w).Encode("ShieldHandler: "+player+" "+status)
}

//DisqualifyHandler is
func DisqualifyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	player := vars["player"]
	status := vars["status"]
	json.NewEncoder(w).Encode("DisqualifyHandler: "+player+" "+status)
}

func main() {
	fmt.Println("Hello World!")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/health", healthHandler)
	router.HandleFunc("/api/battlefield/{name}", GetBattlefield).Methods("GET")
	router.HandleFunc("/api/start/{name}", StartBattlefield).Methods("GET")
	router.HandleFunc("/api/start", StartNonameBattlefield).Methods("GET")
	router.HandleFunc("/api/{player}/shield/{status}", ShieldHandler).Methods("GET")
	router.HandleFunc("/api/{player}/disqualify/{status}", DisqualifyHandler).Methods("GET")

	staticDir := "/static/"
	router.
        PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))
		

	err := http.ListenAndServe(":8080", router)
	if err != nil {
        log.Fatal("ListenAndServe Error: ", err)
    }
}
