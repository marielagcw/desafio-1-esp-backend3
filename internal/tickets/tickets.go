package tickets

import (
	"fmt"
	"time"
)

// Ticket representa un pasaje de una aerolínea
type Ticket struct {
	Id          string
	Nombre      string
	Email       string
	PaisDestino string
	HoraVuelo   string
	Precio      string
}

type Storage struct {
	Tickets []Ticket
}

func (s *Storage) PrintInfo() {
	fmt.Printf("%v+", s.Tickets)
}

// Requerimiento 1: GetTotalTickets calcula cuántas personas viajan a un país determinado.
func (s *Storage) GetTotalTickets(destination string) (int, error) {
	count := 0
	for _, ticket := range s.Tickets {
		if ticket.PaisDestino == destination {
			count++
		}
	}
	return count, nil
}

// Requerimiento 2: GetCountByPeriod calcula cuántas personas viajan en un período de tiempo específico.
func (s *Storage) GetCountByPeriod(period string) (int, error) {
	count := 0
	for _, ticket := range s.Tickets {
		if s.isInPeriod(ticket.HoraVuelo, period) {
			count++
		}
	}
	return count, nil
}

// isInPeriod verifica si el tiempo dado está en el período especificado.
func (s *Storage) isInPeriod(timeStr string, period string) bool {
	t, err := time.Parse("15:04", timeStr)
	if err != nil {
		return false
	}

	switch period {
	case "madrugada":
		return t.Hour() >= 0 && t.Hour() < 6
	case "mañana":
		return t.Hour() >= 7 && t.Hour() < 12
	case "tarde":
		return t.Hour() >= 13 && t.Hour() < 19
	case "noche":
		return t.Hour() >= 20 && t.Hour() <= 23
	default:
		return false
	}
}

// Requerimiento 3: Calcular el porcentaje de personas que viajan a un país determinado en un día.
func (s *Storage) PercentageDestination(destination string, total int) (float64, error) {
	count := 0
	for _, ticket := range s.Tickets {
		if ticket.PaisDestino == destination {
			count++
		}
	}

	if(total == 0){
		return 0, nil
	}

	percentage := float64(count) / float64(total) * 100.0
	return percentage, nil
}
