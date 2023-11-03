package main

import (
	"modak-exercise/src/solution/notification"
	"modak-exercise/src/utils"
	"time"
)

func main() {
	notifications := service.NewNotificationService()
	notificationType := service.Status
	for {
		if limiter, ok := notifications[notificationType]; ok {
			if limiter.Limiter.Allow() {
				limiter.SendNotification(notificationType)
			} else {
				utils.LogRejection(notificationType)
			}
		} else {
			utils.LogInvalidNotificationType(notificationType)
		}
		time.Sleep(1 * time.Second)
	}
}
