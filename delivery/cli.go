package delivery

import (
	"fmt"

	"github.com/golang-db-intensive/config"
	"github.com/golang-db-intensive/model"
	"github.com/golang-db-intensive/repository"
	"github.com/golang-db-intensive/usecase"
)

func Run() {
	config := config.NewConfig()
	db := config.DbConn()
	cstRepo := repository.NewCustomerRepository(db)
	cstUseCase := usecase.NewCustomerUseCase(cstRepo)

	// Select
	// page := 1
	// totalRows := 5
	// customers, err := cstUseCase.GetAll(page, totalRows)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, customer := range customers {
	// 	fmt.Println(customer)
	// }

	// select id,name,address,job from customer where id = 2
	// customer, err := cstRepo.GetById(1)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(customer)

	// Insert
	// addCustomer01 := model.Customer{Name: "Eggy Fachrurrozi", Address: "Bandung", Job: "Backend"}
	// cstRepo.Insert(&addCustomer01)
	// addCustomer02 := model.Customer{Name: "Ahmad Hilmy Muflih", Address: "Sutabaya", Job: "Backend"}
	// cstRepo.Insert(&addCustomer02)
	// addCustomer03 := model.Customer{Name: "Muhammad Kevin", Address: "Jakarta", Job: "Android Developer"}
	// cstRepo.Insert(&addCustomer03)
	// addCustomer04 := model.Customer{Name: "Teguh Pambudi", Address: "Cianjur", Job: "Programmer"}
	// cstRepo.Insert(&addCustomer04)
	// addCustomer05 := model.Customer{Name: "Theofilus Handoyo", Address: "Malang", Job: "Programmer"}
	// cstRepo.Insert(&addCustomer05)
	// addCustomer06 := model.Customer{Name: "Yafi Abyan", Address: "Jakarta", Job: "Programmer"}
	// cstRepo.Insert(&addCustomer06)
	// addCustomer07 := model.Customer{Name: "Suci Mulan", Address: "Bandung", Job: "Secretary"}
	// cstRepo.Insert(&addCustomer07)
	// addCustomer08 := model.Customer{Name: "Anisa Rahma", Address: "Lampung", Job: "Accountan"}
	// cstRepo.Insert(&addCustomer08)

	addCustomer09 := model.Customer{Name: "Edo", Address: "Lampung", Job: "Accountan"}
	err := cstUseCase.Insert(&addCustomer09)
	fmt.Println(err)

	// Update
	// updateCustomer01 := model.Customer{Id: 3, Name: "Jution Candra", Address: "Jakarta", Job: "Programmer"}
	// cstRepo.Update(&updateCustomer01)

	// Delete
	// cstRepo.Delete(2)
}
