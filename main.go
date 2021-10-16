package main

import (
	"log"
	"mock-testing/infra/postgres"
	"mock-testing/usecases"
)

func main() {
	service := usecases.NewCarService(&postgres.InMemoryDb{})

	err := service.Save("Teste", "Teste")
	if err != nil {
		log.Fatal("cannot save")
	}

	v, err := service.GetAllCars()
	if err != nil {
		log.Fatal("cannot get")
	}

	for _, car := range v {
		log.Println("------------")
		log.Println("Model: ", car.Model())
		log.Println("Plate: ", car.Plate())
		log.Println("------------")
	}

}
