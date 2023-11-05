package main

import (
	"encoding/json"
	"fmt"
	"github.com/didip/tollbooth/v7"
	"log"
	"net/http"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func main() {

	message := Message{
		Status: "Request Failed",
		Body:   "Excedió el numero de peticiones.",
	}
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		log.Println("ocurrió un error intentando mapear", err)
	}

	tlbthLimiterType := tollbooth.NewLimiter(2, nil)
	tlbthLimiterType.SetMessageContentType("application/json")
	tlbthLimiterType.SetMessage(string(jsonMessage))

	http.Handle("/notification", tollbooth.LimitFuncHandler(tlbthLimiterType, notificationHandler))
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Ocurrió un error escuchando en el puerto :8080", err)
	}
}

func notificationHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	notificationType := request.URL.Query().Get("type")
	writer.WriteHeader(http.StatusOK)
	message := Message{
		Status: "Successful",
		Body:   fmt.Sprintf("Has enviado correctamente la notificación de tipo: %s", notificationType),
	}
	err := json.NewEncoder(writer).Encode(&message)
	if err != nil {
		return
	}
}
