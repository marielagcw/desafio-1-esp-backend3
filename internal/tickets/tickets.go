package tickets

// Ticket representa un pasaje de una aerolínea
type Ticket struct {
	Id		  	int
	Nombre	  	string
	Email	  	string
	PaisDestino string
	HoraVuelo  	string
	Precio	  	string
}

type Storage struct {
	Tickets []Ticket
}

// ejemplo 1
func GetTotalTickets(destination string) (int, error) {}

// ejemplo 2
func GetMornings(time string) (int, error) {}

// ejemplo 3
func AverageDestination(destination string, total int) (int, error) {}