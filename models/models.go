package models

import "time"

type Reservation struct {
	ID              string    `bson:"id,omitempty"`
	CustomerName    string    `bson:"customer_name,omitempty"`
	TableNumber     int       `bson:"table_number,omitempty"`
	ReservationTime time.Time `bson:"reservation_time,omitempty"`
}
