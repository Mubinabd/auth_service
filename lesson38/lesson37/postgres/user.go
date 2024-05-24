package postgres

import (
	"context"
	"errors"
	"github.com/husanmusa/NT_Golang_10/lesson38/lesson37/models"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type UserDB struct {
	Db *pgx.Conn
}

func NewUser(db *pgx.Conn) *UserDB {
	return &UserDB{db}
}

func (Userdb *UserDB) Create(ctx context.Context, std *models.User) error {
	query := `
		INSERT INTO 
			users (
				name, 
				lastname, 
				phone, 
				email, 
				age, 
				role) 
		VALUES (
				$1, 
				$2, 
				$3, 
				$4, 
				$5, 
				$6)`
	_, err := Userdb.Db.Exec(ctx, query, std.Name, std.Lastname, std.Phone, std.Email, std.Age, std.Role)
	if err != nil {
		return err
	}
	return nil
}

func (Userdb *UserDB) Update(ctx context.Context, std *models.User) error {
	query := `UPDATE users SET name = $1, lastname = $2, phone = $3, email = $4, age = $5, role = $6, updated_at = $7 WHERE id = $8`
	_, err := Userdb.Db.Exec(ctx, query, std.Name, std.Lastname, std.Phone, std.Email, std.Age, std.Role, time.Now(), std.ID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return errors.New("no rows found")
		}
		return err
	}
	return nil
}

func (Userdb *UserDB) Delete(ctx context.Context, id *string) error {
	query := `UPDATE users SET deleted_at = $1 WHERE id = $2`
	_, err := Userdb.Db.Exec(ctx, query, time.Now().Unix(), id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return errors.New("no rows found")
		}
		return err
	}
	return nil
}

func (Userdb *UserDB) GetById(ctx context.Context, id *string) (*models.User, error) {
	var user models.User
	query := `
		SELECT
			id,
			name, 
			lastname, 
			phone, 
			email, 
			age, 
			role,
			created_at,
			updated_at,
			deleted_at
		FROM 
			users 
		WHERE 
			id = $1 AND deleted_at = 0`
	err := Userdb.Db.QueryRow(ctx, query, uuid.MustParse(*id)).Scan(
		&user.ID,
		&user.Name,
		&user.Lastname,
		&user.Phone,
		&user.Email,
		&user.Age,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("no rows found")
		}
		return nil, err
	}
	return &user, nil
}

func (Userdb *UserDB) GetByRole(ctx context.Context, role *string) (*models.User, error) {
	var user models.User
	query := `
		SELECT
			id,
			name, 
			lastname, 
			phone, 
			email, 
			age, 
			role,
			created_at,
			updated_at,
			deleted_at
		FROM 
			users 
		WHERE 
			role = $1 AND deleted_at = 0`
	err := Userdb.Db.QueryRow(ctx, query, models.RoleType(*role)).Scan(
		&user.ID,
		&user.Name,
		&user.Lastname,
		&user.Phone,
		&user.Email,
		&user.Age,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("no rows found")
		}
		return nil, err
	}
	return &user, nil
}

func (Userdb *UserDB) GetAll(ctx context.Context) (*models.AllUsers, error) {
	var users []models.User
	var count uint
	query := `SELECT COUNT(1) FROM users WHERE deleted_at = 0`
	err := Userdb.Db.QueryRow(ctx, query).Scan(&count)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("no rows found")
		}
		return nil, err
	}
	query = `SELECT * FROM users WHERE deleted_at = 0`
	rows, err := Userdb.Db.Query(ctx, query)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("no rows found")
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Lastname,
			&user.Phone,
			&user.Email,
			&user.Age,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return &models.AllUsers{Count: count, Users: users}, nil
}
