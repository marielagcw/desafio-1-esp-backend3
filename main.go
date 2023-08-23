package main

import (
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

const (
	filename = "./tickets.csv"
)

func main() {
	total, err := tickets.GetTotalTickets("Brazil")

	Storage := tickets.Storage{
		Tickets: readFile(filename),
}


// readFile lee el archivo de tickets y devuelve un slice de tickets
func readFile(filename string) []tickets.Ticket {
	file, err := os.readFile(filename)

	iferr != nil {
		panic(err)
	}

	data := strings.Split(string(file), "\n")

	var result []tickets.Ticket

	for i:0; i < len(data); i++ {

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
