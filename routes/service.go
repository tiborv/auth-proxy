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
	service, saveErr := db.Service{}.Init().Save()
	if saveErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Service not created")
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(service)
}

func updateService(w http.ResponseWriter, r *http.Request) {
	service, err := db.ServiceJson(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Service not updated")
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
	service, err := db.ServiceJson(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Service not deleted")
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
