package routers

import (
	"encoding/json"
	"net/http"

	"github.com/bagomkarnath/go_code/assignments/assignmenttwo/domain"
	"github.com/gorilla/mux"
)

// Handler struct
type Handler struct {
	Repository domain.CustomerStore // interface for persistence
}

// Create Post impl
func (h *Handler) Create(rw http.ResponseWriter, r *http.Request) {
	var customer domain.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := h.Repository.Create(customer); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}

// Update Put impl
func (h *Handler) Update(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var customer domain.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := h.Repository.Update(id, customer); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
}

// Get 1 record Get
func (h *Handler) Get(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if customer, err := h.Repository.GetByID(id); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	} else {
		j, err := json.Marshal(customer)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(j)
	}
}

// GetAll all data Get
func (h *Handler) GetAll(rw http.ResponseWriter, r *http.Request) {
	if customers, err := h.Repository.GetAll(); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	} else {
		j, err := json.Marshal(customers)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(j)
	}
}

// Delete a record
func (h *Handler) Delete(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if err := h.Repository.Delete(id); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
}
