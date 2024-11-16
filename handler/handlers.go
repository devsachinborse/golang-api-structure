package handler

import (
	"api-structure/models"
	"api-structure/services"
	"encoding/json"
	"net/http"
)

type ReservationHandler struct {
	service *services.ReservationService
}

func NewReservationHandler(service *services.ReservationService) *ReservationHandler {
	return &ReservationHandler{
		service: service,
	}
}

func (h *ReservationHandler) HandleReservation(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var reservation models.Reservation
		if err := json.NewDecoder(r.Body).Decode(&reservation); err != nil {
			http.Error(w, "Invalid request payload", http.StatusInternalServerError)
			return
		}

		if err := h.service.CreateReservation(reservation); err != nil {
			http.Error(w, "Falid to create reservatin", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Reservation created successfully"})
	}
}
