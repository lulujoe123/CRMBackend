package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	// Import the third-party gorilla/mux package
    "github.com/gorilla/mux"
	"github.com/google/uuid"
)

type Customer struct {
	Id			string `json:"id"`
	Name		string `json:"name"`
	Role 		string `json:"role"`
	Email 		string `json:"email"`
	Phone		string `json:"phone"`
	Contacted	bool   `json:"contacted"`
}

var customers = map[string]Customer{
	"63807978-849d-11ed-a1eb-0242ac120002": {
		Id: "63807978-849d-11ed-a1eb-0242ac120002",
		Name: "Kiera Maddox",
		Role: "Software Engineer",
		Email: "test1@gmail.com",
		Contacted: false,
		Phone: "001",
	},
	"63808292-849d-11ed-a1eb-0242ac120002": {
		Id: "63808292-849d-11ed-a1eb-0242ac120002",
		Name: "Sharon Oneal",
		Role: "Software Engineer",
		Email: "test2@gmail.com",
		Contacted: false,
		Phone: "002",
	},
	"638088aa-849d-11ed-a1eb-0242ac120002": {
		Id: "638088aa-849d-11ed-a1eb-0242ac120002",
		Name: "Shania Parrish",
		Role: "Software Engineer",
		Email: "test3@gmail.com",
		Contacted: true,
		Phone: "003",
	},
	"6381acc6-849d-11ed-a1eb-0242ac120002": {
		Id: "6381acc6-849d-11ed-a1eb-0242ac120002",
		Name: "Sharon Oneal",
		Role: "Software Engineer",
		Email: "test4@gmail.com",
		Contacted: false,
		Phone: "004",
	},
	"6381b180-849d-11ed-a1eb-0242ac120002": {
		Id: "6381b180-849d-11ed-a1eb-0242ac120002",
		Name: "Ifan Duffyl",
		Role: "Software Engineer",
		Email: "test5@gmail.com",
		Contacted: true,
		Phone: "005",
	},
}

var emptyMap = map[string]string{}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(customers)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")

	if _, ok := customers[mux.Vars(r)["id"]]; ok {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customers[mux.Vars(r)["id"]])
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(emptyMap)
	}

}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	
	var customer Customer
	customer.Id = uuid.New().String()
    reqBody, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(reqBody, &customer)

	customers[customer.Id] = customer
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customers[customer.Id])

}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")

	var customer Customer
    reqBody, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(reqBody, &customer)
	id := mux.Vars(r)["id"]
	if _, ok := customers[id]; ok {
		customers[id] = customer
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customers[id])
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(emptyMap)
	}
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")

	id := mux.Vars(r)["id"]

	if _, ok := customers[id]; ok {
		delete(customers, id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customers)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(emptyMap)
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	fmt.Println("Server is starting on port 3000...")

	http.ListenAndServe(":3000", router)
}