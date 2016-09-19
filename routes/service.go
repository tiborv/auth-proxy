package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tiborv/auth-proxy/models"
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
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(services)
}

func createService(w http.ResponseWriter, r *http.Request) {
	service, jsonErr := models.ServiceJson(r.Body)
	if jsonErr != nil {
		fmt.Println("Service create json err:", jsonErr)
		HttpResponse{Status: http.StatusBadRequest, Msg: "Service not created, malformed json"}.Send(w)
		return
	}
	if service.Exists() {
		HttpResponse{Status: http.StatusBadRequest, Msg: "Service with this slug already exists"}.Send(w)
		return
	}
	savedService, saveErr := service.Save()
	if saveErr != nil {
		fmt.Println(saveErr)
		HttpResponse{Status: http.StatusBadRequest, Msg: "Service not created"}.Send(w)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(savedService)
}

func updateService(w http.ResponseWriter, r *http.Request) {
	service, jsonErr := models.ServiceJson(r.Body)
	if jsonErr != nil {
		fmt.Println("Service update json err:", jsonErr)
		HttpResponse{Status: http.StatusBadRequest, Msg: "Service not updated, malformed json"}.Send(w)
		return
	}
	if !service.Exists() {
		fmt.Println("Service update does not exist")
		HttpResponse{Status: http.StatusBadRequest, Msg: "A service with this slug does not exist"}.Send(w)
		return
	}

	_, saveErr := service.Save()
	if saveErr != nil {
		HttpResponse{Status: http.StatusBadRequest, Msg: "Service not updated"}.Send(w)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(service)
}

func deleteService(w http.ResponseWriter, r *http.Request) {
	service, jsonErr := models.ServiceJson(r.Body)
	if jsonErr != nil {
		fmt.Println("Service delete json err:", jsonErr)
		HttpResponse{Status: http.StatusBadRequest, Msg: "Service not deleted, malformed json"}.Send(w)
		return
	}
	deleted := service.Delete()
	if !deleted {
		HttpResponse{Status: http.StatusBadRequest, Msg: "Service not deleted"}.Send(w)
		return
	}
	HttpResponse{Status: http.StatusOK, Msg: "Service deleted"}.Send(w)
}
