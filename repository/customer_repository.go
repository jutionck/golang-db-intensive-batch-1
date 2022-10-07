package repository

import (
	"database/sql"

	"github.com/golang-db-intensive/model"
)

type CustomerRepository interface {
	Insert(customer *model.Customer) error
	Update(customer *model.Customer) error
	Delete(id int) error
	GetAll(page int, totalRows int) ([]model.Customer, error)
	GetById(id int) (model.Customer, error)
}

type customerRepository struct {
	db *sql.DB
}

func (c *customerRepository) Insert(customer *model.Customer) error {
	_, err := c.db.Exec("insert into customer(name,address,job) values ($1,$2,$3)", customer.Name, customer.Address, customer.Job)
	if err != nil {
		return err
	}

	return nil
}

func (c *customerRepository) Update(customer *model.Customer) error {
	_, err := c.db.Exec("update customer set name=$1, address=$2, job=$3 where id=$4", customer.Name, customer.Address, customer.Job, customer.Id)
	if err != nil {
		return err
	}

	return nil
}
func (c *customerRepository) Delete(id int) error {
	_, err := c.db.Exec("delete from customer where id=$1", id)
	if err != nil {
		return err
	}

	return nil
}
func (c *customerRepository) GetAll(page int, totalRows int) ([]model.Customer, error) {
	// pagination
	limit := totalRows
	offset := limit * (page - 1)
	rows, err := c.db.Query("select id,name,address,job from customer limit $1 offset $2", limit, offset)
	if err != nil {
		return nil, err
	}

	var customers []model.Customer
	for rows.Next() {
		var customer model.Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Address, &customer.Job)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (c *customerRepository) GetById(id int) (model.Customer, error) {
	var customer model.Customer

	// if err := c.db.QueryRow("select id,name,address,job from customer where id=$1", id).Scan(&customer.Id, &customer.Name, &customer.Address, &customer.Job); err != nil {
	// 	return model.Customer{}, err
	// }

	err := c.db.QueryRow("select id,name,address,job from customer where id=$1", id).Scan(&customer.Id, &customer.Name, &customer.Address, &customer.Job)
	if err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	cstRepo := new(customerRepository)
	cstRepo.db = db
	return cstRepo
}
