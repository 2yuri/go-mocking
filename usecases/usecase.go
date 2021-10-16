package usecases

import "mock-testing/domain"

type CarService struct {
	i domain.CarStorage
}

func NewCarService(i domain.CarStorage) *CarService {
	return &CarService{i: i}
}

func (s *CarService) GetAllCars() ([]domain.Car, error) {
	return s.i.GetCar()
}

func (s *CarService) Save(plate, model string) error {
	return s.i.SaveCar(domain.NewCar(plate, model))
}
