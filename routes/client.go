package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tiborv/prxy/models"
)

const clientPath = "/api/client"

func init() {
	mux.Handle(clientPath+"/list", RequireAuth(listClient))
	mux.Handle(clientPath+"/update", RequireAuth(updateClient))
	mux.Handle(clientPath+"/create", RequireAuth(createClient))
	mux.Handle(clientPath+"/delete", RequireAuth(deleteClient))

}

func listClient(w http.ResponseWriter, r *http.Request) {
	clients, _ := models.FindAllClients()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(clients)
}

func createClient(w http.ResponseWriter, r *http.Request) {
	client, jsonErr := models.ClientJson(r.Body)
	if jsonErr != nil {
		fmt.Println("Client create jsonErr:", jsonErr)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Client not created")
		return
	}
	savedClient, saveErr := client.Init().Save()
	if saveErr != nil {
		fmt.Println("Client create err:", saveErr)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Client not created")
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(savedClient)
}

func updateClient(w http.ResponseWriter, r *http.Request) {
	client, jsonErr := models.ClientJson(r.Body)
	if jsonErr != nil {
		fmt.Println("Client update jsonErr:", jsonErr)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Client not updated")
		return
	}
	savedClient, saveErr := client.Save()
	if saveErr != nil {
		fmt.Println("Client update err:", saveErr)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Client not updated")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(savedClient)
}

func deleteClient(w http.ResponseWriter, r *http.Request) {
	client, jsonErr := models.ClientJson(r.Body)
	if jsonErr != nil {
		fmt.Println("Client delete jsonErr:", jsonErr)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Client not deleted")
		return
	}
	deleted := client.Delete()
	if deleted {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Client deleted")
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, "Client not deleted")

}
