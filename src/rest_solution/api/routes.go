package api

import "net/http"

func InitRoutes() {
	http.HandleFunc("/send-message", SendMessageNotification)
}
