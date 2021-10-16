package domain

type Car struct {
	 model string
	 plate string
}

func (c Car) Model() string {
	return c.model
}

func (c Car) Plate() string {
	return c.plate
}

func NewCar(model string, plate string) *Car {
	return &Car{model: model, plate: plate}
}