package domain

// Customer struct
type Customer struct {
	id, name, email string
}

// CustomerStore interface
type CustomerStore interface {
	Create(Customer) error
	Update(string, Customer) error
	Delete(string) error
	GetByID(string) (Customer, error)
	GetAll() ([]Customer, error)
}

// GetID returns id of Customer
func (c *Customer) GetID() string {
	return (*c).id
}

// SetID set ID
func (c *Customer) SetID(custID string) {
	(*c).id = custID
}

// GetName returns name
func (c *Customer) GetName() string {
	return (*c).name
}

// SetName sets name
func (c *Customer) SetName(n string) {
	(*c).name = n
}

// GetEmail returns email
func (c *Customer) GetEmail() string {
	return (*c).email
}

// SetEmail sets email
func (c *Customer) SetEmail(em string) {
	(*c).email = em
}

// NewCustomer with all fields
// func (c *Customer) NewCustomer(cid, cname, cemail string) {
// 	(*c).id = cid
// 	(*c).name = cname
// 	(*c).email = cemail
// }
