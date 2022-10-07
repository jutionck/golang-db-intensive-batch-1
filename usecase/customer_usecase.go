package usecase

import (
	"errors"

	"github.com/golang-db-intensive/model"
	"github.com/golang-db-intensive/repository"
)

type CustomerUseCase interface {
	RegisterNewCustomer(customer *model.Customer) error
	UpdateCustomer(customer *model.Customer) error
	DeleteCustomer(id int) error
	FindAllCustomer(page int, totalRows int) ([]model.Customer, error)
	FindCustomerById(id int) (model.Customer, error)
	FindCustomerByName(name string) ([]model.Customer, error)
}

type customerUseCase struct {
	customerRepository repository.CustomerRepository
}

func (c *customerUseCase) RegisterNewCustomer(customer *model.Customer) error {
	if customer.Name == "" {
		return errors.New("Customer name can't be empty")
	}
	return c.customerRepository.Insert(customer)
}

func (c *customerUseCase) UpdateCustomer(customer *model.Customer) error {
	return c.customerRepository.Update(customer)
}
func (c *customerUseCase) DeleteCustomer(id int) error {
	return c.customerRepository.Delete(id)
}

func (c *customerUseCase) FindAllCustomer(page int, totalRows int) ([]model.Customer, error) {
	return c.customerRepository.GetAll(page, totalRows)
}

func (c *customerUseCase) FindCustomerByName(name string) ([]model.Customer, error) {
	return c.customerRepository.GetByName(name)
}

func (c *customerUseCase) FindCustomerById(id int) (model.Customer, error) {
	return c.customerRepository.GetById(id)
}

func NewCustomerUseCase(customerRepository repository.CustomerRepository) CustomerUseCase {
	cstUseCase := new(customerUseCase)
	cstUseCase.customerRepository = customerRepository
	return cstUseCase
}
