package delivery

import (
	"fmt"

	"github.com/golang-db-intensive/config"
	"github.com/golang-db-intensive/model"
	"github.com/golang-db-intensive/repository"
	"github.com/golang-db-intensive/usecase"
	"github.com/golang-db-intensive/utils"
)

func Run() {
	config := config.NewConfig()
	db := config.DbConn()
	cstRepo := repository.NewCustomerRepository(db)
	cstUseCase := usecase.NewCustomerUseCase(cstRepo)

	// Select
	page := 2
	totalRows := 10
	customers, err := cstUseCase.FindAllCustomer(page, totalRows)
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
	addCustomer01 := model.Customer{
		Name:    "Steve Job",
		Address: utils.NewNullString("UK"),
		Job:     "CEO"}
	cstRepo.Insert(&addCustomer01)

	// Update
	// updateCustomer01 := model.Customer{Id: 3, Name: "Jution Candra", Address: "Jakarta", Job: "Programmer"}
	// cstRepo.Update(&updateCustomer01)

	// Delete
	// cstRepo.Delete(2)

	// Find customer by name
	// customers, err := cstUseCase.FindCustomerByName("b")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, customer := range customers {
	// 	fmt.Println(customer)
	// }
}
