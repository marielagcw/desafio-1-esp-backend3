package main

import (
	"github.com/marielagcw/desafio-1-esp-backend3/internal/tickets"
	"log"
	"os"
	"strings"
	"fmt"
)

const (
	filename = "./tickets.csv"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	

	Storage := tickets.Storage{
		Tickets: readFile(filename),
	}

	Storage.PrintInfo()

	total, err := Storage.GetTotalTickets("Brazil")
		if err != nil {
			panic(err)
		}
		fmt.Println("Total de tickets a destino elegido:", total)
	
	totalCantidadTickets := len(Storage.Tickets)	

	percentageTravellersPerDay, err := Storage.PercentageDestination("Brazil", totalCantidadTickets)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Porcentaje de personas que viajan a este pais en un dia: %.2f%%\n", percentageTravellersPerDay)

}
	
// readFile lee el archivo de tickets y devuelve un slice de tickets
func readFile(filename string) []tickets.Ticket {
	file, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	data := strings.Split(string(file), "\n")

	var result []tickets.Ticket

	for i := 0; i < len(data); i++ {

		if len(data[i]) > 0 {
			file := strings.Split(data[i], ",")
			ticket := tickets.Ticket{
				Id: file[0],
				Nombre: file[1],
				Email: file[2],
				PaisDestino: file[3],
				HoraVuelo: file[4],
				Precio: file[5],
			}		
			result = append(result, ticket)
		}

	}

	return result
}