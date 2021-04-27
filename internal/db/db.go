package db

import (
	"database/sql"

	"github.com/TutorialEdge/go-grpc-services-course/internal/rocket"
)

// DB
type Store struct {
	db *sql.DB
}

// New - returns a new Store
func New() Store {
	return Store{
		db: nil,
	}
}

func (s Store) GetRocketByID(id string) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}

func (s Store) InsertRocket(rkt rocket.Rocket) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}

func (s Store) DeleteRocket(id string) error {
	return nil
}
