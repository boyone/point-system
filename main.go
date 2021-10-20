package main

import (
	"encoding/json"
	"net/http"
	"point-system/point"

	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()

	router.Post("/api/v1/calculatePoints", CalculatePointHandler)

	http.ListenAndServe(":3000", router)
}

type PointResponse struct {
	Value int `json:"value"`
}

type PriceRequest struct {
	Value float64 `json:"value,omitempty"`
}

func CalculatePointHandler(w http.ResponseWriter, r *http.Request) {
	price := PriceRequest{}
	json.NewDecoder(r.Body).Decode(&price)

	pointResponse := PointResponse{}
	pointResponse.Value = point.CalculatePoint(price.Value)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&pointResponse)
}
