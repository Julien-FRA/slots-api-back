package controller

import (
	"app/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	_ "strconv"

	"github.com/gorilla/mux"
)

func GetAllAppointmentsController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	shop, err := model.GetAllAppointments()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(shop)
	}
}

func GetCustomerAppointmentsController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	appointment, err := model.GetCustomerAppointments(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(appointment)
}

func GetShopAppointmentsController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	appointment, err := model.GetShopAppointments(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(appointment)
}

func CreateAppointmentController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var appointment model.Appointment

	var reqBody struct {
		AppointmentJSON string `json:"appointmentJSON"`
	}

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("Received employee reqBody: %+v\n", reqBody)
	err = json.Unmarshal([]byte(reqBody.AppointmentJSON), &appointment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("Received employee data: %+v\n", appointment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = model.CreateAppointment(appointment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func UpdateAppointmentController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var appointment model.Appointment

	var reqBody struct {
		AppointmentJSON string `json:"appointmentJSON"`
	}

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal([]byte(reqBody.AppointmentJSON), &appointment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("Received shop data: %+v\n", appointment)
	err = model.UpdateAppointment(appointment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func DeleteAppointmentController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)
	idStr := param["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = model.DeleteAppointment(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
