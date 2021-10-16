package usecases

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"mock-testing/domain"
	generated "mock-testing/infra/mock"
	"testing"
)

func TestSaveMethod(t *testing.T) {
	controller := gomock.NewController(t)
	mockCar := generated.NewMockCarStorage(controller)

	car := domain.NewCar("test", "yuri")
	svc := NewCarService(mockCar)

	t.Run("should return an error on save", func(t *testing.T) {
		mockCar.EXPECT().SaveCar(car).DoAndReturn(func(car *domain.Car) error {
			if car.Model() == "test" {
				return errors.New("erro qualquer")
			}

			return nil
		})

		err := svc.Save("test", "yuri")
		assert.Error(t, err)
	})

	t.Run("should return another error on save", func(t *testing.T) {
		mockCar.EXPECT().SaveCar(car).Return(errors.New("erro qualquer"))
		err := svc.Save("test", "yuri")
		assert.Error(t, err)
	})

	t.Run("should save correctly", func(t *testing.T) {
		mockCar.EXPECT().SaveCar(car).Return(nil)

		err := svc.Save("test", "yuri")
		assert.NoError(t, err)
	})
}

func TestExtraMethods(t *testing.T) {
	controller := gomock.NewController(t)
	mockDao := generated.NewMockCarStorage(controller)

	car := domain.NewCar("test", "yuri")
	svc := NewCarService(mockDao)


	t.Run("should return an error on get", func(t *testing.T) {
		mockDao.EXPECT().GetCar().Return(nil, errors.New("Deu erro")).MaxTimes(2).MinTimes(2)

		v, err := svc.GetAllCars()
		assert.Error(t, err)
		assert.Nil(t, v)

		v2, err2 := svc.GetAllCars()
		assert.Error(t, err2)
		assert.Nil(t, v2)
	})

	t.Run("should return no error and an array", func(t *testing.T) {
		var cars []domain.Car
		cars = append(cars, *car)
		cars = append(cars, *car)
		mockDao.EXPECT().GetCar().Return(cars, nil)
		v, err := svc.GetAllCars()
		assert.NoError(t, err)
		assert.NotEmpty(t, v)
	})
}

func TestGetMethod(t *testing.T) {
	controller := gomock.NewController(t)
	mockDao := generated.NewMockCarStorage(controller)
	car := domain.NewCar("test", "yuri")
	svc := NewCarService(mockDao)

	t.Run("should return an error on get", func(t *testing.T) {
		mockDao.EXPECT().GetCar().Return(nil, errors.New("Deu erro"))
		v, err := svc.GetAllCars()
		assert.Error(t, err)
		assert.Nil(t, v)
	})

	t.Run("should return no error and an array", func(t *testing.T) {
		var cars []domain.Car
		cars = append(cars, *car)
		cars = append(cars, *car)
		mockDao.EXPECT().GetCar().Return(cars, nil)
		v, err := svc.GetAllCars()
		assert.NoError(t, err)
		assert.NotEmpty(t, v)
	})
}
