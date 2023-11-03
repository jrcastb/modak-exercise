package service

import (
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

const (
	Status    = "status"
	News      = "news"
	Marketing = "marketing"
)

var (
	Types = map[string][]interface{}{
		Status:    {2, 3 * time.Second},
		News:      {2, 10 * time.Second},
		Marketing: {3, 10 * time.Second},
	}
)

type RateLimiter struct {
	NotificationType string
	Limiter          *rate.Limiter
}

type Message interface {
	SendNotification(notificationType string)
}

func (r RateLimiter) SendNotification(notificationType string) {
	value := fmt.Sprintf("Notificación de tipo %s enviada.", notificationType)
	fmt.Printf("%v %v\n", time.Now().Format("15:04:05"), value)
}

func NewNotificationService() map[string]*RateLimiter {
	limiters := make(map[string]*RateLimiter)

	limiters[Status] = &RateLimiter{
		NotificationType: Status,
		Limiter:          newLimiter(Types[Status][0].(int), Types[Status][1].(time.Duration)), //2 cada segundo
	}
	limiters[News] = &RateLimiter{
		NotificationType: News,
		Limiter:          newLimiter(Types[News][0].(int), Types[News][1].(time.Duration)), //2 cada 10 segundos
	}
	limiters[Marketing] = &RateLimiter{
		NotificationType: Marketing,
		Limiter:          newLimiter(Types[Marketing][0].(int), Types[Marketing][1].(time.Duration)), //3 cada 5 segundos
	}
	return limiters
}

func Per(eventCount int, duration time.Duration) rate.Limit {
	return rate.Every(duration / time.Duration(eventCount))
}

func newLimiter(eventCount int, duration time.Duration) *rate.Limiter {
	return rate.NewLimiter(Per(eventCount, duration), 1)
}
