package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

var contacts []Contact

func main() {
	r := mux.NewRouter()
	contacts = append(contacts, Contact{Name: "Friend_1", Phone: "98xxx-xxxxx", Email: "person1@mail.com"})
	contacts = append(contacts, Contact{Name: "Friend_2", Phone: "96xxx-xxxxx", Email: "person2@mail.com"})
	contacts = append(contacts, Contact{Name: "Friend_3", Phone: "97xxx-xxxxx", Email: "person3@mail.com"})

	// Route handles & endpoints
	r.HandleFunc("/contacts", getContacts).Methods("GET")
	r.HandleFunc("/contacts/{name}", getContact).Methods("GET")
	// r.HandleFunc("/contacts", createContact).Methods("POST")
	// r.HandleFunc("/contacts/{name}", updateContact).Methods("PUT")
	// r.HandleFunc("/contacts/{name}", deleteContact).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":3000", r))
}

// Get all contacts
func getContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}

func getContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params

	for _, item := range contacts {
		if item.Name == params["name"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
	json.NewEncoder(w).Encode(&Contact{})
}

func createContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var contact Contact
	_ = json.NewDecoder(r.Body).Decode(&contact)
	contacts = append(contacts, contact)
	json.NewEncoder(w).Encode(contact)
}

// Delete contact
func deleteContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range contacts {
		if item.Name == params["name"] {
			contacts = append(contacts[:idx], contacts[idx+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(contacts)
}

// Update contact
func updateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range contacts {
		if item.Name == params["name"] {
			contacts = append(contacts[:idx], contacts[idx+1:]...)
			var contact Contact
			_ = json.NewDecoder(r.Body).Decode(&contact)
			contact.Name = params["name"]
			contacts = append(contacts, contact)
			json.NewEncoder(w).Encode(contact)
			return
		}
	}
}
