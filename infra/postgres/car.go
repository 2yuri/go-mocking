package postgres

import (
	"mock-testing/domain"
)

//InMemoryDb é uma implementação fake, o nome de postgres é só para exemplificação
type InMemoryDb struct {
}

var db []domain.Car

func (i InMemoryDb) SaveCar(car *domain.Car) error {
	db = append(db, *car)

	return nil
}

func (i InMemoryDb) GetCar() ([]domain.Car, error) {
	return db, nil
}
