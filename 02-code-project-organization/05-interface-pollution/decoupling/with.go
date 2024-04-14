package main

type customerStorer interface {
	StoreCustomer(Customer) error
}

type CustomerService2 struct {
	store customerStorer
}

func (cs CustomerService2) CreateNewCustomer(id string) error {
	customer := Customer{id: id}
	return cs.store.StoreCustomer(customer)
}
