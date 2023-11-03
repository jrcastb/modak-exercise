package utils

import "fmt"

func LogRejection(nottificationType string) {
	fmt.Printf("Solicitud de notificación de tipo %s rechazada: Limite de velocidad excedido\n", nottificationType)
}
func LogInvalidNotificationType(notificationType string) {
	fmt.Printf("Tipo de notificación: %s no es válido.\n", notificationType)
}
