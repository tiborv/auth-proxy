package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tiborv/prxy/db"
)

const servicePath = "/api/service"

func init() {
	mux.Handle(servicePath+"/list", RequireUser(listService))
	mux.Handle(servicePath+"/update", RequireUser(updateService))
	mux.Handle(servicePath+"/delete", RequireUser(deleteService))
	mux.Handle(servicePath+"/create", RequireUser(createService))
}

func listService(w http.ResponseWriter, r *http.Request) {
	services, _ := db.FindAllServices()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(services)
}

func createService(w http.ResponseWriter, r *http.Request) {
	service, jsonErr := db.ServiceJson(r.Body)
	if jsonErr != nil {
		fmt.Println("Service create json err:", jsonErr)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Service not created, malformed json")
		return
	}
	savedService, saveErr := service.Init().Save()
	if saveErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Service not created")
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(savedService)
}

func updateService(w http.ResponseWriter, r *http.Request) {
	service, jsonErr := db.ServiceJson(r.Body)
	if jsonErr != nil {
		fmt.Println("Service update json err:", jsonErr)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Service not updated, malformed json")
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
	service, jsonErr := db.ServiceJson(r.Body)
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
