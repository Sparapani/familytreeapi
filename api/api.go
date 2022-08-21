package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Sparapani/familytreeapi/graph"
	"github.com/Sparapani/familytreeapi/services"
	"github.com/Sparapani/familytreeapi/types"
	"github.com/gorilla/mux"
)

func StartRouter() {

	router := mux.NewRouter()

	router.HandleFunc("/familytree/", getFamilyTree).Methods("GET")
	router.HandleFunc("/familytree/relative/", getRelativeByID).Methods("GET")
	router.HandleFunc("/familytree/baconsnumber/", getBaconsNumber).Methods("GET")
	router.HandleFunc("/familytree/member/", addMemberInFamilyTree).Methods("POST")
	router.HandleFunc("/familytree/member/", updateMemberInFamilyTree).Methods("PUT")
	router.HandleFunc("/familytree/member/", removeMemberInFamilyTree).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getFamilyTree(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	members := services.GetAllMembers()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(members)

}

func getRelativeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var member graph.Member

	json.NewDecoder(r.Body).Decode(&member)

	relatives := services.GetRelativesByID(member)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(relatives)

}

func getBaconsNumber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var members types.BaconsNumber

	json.NewDecoder(r.Body).Decode(&members)

	baconsNumber := services.GetBaconsNumber(members)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(baconsNumber)

}

func addMemberInFamilyTree(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newMember graph.Member

	json.NewDecoder(r.Body).Decode(&newMember)

	newMember.ID = services.AddMember(newMember.Name)

	if newMember.ID == -1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrRequiredFieldsEmpty.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newMember.ID)
}

func updateMemberInFamilyTree(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var updateMember graph.Member

	json.NewDecoder(r.Body).Decode(&updateMember)

	ok := services.UpdateMember(updateMember)

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrBadRequiredFields.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("OK")
}

func removeMemberInFamilyTree(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var removeMember graph.Member

	json.NewDecoder(r.Body).Decode(&removeMember)

	ok := services.RemoveMember(removeMember)

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrBadRequiredFields.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("OK")
}
