package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartRouter() {

	router := mux.NewRouter()

	router.HandleFunc("/familytree/", getFamilyTree).Methods("GET")
	router.HandleFunc("/familytree/", addMemberInFamilyTree).Methods("POST")
	router.HandleFunc("/familytree/{id}", updateMemberInFamilyTree).Methods("PUT")
	router.HandleFunc("/familytree/{id}", removeMemberInFamilyTree).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getFamilyTree(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("OK")

}

func addMemberInFamilyTree(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("OK")
}

func updateMemberInFamilyTree(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("OK")
}

func removeMemberInFamilyTree(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("OK")
}
