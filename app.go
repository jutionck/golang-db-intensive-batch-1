package main

import (
	"fmt"

	"github.com/golang-db-intensive/config"
	"github.com/golang-db-intensive/repository"
	_ "github.com/lib/pq"
)

func main() {

	config := config.NewConfig()
	db := config.DbConn()
	cstRepo := repository.NewCustomerRepository(db)

	// Select
	customers, err := cstRepo.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, customer := range customers {
		fmt.Println(customer)
	}

	// select id,name,address,job from customer where id = 2
	// customer, err := cstRepo.GetById(1)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(customer)

	// Insert
	// addCustomer01 := model.Customer{Name: "Fadli", Address: "Cianjur", Job: "Programmer"}
	// cstRepo.Insert(&addCustomer01)

	// Update
	// updateCustomer01 := model.Customer{Id: 3, Name: "Jution Candra", Address: "Jakarta", Job: "Programmer"}
	// cstRepo.Update(&updateCustomer01)

	// Delete
	cstRepo.Delete(2)
}
