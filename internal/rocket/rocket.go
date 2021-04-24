package rocket

type Rocket struct {
	ID   string
	Name string
	Type string
}

// Store - defines the interface we need to satisfy for our
// service to work correctly
type Store interface {
	GetRocketByID(id string) (Rocket, error)
	InsertRocket(rkt Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

// Service -
type Service struct {
	Store Store
}

// GetRocketByID - retrieves a rocket from the store by ID
func (s Service) GetRocketByID(id string) (Rocket, error) {
	rkt, err := s.Store.GetRocketByID(id)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// AddRocket - Adds a rocket to our store
func (s Service) AddRocket(rkt Rocket) (Rocket, error) {
	rkt, err := s.Store.InsertRocket(rkt)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// DeleteRocket - deletes a rocket - most likely rapid
// unscheduled disassembly
func (s Service) DeleteRocket(id string) error {
	err := s.Store.DeleteRocket(id)
	if err != nil {
		return err
	}
	return nil
}
