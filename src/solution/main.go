package main

import (
	"fmt"
	"modak-exercise/src/solution/notification"
	"time"
)

func main() {
	/*limiters := make(map[string]*RateLimiter)
	limiters[Status] = &RateLimiter{
		NotificationType: Status,
		Limiter:          rate.NewLimiter(Per(2, 3*time.Second), 1), //2 cada segundo
	}
	limiters[News] = &RateLimiter{
		NotificationType: News,
		Limiter:          rate.NewLimiter(Per(2, 10*time.Second), 1), //2 cada 10 segundos
	}
	limiters[Marketing] = &RateLimiter{
		NotificationType: Marketing,
		Limiter:          rate.NewLimiter(Per(3, 10*time.Second), 1), //3 cada 5 segundos
	}*/
	notifications := service.NewNotificationService()
	for {
		notificationType := service.Status
		if limiter, ok := notifications[notificationType]; ok {
			if limiter.Limiter.Allow() {
				limiter.SendNotification(notificationType)
			} else {
				fmt.Printf("Solicitud de notificación de tipo %s rechazada: limite de velocidad excedido.\n", notificationType)
			}
		} else {
			fmt.Printf("Tipo de notificación no valido: %s.\n", notificationType)
		}
		time.Sleep(1 * time.Second)
	}
}
