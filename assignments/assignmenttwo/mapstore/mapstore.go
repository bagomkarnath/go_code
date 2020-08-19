package mapstore

import (
	"errors"

	"github.com/bagomkarnath/go_code/assignments/assignmenttwo/domain"
)

// MapStore to store data
type MapStore struct {
	store map[string]domain.Customer
}

// NewMapStore factory method for new instance of MapStore
func NewMapStore() *MapStore {
	return &MapStore{store: make(map[string]domain.Customer)}
}

// Create customer data
func (ms *MapStore) Create(c domain.Customer) error {
	if ms != nil {
		(*ms).store[c.GetID()] = c
		return nil
	}
	return errors.New("MapStore not initialized")
}

// Update customer data
func (ms *MapStore) Update(custid string, c domain.Customer) error {
	_, err := (*ms).GetByID(custid)
	if err != nil {
		return err
	}
	var upCust domain.Customer
	upCust.SetID(c.GetID())
	upCust.SetName(c.GetName())
	upCust.SetEmail(c.GetEmail())
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
	return cust, errors.New("Customer doesn't exists")

}

// GetAll customer data
func (ms *MapStore) GetAll() ([]domain.Customer, error) {
	if len((*ms).store) == 0 {
		return nil, errors.New("No data in Map Store")
	}
	customers := make([]domain.Customer, 0, len((*ms).store))
	for _, v := range (*ms).store {
		customers = append(customers, v)
	}
	return customers, nil
}
