package store

func (s *Store) GetContract(id string) (*Contract, error) {
	return &Contract{}, nil
}

func (s *Store) SetContract(id string, c *Contract) error {
	return nil
}

func (s *Store) GetCustomer(id string) (*Customer, error) {
	return &Customer{}, nil
}

func (s *Store) SetCustomer(id string, c *Customer) error {
	return nil
}
