package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tiborv/prxy/models"
)

const servicePath = "/api/service"

func init() {
	mux.Handle(servicePath+"/list", RequireAuth(listService))
	mux.Handle(servicePath+"/update", RequireAuth(updateService))
	mux.Handle(servicePath+"/delete", RequireAuth(deleteService))
	mux.Handle(servicePath+"/create", RequireAuth(createService))
}

func listService(w http.ResponseWriter, r *http.Request) {
	services, _ := models.FindAllServices()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(services)
}

func createService(w http.ResponseWriter, r *http.Request) {
	service, jsonErr := models.ServiceJson(r.Body)
	if jsonErr != nil {
		fmt.Println("Service create json err:", jsonErr)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Service not created, malformed json")
		return
	}
	if service.Exists() {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Service with this slug already exists")
		return
	}
	savedService, saveErr := service.Save()
	if saveErr != nil {
		fmt.Println(saveErr)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Service not created")
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(savedService)
}

func updateService(w http.ResponseWriter, r *http.Request) {
	service, jsonErr := models.ServiceJson(r.Body)
	if jsonErr != nil {
		fmt.Println("Service update json err:", jsonErr)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Service not updated, malformed json")
		return
	}
	if !service.Exists() {
		fmt.Println("Service update does not exist")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "A service with this slug does not exist")
		return
	}

	_, saveErr := service.Save()
	if saveErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Service not updated")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(service)
}

func deleteService(w http.ResponseWriter, r *http.Request) {
	service, jsonErr := models.ServiceJson(r.Body)
	if jsonErr != nil {
		fmt.Println("Service delete json err:", jsonErr)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Service not deleted, malformed json")
		return
	}
	deleted := service.Delete()
	if deleted {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Service deleted")
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, "Service not deleted")

}
