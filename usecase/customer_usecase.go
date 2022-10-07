package usecase

import (
	"errors"

	"github.com/golang-db-intensive/model"
	"github.com/golang-db-intensive/repository"
)

type CustomerUseCase interface {
	Insert(customer *model.Customer) error
	Update(customer *model.Customer) error
	Delete(id int) error
	GetAll(page int, totalRows int) ([]model.Customer, error)
	GetById(id int) (model.Customer, error)
}

type customerUseCase struct {
	customerRepository repository.CustomerRepository
}

func (c *customerUseCase) Insert(customer *model.Customer) error {
	if customer.Name == "" {
		return errors.New("Customer name can't be empty")
	}
	return c.customerRepository.Insert(customer)
}

func (c *customerUseCase) Update(customer *model.Customer) error {
	return c.customerRepository.Update(customer)
}
func (c *customerUseCase) Delete(id int) error {
	return c.customerRepository.Delete(id)
}
func (c *customerUseCase) GetAll(page int, totalRows int) ([]model.Customer, error) {
	return c.customerRepository.GetAll(page, totalRows)
}

func (c *customerUseCase) GetById(id int) (model.Customer, error) {
	return c.customerRepository.GetById(id)
}

func NewCustomerUseCase(customerRepository repository.CustomerRepository) CustomerUseCase {
	cstUseCase := new(customerUseCase)
	cstUseCase.customerRepository = customerRepository
	return cstUseCase
}
