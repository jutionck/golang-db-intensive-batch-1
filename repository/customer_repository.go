package repository

import (
	"github.com/golang-db-intensive/model"
	"github.com/golang-db-intensive/utils"
	"github.com/jmoiron/sqlx"
)

type CustomerRepository interface {
	Insert(customer *model.Customer) error
	Update(customer *model.Customer) error
	Delete(id int) error
	GetAll(page int, totalRows int) ([]model.Customer, error)
	GetById(id int) (model.Customer, error)
}

type customerRepository struct {
	db *sqlx.DB
}

func (c *customerRepository) Insert(customer *model.Customer) error {
	_, err := c.db.NamedExec(utils.INSERT_CUSTOMER, customer)
	if err != nil {
		return err
	}

	return nil
}

func (c *customerRepository) Update(customer *model.Customer) error {
	_, err := c.db.NamedExec(utils.UPDATE_CUSTOMER, customer)
	if err != nil {
		return err
	}

	return nil
}
func (c *customerRepository) Delete(id int) error {
	_, err := c.db.Exec(utils.DELETE_CUSTOMER, id)
	if err != nil {
		return err
	}

	return nil
}
func (c *customerRepository) GetAll(page int, totalRows int) ([]model.Customer, error) {
	// pagination
	limit := totalRows
	offset := limit * (page - 1)
	var customers []model.Customer
	err := c.db.Select(&customers, utils.SELECT_CUSTOMER_LIST, limit, offset)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (c *customerRepository) GetById(id int) (model.Customer, error) {
	var customer model.Customer
	err := c.db.Get(&customer, utils.SELECT_CUSTOMER_ID, id)
	if err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}

func NewCustomerRepository(db *sqlx.DB) CustomerRepository {
	cstRepo := new(customerRepository)
	cstRepo.db = db
	return cstRepo
}
