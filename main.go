package main

import (
	"fmt"
	"github.com/marielagcw/desafio-1-esp-backend3/internal/tickets"
	"log"
	"os"
	"strings"
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

	//Storage.PrintInfo()

	// Creación de los channels para las go rutines
	canalInt := make(chan int)
	defer close(canalInt)
	canalFloat := make(chan float64)
	defer close(canalFloat)
	canalError := make(chan error)
	defer close(canalError)

	// Ejecución de las go rutines
	fmt.Println(" -------------------------------- Ejecución Requerimiento 1 -------------------------------- ")
	go func(chan int, chan error) {

		total, err := Storage.GetTotalTickets("Brazil")
		if err != nil {
			canalError <- err
			return
		}

		canalInt <- total

	}(canalInt, canalError)

	select {
	case total := <-canalInt:
		fmt.Printf("Total de tickets para el destino solicitado es: %d\n", total)
	case err := <-canalError:
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println(" -------------------------------- Ejecución Requerimiento 2 -------------------------------- ")

	go func(chan int, chan error) {

		count, err := Storage.GetCountByPeriod("madrugada")
		if err != nil {
			canalError <- err
			return
		}

		canalInt <- count

	}(canalInt, canalError)

	select {
	case count := <-canalInt:
		fmt.Printf("Total de tickets para el periodo solicitado es: %d\n", count)
	case err := <-canalError:
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println(" -------------------------------- Ejecución Requerimiento 3 -------------------------------- ")
	go func(chan float64, chan error) {

		totalCantidadTickets := len(Storage.Tickets)

		percentageTravellersPerDay, err := Storage.PercentageDestination("Brazil", totalCantidadTickets)
		if err != nil {
			canalError <- err
			return
		}

		canalFloat <- percentageTravellersPerDay

	}(canalFloat, canalError)

	select {
	case percentageTravellersPerDay := <-canalFloat:
		fmt.Printf("Porcentaje de personas que viajan a este pais en un dia: %.2f%%\n", percentageTravellersPerDay)
	case err := <-canalError:
		log.Fatal(err)
		os.Exit(1)
	}
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
				Id:          file[0],
				Nombre:      file[1],
				Email:       file[2],
				PaisDestino: file[3],
				HoraVuelo:   file[4],
				Precio:      file[5],
			}
			result = append(result, ticket)
		}

	}

	return result
}
