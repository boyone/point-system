package main

import (
	"encoding/json"
	"log"
	"net/http"
	"point-system/point"

	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()

	router.Post("/api/v1/calculatePoints", CalculatePointHandler)

	log.Fatal(http.ListenAndServe(":3000", router))
}

type PointResponse struct {
	Value int `json:"value"`
}

type PriceRequest struct {
	Value float64 `json:"value,omitempty"`
}

func CalculatePointHandler(w http.ResponseWriter, r *http.Request) {
	price := PriceRequest{}
	err := json.NewDecoder(r.Body).Decode(&price)
	if err != nil {
		log.Println(err)
	}

	pointResponse := PointResponse{}
	pointResponse.Value = point.CalculatePoint(price.Value)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&pointResponse)
	if err != nil {
		log.Println(err)
	}
}
