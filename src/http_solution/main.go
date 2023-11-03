package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"modak-exercise/src/utils"
	"net/http"
	"time"
)

type RateLimiter struct {
	NotificationType string
	Limiter          *rate.Limiter
}

var limiters map[string]*RateLimiter

func main() {
	limiters = make(map[string]*RateLimiter)
	limiters["status"] = &RateLimiter{
		NotificationType: "status",
		Limiter:          rate.NewLimiter(rate.Every(10*time.Second), 1),
	}
	limiters["news"] = &RateLimiter{
		NotificationType: "news",
		Limiter:          rate.NewLimiter(rate.Every(30*time.Second), 1),
	}

	http.HandleFunc("/send-message", sendNotificationHandler)

	fmt.Println("Servidor escuchando en el puerto 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func sendNotificationHandler(w http.ResponseWriter, r *http.Request) {
	notificationType := r.URL.Query().Get("type")

	if limiter, ok := limiters[notificationType]; ok {
		if limiter.Limiter.Allow() {
			sendMessage(notificationType)
		} else {
			w.WriteHeader(http.StatusTooManyRequests)
			utils.LogRejection(notificationType)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		utils.LogInvalidNotificationType(notificationType)
	}
}

func sendMessage(notificationType string) {
	value := fmt.Sprintf("Notificación de tipo %s enviada.", notificationType)
	fmt.Printf("%v %v\n", time.Now().Format("15:04:05"), value)
}
