package domain

// Customer struct
type Customer struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CustomerStore interface
type CustomerStore interface {
	Create(Customer) error
	Update(string, Customer) error
	Delete(string) error
	GetByID(string) (Customer, error)
	GetAll() ([]Customer, error)
}
