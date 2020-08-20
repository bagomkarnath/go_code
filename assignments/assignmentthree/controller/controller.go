package controller

import (
	"fmt"

	"github.com/bagomkarnath/go_code/assignments/assignmenttwo/domain"
)

// CustomerController as controller
type CustomerController struct {
	// Explicit dependency and declarative programming that hides dependent logic
	store domain.CustomerStore // It can be any Store including MapStore
}

// Add to add data
func (controller CustomerController) Add(cust domain.Customer) {
	err := controller.store.Create(cust)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("New Customer has been created")
}

// ShowAll displays all customers
func (controller CustomerController) ShowAll() {
	customers, err := controller.store.GetAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("All data : ", customers)
}

// Delete a customer
func (controller CustomerController) Delete(custid string) {
	err := controller.store.Delete(custid)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Customer record deleted")
}

// Update a customer
func (controller CustomerController) Update(custid string, c domain.Customer) {
	err := controller.store.Update(custid, c)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Customer record updated")
}

// Get customer by id
func (controller CustomerController) Get(custid string) {
	cust, err := controller.store.GetByID(custid)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Customer : ", cust)
}
