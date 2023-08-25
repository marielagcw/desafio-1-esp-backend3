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
	fmt.Printf("Total de tickets para %s: %d\n", destination, count)
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
	fmt.Printf("Total de tickets para %s: %d\n", period, count)
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


// ejemplo 1
// func GetTotalTickets(destination string) (int, error) {}

// ejemplo 2
// func GetMornings(time string) (int, error) {}

// ejemplo 3
// func AverageDestination(destination string, total int) (int, error) {}
