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

func (s *Storage) GetTotalTickets(destination string) (int, error) {
	count := 0
	for _, ticket := range s.Tickets {
		if ticket.PaisDestino == destination {
			count++
		}
	}
	return count, nil
}




// ejemplo 1
// func GetTotalTickets(destination string) (int, error) {}

// ejemplo 2
// func GetMornings(time string) (int, error) {}

// ejemplo 3
// func AverageDestination(destination string, total int) (int, error) {}
