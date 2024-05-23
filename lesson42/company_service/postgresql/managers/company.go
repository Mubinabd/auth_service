package managers

import (
	"database/sql"
	"fmt"
	"project/models"
)

type CompanyManager struct {
	conn *sql.DB
}

func NewCompanyManager(db *sql.DB) *CompanyManager {
	return &CompanyManager{conn: db}
}

func (cm *CompanyManager) CreateCompany(company *models.CompanyCreated) error {
	query := "INSERT INTO companies (name, location, workers) VALUES ($1, $2, $3)"
	_, err := cm.conn.Exec(query, company.Name, company.Location, company.Workers)
	return err
}

func (cm *CompanyManager) GetCompanyByID(companyID string) (*models.Company, error) {
	query := "SELECT id, name, location, workers FROM companies WHERE id = $1 AND deleted_at = 0"
	row := cm.conn.QueryRow(query, companyID)
	company := &models.Company{}
	err := row.Scan(
		&company.CompanyID, &company.Name, &company.Location, &company.Workers,
	)
	if err != nil {
		return nil, err
	}
	return company, nil
}

func (cm *CompanyManager) GetAllCompanies(location string) (*models.Companies, error) {
	query := "SELECT id, name, location, workers FROM companies WHERE deleted_at = 0"

	var args []interface{}
	paramIndex := 1
	if location != "" {
		query += " AND location = $" + fmt.Sprint(paramIndex)
		args = append(args, location)
		paramIndex++
	}
	
	rows, err := cm.conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companies []models.Company
	var count int
	for rows.Next() {
		company := &models.Company{}
		err := rows.Scan(
			&company.CompanyID, &company.Name,
			&company.Location, &company.Workers,
		)
		if err != nil {
			return nil, err
		}
		companies = append(companies, *company)
		count++
	}
	return &models.Companies{Companies: companies, Count: count}, nil
}

func (cm *CompanyManager) UpdateCompany(company *models.CompanyUpdated) error {
	temp, err := cm.GetCompanyByID(company.CompanyID)
	if err != nil {
		return err
	}
	if company.Name == "" {
		company.Name = temp.Name
	}
	if company.Location == "" {
		company.Location = temp.Location
	}
	if company.Workers == 0 {
		company.Workers = temp.Workers
	}

	query := "UPDATE companies SET name = $1, location = $2, workers = $3, updated_at = now() WHERE id = $4 AND deleted_at = 0"
	_, err = cm.conn.Exec(query, company.Name, company.Location, company.Workers, company.CompanyID)
	return err
}

func (cm *CompanyManager) DeleteCompany(companyID string) error {
	query := "UPDATE companies SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1 AND deleted_at = 0"
	_, err := cm.conn.Exec(query, companyID)
	return err
}
