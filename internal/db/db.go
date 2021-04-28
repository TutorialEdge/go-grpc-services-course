package db

import (
	"fmt"
	"os"

	"github.com/TutorialEdge/go-grpc-services-course/internal/rocket"

	"github.com/jmoiron/sqlx"
)

// DB
type Store struct {
	db *sqlx.DB
}

// New - returns a new Store
func New() (Store, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")

	connectString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)

	db, err := sqlx.Connect("postgres", connectString)
	if err != nil {
		return Store{}, err
	}
	return Store{
		db: db,
	}, nil
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
