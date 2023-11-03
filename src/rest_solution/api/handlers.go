package api

import (
	"modak-exercise/src/rest_solution/api/notification"
	"net/http"
)

type MessageHandler struct {
	service *notification.RateLimiter
}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{
		service: notification.NewMessageService(),
	}
}

func SendMessageNotification(writter http.ResponseWriter, request *http.Request) {
	writter.Header().Set("Content-Type", "application/json")

	notificationType := request.URL.Query().Get("type")
	notifications := NewMessageHandler().service.NewNotificationService()

	if limiter, ok := notifications[notificationType]; ok {
		if limiter.Limiter.Allow() {
			limiter.SendMessage(notificationType)
		} else {
			writter.WriteHeader(http.StatusTooManyRequests)
		}
	} else {
		writter.WriteHeader(http.StatusBadRequest)
	}
}
