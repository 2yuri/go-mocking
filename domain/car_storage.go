package domain

type CarStorage interface {
	SaveCar(car *Car) error
	GetCar() ([]Car, error)
}
