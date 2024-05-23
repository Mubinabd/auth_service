package managers

import (
	"database/sql"
	"fmt"
	"project/models"
)

type UserManager struct {
	conn *sql.DB
}

func NewUserManager(db *sql.DB) *UserManager {
	return &UserManager{conn: db}
}

func (um *UserManager) CreateUser(user *models.UserCreated) error {
	query := "INSERT INTO users (name, email, phone_number, birthday, gender) VALUES ($1, $2, $3, $4, $5)"
	_, err := um.conn.Exec(query, user.Name, user.Email, user.PhoneNumber, user.BirthDate, user.Gender)
	return err
}

func (um *UserManager) GetUserByID(userID string) (*models.User, error) {
	query := `
		SELECT id, name, email, phone_number, birthday, gender, created_at, updated_at, deleted_at FROM users
		WHERE id = $1 AND deleted_at = 0
		`
	row := um.conn.QueryRow(query, userID)
	user := &models.User{}
	err := row.Scan(
		&user.UserID, &user.Name,
		&user.Email, &user.PhoneNumber,
		&user.BirthDate, &user.Gender,
		&user.CreatedAt, &user.UpdatedAt,
		&user.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (um *UserManager) GetAllUsers(from, to, gender string) (*models.Users, error) {
	query := `
		SELECT id, name, email, phone_number, birthday, gender, created_at, updated_at, deleted_at FROM users 
		WHERE deleted_at = 0 
		`
	var args []interface{}
	paramIndex := 1

	if from != "" {
		query += fmt.Sprintf(" AND EXTRACT(year FROM age(birthday))	>= $%d", paramIndex)
		args = append(args, from)
		paramIndex++
	}
	if to != "" {
		query += fmt.Sprintf(" AND EXTRACT(year FROM age(birthday)) <= $%d", paramIndex)
		args = append(args, to)
		paramIndex++
	}
	if gender != "" {
		query += fmt.Sprintf(" AND gender = $%d", paramIndex)
		args = append(args, gender)
		paramIndex++
	}
	
	rows, err := um.conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	var count int
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(
			&user.UserID, &user.Name,
			&user.Email, &user.PhoneNumber,
			&user.BirthDate, &user.Gender,
			&user.CreatedAt, &user.UpdatedAt,
			&user.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
		count++
	}
	return &models.Users{Users: users, Count: count}, nil
}

func (um *UserManager) UpdateUser(user *models.UserUpdated) error {
	tempUser, err := um.GetUserByID(user.UserID)
	if err != nil {
		return err
	}

	if user.Name == "" {
		user.Name = tempUser.Name
	}
	if user.Email == "" {
		user.Email = tempUser.Email
	}
	if user.PhoneNumber == "" {
		user.PhoneNumber = tempUser.PhoneNumber
	}
	if user.BirthDate == "" {
		user.BirthDate = tempUser.BirthDate
	}
	if user.Gender == "" {
		user.Gender = tempUser.Gender
	}

	query := `
		UPDATE users SET name = $1, email = $2, phone_number = $3, birthday = $4, gender = $5, updated_at = NOW()
		WHERE id = $6 AND deleted_at = 0
	`
	_, err = um.conn.Exec(query, user.Name, user.Email, user.PhoneNumber, user.BirthDate, user.Gender, user.UserID)
	return err
}

func (um *UserManager) DeleteUser(userID string) error {
	tx, err := um.conn.Begin()
	if err != nil {
		return err
	}
	queries := []string{
		"UPDATE users SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1 AND deleted_at = 0",
		"UPDATE resumes SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE user_id = $1 AND deleted_at = 0",
		"UPDATE interviews SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE user_id = $1 AND deleted_at = 0",
	}

	for _, query := range queries {
		if _, err := tx.Exec(query, userID); err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}
