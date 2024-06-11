package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Mubinabd/auth_service/config"
	"github.com/Mubinabd/auth_service/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db    *sql.DB
	UserS storage.UserI
}

func ConnectDb(cfg config.Config) (*Storage, error) {
	con := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.PostgresHost, cfg.PostgresUser, cfg.PostgresUser, cfg.PostgresPort, cfg.PostgresDatabase)
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	var str = &Storage{Db: db}
	str.UserS = NewUserManager(db)

	return str, nil
}

func (s *Storage) User() storage.UserI {
	if s.UserS == nil {
		s.UserS = NewUserManager(s.Db)
	}

	return s.UserS
}
