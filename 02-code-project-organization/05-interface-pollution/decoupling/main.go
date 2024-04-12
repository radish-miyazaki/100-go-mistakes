package main

//type CustomerSercvice struct {
//	store Store
//}

type customerStorer interface {
	StoreCustomer(Customer) error
}

type CustomerService struct {
	store customerStorer
}

func (cs CustomerService) CreateNewCustomer(id string) error {
	customer := Customer{id: id}
	return cs.store.StoreCustomer(customer)
}

type Store struct{}

type Customer struct {
	id string
}

func (s Store) StoreCustomer(customer Customer) error {
	return nil
}
