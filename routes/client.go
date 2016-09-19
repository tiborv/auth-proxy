package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tiborv/auth-proxy/models"
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
		HttpResponse{Status: http.StatusBadRequest, Msg: "Client not created"}.Send(w)
		return
	}
	savedClient, saveErr := client.Init().Save()
	if saveErr != nil {
		fmt.Println("Client create err:", saveErr)
		HttpResponse{Status: http.StatusBadRequest, Msg: "Client not created"}.Send(w)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(savedClient)
}

func updateClient(w http.ResponseWriter, r *http.Request) {
	client, jsonErr := models.ClientJson(r.Body)
	if jsonErr != nil {
		fmt.Println("Client update jsonErr:", jsonErr)
		HttpResponse{Status: http.StatusBadRequest, Msg: "Client not updated"}.Send(w)
		return
	}
	savedClient, saveErr := client.Save()
	if saveErr != nil {
		fmt.Println("Client update err:", saveErr)
		HttpResponse{Status: http.StatusBadRequest, Msg: "Client not updated"}.Send(w)

		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(savedClient)
}

func deleteClient(w http.ResponseWriter, r *http.Request) {
	client, jsonErr := models.ClientJson(r.Body)
	if jsonErr != nil {
		fmt.Println("Client delete jsonErr:", jsonErr)
		HttpResponse{Status: http.StatusBadRequest, Msg: "Client not deleted"}.Send(w)
		return
	}
	deleted := client.Delete()
	if deleted {
		HttpResponse{Status: http.StatusOK, Msg: "Client deleted"}.Send(w)
		return
	}
	HttpResponse{Status: http.StatusBadRequest, Msg: "Client not deleted"}.Send(w)
}
