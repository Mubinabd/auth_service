package managers

import (
	"database/sql"
	"fmt"
	"project/models"
)

type RecruiterManager struct {
	conn *sql.DB
}

func NewRecruiterManager(db *sql.DB) *RecruiterManager {
	return &RecruiterManager{conn: db}
}

func (rcm *RecruiterManager) CreateRecruiter(recruiter *models.RecruiterCreated) error {
	query := "INSERT INTO recruiters (name, email, phone_number, birthday, gender, company_id) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := rcm.conn.Exec(query, recruiter.Name, recruiter.Email, recruiter.PhoneNumber, recruiter.BirthDate, recruiter.Gender, recruiter.CompanyID)
	return err
}

func (rcm *RecruiterManager) GetRecruiterByID(recruiterID string) (*models.Recruiter, error) {
	query := `
		SELECT r.id, r.name, r.email, r.phone_number, r.birthday, r.gender, 
		c.id, c.name, c.location, c.workers, c.created_at, c.updated_at, c.deleted_at, 
		r.created_at, r.updated_at, r.deleted_at 
		FROM recruiters r
		JOIN companies c ON r.company_id = c.id
		WHERE r.id = $1 AND r.deleted_at = 0
	`
	row := rcm.conn.QueryRow(query, recruiterID)
	recruiter := &models.Recruiter{}
	err := row.Scan(
		&recruiter.ID, &recruiter.Name, &recruiter.Email, &recruiter.PhoneNumber,
		&recruiter.BirthDate, &recruiter.Gender,
		&recruiter.Company.CompanyID, &recruiter.Company.Name, &recruiter.Company.Location, &recruiter.Company.Workers, &recruiter.Company.CreatedAt, &recruiter.Company.UpdatedAt, &recruiter.Company.DeletedAt,
		&recruiter.CreatedAt, &recruiter.UpdatedAt, &recruiter.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return recruiter, nil
}

func (rcm *RecruiterManager) GetAllRecruiters(companyID, gender, age string) (*models.Recruiters, error) {
	query := `
		SELECT 
		r.id, r.name, r.email, r.phone_number, r.birthday, r.gender, 
		c.id, c.name, c.location, c.workers, c.created_at, c.updated_at, c.deleted_at, 
		r.created_at, r.updated_at, r.deleted_at 
		FROM recruiters r
		JOIN companies c ON r.company_id = c.id
		WHERE r.deleted_at = 0 AND c.deleted_at = 0
	`
	var args []interface{}
	paramIndex := 1

	if companyID != "" {
		query += fmt.Sprintf(" AND c.id = $%d", paramIndex)
		args = append(args, companyID)
		paramIndex++
	}
	if gender != "" {
		query += fmt.Sprintf(" AND r.gender = $%d", paramIndex)
		args = append(args, gender)
		paramIndex++
	}
	if age != "" {
		query += fmt.Sprintf(" AND EXTRACT(year FROM age(r.birthday)) = $%d", paramIndex)
		args = append(args, age)
		paramIndex++
	}

	rows, err := rcm.conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recruiters []models.Recruiter
	var count int
	for rows.Next() {
		recruiter := &models.Recruiter{}
		err := rows.Scan(
			&recruiter.ID, &recruiter.Name, &recruiter.Email, &recruiter.PhoneNumber,
			&recruiter.BirthDate, &recruiter.Gender,
			&recruiter.Company.CompanyID, &recruiter.Company.Name, &recruiter.Company.Location, &recruiter.Company.Workers, &recruiter.Company.CreatedAt, &recruiter.Company.UpdatedAt, &recruiter.Company.DeletedAt,
			&recruiter.CreatedAt, &recruiter.UpdatedAt, &recruiter.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		recruiters = append(recruiters, *recruiter)
		count++
	}
	return &models.Recruiters{Recruiters: recruiters, Count: count}, nil
}

func (rcm *RecruiterManager) UpdateRecruiter(recruiter *models.RecruiterUpdated) error {
	tempRecruiter, err := rcm.GetRecruiterByID(recruiter.ID)
	if err != nil {
		return err
	}
	if recruiter.Name == "" {
		recruiter.Name = tempRecruiter.Name
	}
	if recruiter.Email == "" {
		recruiter.Email = tempRecruiter.Email
	}
	if recruiter.PhoneNumber == "" {
		recruiter.PhoneNumber = tempRecruiter.PhoneNumber
	}
	if recruiter.BirthDate == "" {
		recruiter.BirthDate = tempRecruiter.BirthDate
	}
	if recruiter.Gender == "" {
		recruiter.Gender = tempRecruiter.Gender
	}
	query := "UPDATE recruiters SET name = $1, email = $2, phone_number = $3, birthday = $4, gender = $5, updated_at = now() WHERE id = $6"
	_, err = rcm.conn.Exec(query, recruiter.Name, recruiter.Email, recruiter.PhoneNumber, recruiter.BirthDate, recruiter.Gender, recruiter.ID)
	return err
}

func (rcm *RecruiterManager) DeleteRecruiter(recruiterID string) error {
	query := "UPDATE recruiters SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1"
	_, err := rcm.conn.Exec(query, recruiterID)
	return err
}
