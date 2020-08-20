package mapstore

import (
	"github.com/bagomkarnath/go_code/assignments/assignmenttwo/domain"
	"github.com/bagomkarnath/go_code/assignments/assignmenttwo/err"
)

// MapStore to store data
type MapStore struct {
	store map[string]domain.Customer
}

// NewMapStore factory method for new instance of MapStore
func NewMapStore() *MapStore {
	myStore := make(map[string]domain.Customer)

	customer1 := domain.Customer{ID: "90", Name: "Sachin", Email: "sachin@gmail.com"}
	customer2 := domain.Customer{ID: "91", Name: "Saurav", Email: "saurav@gmail.com"}
	customer3 := domain.Customer{ID: "92", Name: "Laxman", Email: "laxman@gmail.com"}

	myStore["90"] = customer1
	myStore["91"] = customer2
	myStore["92"] = customer3

	return &MapStore{store: myStore}
}

// Create customer data
func (ms *MapStore) Create(c domain.Customer) error {
	if ms != nil {
		(*ms).store[c.ID] = c
		return nil
	}
	return err.ErrMapStoreNotInitialized
}

// Update customer data
func (ms *MapStore) Update(custid string, c domain.Customer) error {
	_, err := (*ms).GetByID(custid)
	if err != nil {
		return err
	}
	upCust := domain.Customer{ID: c.ID, Name: c.Name, Email: c.Email}
	(*ms).store[custid] = upCust
	return nil
}

// Delete customer data
func (ms *MapStore) Delete(custid string) error {
	_, err := (*ms).GetByID(custid)
	if err != nil {
		return err
	}
	delete((*ms).store, custid)
	return nil
}

// GetByID customer data
func (ms *MapStore) GetByID(custid string) (domain.Customer, error) {
	cust, ok := (*ms).store[custid]
	if ok {
		return cust, nil
	}
	return cust, err.ErrCustomerDoesNotExists

}

// GetAll customer data
func (ms *MapStore) GetAll() ([]domain.Customer, error) {
	if len((*ms).store) == 0 {
		return nil, err.ErrNoDataInMapStore
	}
	customers := make([]domain.Customer, 0, len((*ms).store))
	for _, v := range (*ms).store {
		customers = append(customers, v)
	}
	return customers, nil
}
