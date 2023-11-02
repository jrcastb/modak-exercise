package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/time/rate"
	"log"
	"net/http"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func endpointHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	message := Message{Status: "successful", Body: "La notificación ha sido entregada correctamente."}

	err := json.NewEncoder(w).Encode(&message)
	if err != nil {
		return
	}
}

func rateLimiter(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	limiter := rate.NewLimiter(2, 4)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			message := Message{
				Status: "failed",
				Body:   "La notificación no ha podido entregarse",
			}
			w.WriteHeader(http.StatusTooManyRequests)
			err := json.NewEncoder(w).Encode(&message)
			if err != nil {
				log.Println("Ocurrió un error codificando el mensaje")
				return
			}
			return
		} else {
			next(w, r)
		}
	},
	)
}

func main() {
	http.Handle("/ping", rateLimiter(endpointHandler))
	err := http.ListenAndServe(":8080", nil)
	fmt.Println("Servidor escuchando en el puerto 8080")
	if err != nil {
		log.Println("Ocurrió un error al escuchar en el puerto 8080", err)
	}
}
