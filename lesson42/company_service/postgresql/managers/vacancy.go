package managers

import (
	"database/sql"
	"fmt"
	"project/models"
)

type VacancyManager struct {
	conn *sql.DB
}

func NewVacancyManager(db *sql.DB) *VacancyManager {
	return &VacancyManager{conn: db}
}

func (vm *VacancyManager) CreateVacancy(v *models.VacancyCreated) error {
	query := "INSERT INTO vacancies (name, position, min_exp, company_id, description) VALUES ($1, $2, $3, $4, $5)"
	_, err := vm.conn.Exec(query, v.Name, v.Position, v.Experience, v.CompanyID, v.Description)
	return err
}

func (vm *VacancyManager) GetVacancyByID(id string) (*models.Vacancy, error) {
	query := `
		SELECT v.id, v.name, v.position, v.min_exp, 
		c.id, c.name, c.location, c.workers, c.created_at, c.updated_at, c.deleted_at,
		v.description, v.created_at, v.updated_at, v.deleted_at 
		FROM vacancies v
		JOIN companies c ON v.company_id = c.id
		WHERE v.id = $1 AND v.deleted_at = 0 and c.deleted_at = 0
	`
	vacancy := &models.Vacancy{}
	err := vm.conn.QueryRow(query, id).Scan(
		&vacancy.VacancyID, &vacancy.Name, &vacancy.Position, &vacancy.Experience,
		&vacancy.Company.CompanyID, &vacancy.Company.Name, &vacancy.Company.Location, &vacancy.Company.Workers, &vacancy.Company.CreatedAt, &vacancy.Company.UpdatedAt, &vacancy.Company.DeletedAt,
		&vacancy.Description, &vacancy.CreatedAt, &vacancy.UpdatedAt, &vacancy.DeletedAt,
	)

	if err != nil {
		return nil, err
	}
	return vacancy, nil
}

func (vm *VacancyManager) GetAllVacancies(position, min_exp, company_id string) (*models.Vacancies, error) {
	query := `
		SELECT v.id, v.name, v.position, v.min_exp, 
		c.id, c.name, c.location, c.workers, c.created_at, c.updated_at, c.deleted_at,
		v.description, v.created_at, v.updated_at, v.deleted_at 
		FROM vacancies v
		JOIN companies c ON v.company_id = c.id
		WHERE v.deleted_at = 0 and c.deleted_at = 0
	`
	var args []interface{}
	paramIndex := 1
	if position != "" {
		query += fmt.Sprintf(" AND position = $%d", paramIndex)
		args = append(args, position)
		paramIndex++
	}
	if min_exp != "" {
		query += fmt.Sprintf(" AND min_exp >= $%d", paramIndex)
		args = append(args, min_exp)
		paramIndex++
	}
	if company_id != "" {
		query += fmt.Sprintf(" AND company_id = $%d", paramIndex)
		args = append(args, company_id)
		paramIndex++
	}
	rows, err := vm.conn.Query(query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vacancies := []models.Vacancy{}
	var count int
	for rows.Next() {
		vacancy := &models.Vacancy{}
		err := rows.Scan(
			&vacancy.VacancyID, &vacancy.Name, &vacancy.Position, &vacancy.Experience,
			&vacancy.Company.CompanyID, &vacancy.Company.Name, &vacancy.Company.Location, &vacancy.Company.Workers, &vacancy.Company.CreatedAt, &vacancy.Company.UpdatedAt, &vacancy.Company.DeletedAt,
			&vacancy.Description, &vacancy.CreatedAt, &vacancy.UpdatedAt, &vacancy.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		vacancies = append(vacancies, *vacancy)
		count++
	}
	return &models.Vacancies{Vacancies: vacancies, Count: count}, nil
}

func (vm *VacancyManager) UpdateVacancy(v *models.VacancyUpdated) error {
	tempVacancy, err := vm.GetVacancyByID(v.VacancyID)
	if err != nil {
		return err
	}
	if v.Name == "" {
		tempVacancy.Name = v.Name
	}
	if v.Position == "" {
		tempVacancy.Position = v.Position
	}

	query := "UPDATE vacancies SET name = $1, position = $2, min_exp = $3, description = $4 WHERE id = $5"
	_, err = vm.conn.Exec(query, v.Name, v.Position, v.Experience, v.Description, v.VacancyID)
	return err
}

func (vm *VacancyManager) DeleteVacancy(id string) error {
	query := "UPDATE vacancies SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1"
	_, err := vm.conn.Exec(query, id)
	return err
}