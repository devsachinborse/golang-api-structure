package services

import (
	"api-structure/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type ReservationService struct {
	collection *mongo.Collection
}

func NewReservationService(db *mongo.Database) *ReservationService {
	return &ReservationService{
		collection: db.Collection("reservations"),
	}
}

func (s *ReservationService) CreateReservation(reservation models.Reservation) error {
	reservation.ReservationTime = time.Now()
	_, err := s.collection.InsertOne(context.Background(), reservation)
	return err
}
