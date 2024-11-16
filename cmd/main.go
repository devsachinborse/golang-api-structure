package main

import (
	"api-structure/config"
	"api-structure/database"
	"api-structure/handler"
	"api-structure/services"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//load config
	cfg := config.LoadConfig()

	//initialize database
	db := database.ConnectMongoDB(cfg)

	//create service and handlers
	reservationService := services.NewReservationService(db)
	reservationHandler := handler.NewReservationHandler(reservationService)

	//setup routes
	http.HandleFunc("/reservations", reservationHandler.HandleReservation)

	fmt.Printf("Server is running on http://localhost:%s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, nil); err != nil {
		log.Fatal(err)
	}
}
