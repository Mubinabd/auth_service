package managers

import (
	"database/sql"
	"fmt"
	"project/models"
)

type InterviewManager struct {
	conn *sql.DB
}

func NewInterviewManager(db *sql.DB) *InterviewManager {
	return &InterviewManager{conn: db}
}

func (im *InterviewManager) CreateInterview(interview *models.InterviewCreated) error {
	var age int
	err := im.conn.QueryRow("SELECT EXTRACT(YEAR FROM AGE(birthday)) FROM users WHERE id = $1", interview.UserID).Scan(&age)
	if err != nil {
		return err
	}
	if age < 18 {
		return fmt.Errorf("user is under 18 years old")
	}
	query1 := "SELECT position FROM resumes WHERE user_id = $1"
	var rPosition string
	err = im.conn.QueryRow(query1, interview.UserID).Scan(&rPosition)
	if err != nil {
		return err
	}
	var vPosition string
	query2 := "SELECT position from vacancies where id = $1"
	err = im.conn.QueryRow(query2, interview.VacancyID).Scan(&vPosition)
	if err != nil {
		return err
	}
	if vPosition != rPosition {
		return fmt.Errorf("user has different position")
	}

	query := "INSERT INTO interviews (user_id, vacancy_id, recruiter_id, interview_date) VALUES ($1, $2, $3, $4)"
	_, err = im.conn.Exec(query, interview.UserID, interview.VacancyID, interview.RecruiterID, interview.Date)
	if err != nil {
		return err
	}
	return nil
}

func (im *InterviewManager) GetInterviewByID(interviewID string) (*models.Interview, error) {
	query := `
	SELECT i.id, 
		u.id, u.name, u.email, u.phone_number, u.birthday, u.gender, u.created_at, u.updated_at, u.deleted_at,
		v.id, v.name, v.position, v.min_exp, 
			c.id, c.name, c.location, c.workers, c.created_at, c.updated_at, c.deleted_at,
		v.description, v.created_at, v.updated_at, v.deleted_at,
		r.id, r.name, r.email, r.phone_number, r.birthday, r.gender,
			c2.id, c2.name, c2.location, c2.workers, c2.created_at, c2.updated_at, c2.deleted_at,
		r.created_at, r.updated_at, r.deleted_at,
	i.interview_date 
	FROM interviews i
	JOIN users u ON i.user_id = u.id
	JOIN vacancies v ON i.vacancy_id = v.id
	JOIN companies c ON v.company_id = c.id
	JOIN recruiters r ON i.recruiter_id = r.id
	JOIN companies c2 ON r.company_id = c2.id
	WHERE i.id = $1 AND i.deleted_at = 0 and u.deleted_at = 0 and v.deleted_at = 0 and c.deleted_at = 0 and r.deleted_at = 0 and c2.deleted_at = 0
	`
	row := im.conn.QueryRow(query, interviewID)
	interview := &models.Interview{}
	err := row.Scan(
		&interview.InterviewID,
		&interview.User.UserID, &interview.User.Name, &interview.User.Email, &interview.User.PhoneNumber, &interview.User.BirthDate, &interview.User.Gender, &interview.User.CreatedAt, &interview.User.UpdatedAt, &interview.User.DeletedAt,
		&interview.Vacancy.VacancyID, &interview.Vacancy.Name, &interview.Vacancy.Position, &interview.Vacancy.Experience,
		&interview.Vacancy.Company.CompanyID, &interview.Vacancy.Company.Name, &interview.Vacancy.Company.Location, &interview.Vacancy.Company.Workers, &interview.Vacancy.Company.CreatedAt, &interview.Vacancy.Company.UpdatedAt, &interview.Vacancy.Company.DeletedAt,
		&interview.Vacancy.Description, &interview.Vacancy.CreatedAt, &interview.Vacancy.UpdatedAt, &interview.Vacancy.DeletedAt,
		&interview.Recruiter.ID, &interview.Recruiter.Name, &interview.Recruiter.Email, &interview.Recruiter.PhoneNumber, &interview.Recruiter.BirthDate, &interview.Recruiter.Gender,
		&interview.Recruiter.Company.CompanyID, &interview.Recruiter.Company.Name, &interview.Recruiter.Company.Location, &interview.Recruiter.Company.Workers, &interview.Recruiter.Company.CreatedAt, &interview.Recruiter.Company.UpdatedAt, &interview.Recruiter.Company.DeletedAt,
		&interview.Recruiter.CreatedAt, &interview.Recruiter.UpdatedAt, &interview.Recruiter.DeletedAt,
		&interview.Date,
	)
	if err != nil {
		return nil, err
	}
	return interview, nil
}

func (im *InterviewManager) GetAllInterviews(companyID string) (*models.Interviews, error) {
	query := `
	SELECT 
		i.id,
		u.id, u.name, u.email, u.phone_number, u.birthday, u.gender, u.created_at, u.updated_at, u.deleted_at,
		v.id, v.name, v.position, v.min_exp,
			c.id, c.name, c.location, c.workers, c.created_at, c.updated_at, c.deleted_at,
		v.description, v.created_at, v.updated_at, v.deleted_at,
		r.id, r.name, r.email, r.phone_number, r.birthday, r.gender, 
			c2.id, c2.name, c2.location, c2.workers, c2.created_at, c2.updated_at, c2.deleted_at,
		r.created_at, r.updated_at, r.deleted_at,
		i.interview_date, i.created_at, i.updated_at, i.deleted_at
	FROM interviews i 
	JOIN users u on u.id = i.user_id
	JOIN vacancies v on v.id = i.vacancy_id
	JOIN companies c on c.id = v.company_id
	JOIN recruiters r on r.id = i.recruiter_id
	JOIN companies c2 on c2.id = r.company_id
	WHERE i.deleted_at = 0 AND u.deleted_at = 0 AND v.deleted_at = 0 AND c.deleted_at = 0 AND r.deleted_at = 0 AND c2.deleted_at = 0
	`
	var args []interface{}
	paramIndex := 1
	if companyID != "" {
		query += " AND c.id = $1"
		args = append(args, companyID)
		paramIndex++
	}

	rows, err := im.conn.Query(query, args...)
	
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var interviews []models.Interview
	var count int
	for rows.Next() {
		interview := &models.Interview{}
		err := rows.Scan(
			&interview.InterviewID,
			&interview.User.UserID, &interview.User.Name, &interview.User.Email, &interview.User.PhoneNumber, &interview.User.BirthDate, &interview.User.Gender, &interview.User.CreatedAt, &interview.User.UpdatedAt, &interview.User.DeletedAt,
			&interview.Vacancy.VacancyID, &interview.Vacancy.Name, &interview.Vacancy.Position, &interview.Vacancy.Experience,
			&interview.Vacancy.Company.CompanyID, &interview.Vacancy.Company.Name, &interview.Vacancy.Company.Location, &interview.Vacancy.Company.Workers, &interview.Vacancy.Company.CreatedAt, &interview.Vacancy.Company.UpdatedAt, &interview.Vacancy.Company.DeletedAt,
			&interview.Vacancy.Description, &interview.Vacancy.CreatedAt, &interview.Vacancy.UpdatedAt, &interview.Vacancy.DeletedAt,
			&interview.Recruiter.ID, &interview.Recruiter.Name, &interview.Recruiter.Email, &interview.Recruiter.PhoneNumber, &interview.Recruiter.BirthDate, &interview.Recruiter.Gender,
			&interview.Recruiter.Company.CompanyID, &interview.Recruiter.Company.Name, &interview.Recruiter.Company.Location, &interview.Recruiter.Company.Workers, &interview.Recruiter.Company.CreatedAt, &interview.Recruiter.Company.UpdatedAt, &interview.Recruiter.Company.DeletedAt,
			&interview.Recruiter.CreatedAt, &interview.Recruiter.UpdatedAt, &interview.Recruiter.DeletedAt,
			&interview.Date, &interview.CreatedAt, &interview.UpdatedAt, &interview.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		interviews = append(interviews, *interview)
		count++
	}
	return &models.Interviews{Interviews: interviews, Count: count}, nil
}

func (im *InterviewManager) GetInterviewsBy(id string, userID, recruiterID bool) (*models.Interviews, error) {
	fmt.Println(id, userID, recruiterID)
	var query string
	if userID {
		query = `
		SELECT 
			i.id, 
			u.id, u.name, u.email, u.phone_number, u.birthday, u.gender, u.created_at, u.updated_at, u.deleted_at, 
			v.id, v.name, v.position, v.min_exp,
			c.id, c.name, c.location, c.workers, c.created_at, c.updated_at, c.deleted_at, 
			v.description, v.created_at, v.updated_at, v.deleted_at, 
			rc.id, rc.name, rc.email, rc.phone_number, rc.birthday, rc.gender, 
			c2.id, c2.name, c2.location, c2.workers, c2.created_at, c2.updated_at, c2.deleted_at, 
			rc.created_at, rc.updated_at, rc.deleted_at, 
			i.interview_date, i.created_at, i.updated_at, i.deleted_at
		FROM interviews i 
		JOIN users u ON u.id = i.user_id 
		JOIN vacancies v ON v.id = i.vacancy_id 
		JOIN companies c ON c.id = v.company_id 
		JOIN recruiters rc ON rc.id = i.recruiter_id 
		JOIN companies c2 ON c2.id = v.company_id 
		WHERE i.user_id = $1 AND i.deleted_at = 0 AND u.deleted_at = 0 AND v.deleted_at = 0 AND rc.deleted_at = 0;
	
		`
	} else if recruiterID {
		query = `
		SELECT 
			i.id, 
			u.id, u.name, u.email, u.phone_number, u.birthday, u.gender, u.created_at, u.updated_at, u.deleted_at, 
			v.id, v.name, v.position, v.min_exp,
			c.id, c.name, c.location, c.workers, c.created_at, c.updated_at, c.deleted_at, 
			v.description, v.created_at, v.updated_at, v.deleted_at, 
			rc.id, rc.name, rc.email, rc.phone_number, rc.birthday, rc.gender, 
			c2.id, c2.name, c2.location, c2.workers, c2.created_at, c2.updated_at, c2.deleted_at, 
			rc.created_at, rc.updated_at, rc.deleted_at, 
			i.interview_date, i.created_at, i.updated_at, i.deleted_at
		FROM interviews i 
		JOIN users u ON u.id = i.user_id 
		JOIN vacancies v ON v.id = i.vacancy_id 
		JOIN companies c ON c.id = v.company_id 
		JOIN recruiters rc ON rc.id = i.recruiter_id 
		JOIN companies c2 ON c2.id = v.company_id 
		WHERE i.recruiter_id = $1 AND i.deleted_at = 0 AND u.deleted_at = 0 AND v.deleted_at = 0 AND rc.deleted_at = 0;

		`
	}
	rows, err := im.conn.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var interviews []models.Interview
	var count int
	for rows.Next() {
		interview := &models.Interview{}
		err := rows.Scan(
			&interview.InterviewID,
			&interview.User.UserID, &interview.User.Name, &interview.User.Email, &interview.User.PhoneNumber, &interview.User.BirthDate, &interview.User.Gender, &interview.User.CreatedAt, &interview.User.UpdatedAt, &interview.User.DeletedAt,
			&interview.Vacancy.VacancyID, &interview.Vacancy.Name, &interview.Vacancy.Position, &interview.Vacancy.Experience,
			&interview.Vacancy.Company.CompanyID, &interview.Vacancy.Company.Name, &interview.Vacancy.Company.Location, &interview.Vacancy.Company.Workers, &interview.Vacancy.Company.CreatedAt, &interview.Vacancy.Company.UpdatedAt, &interview.Vacancy.Company.DeletedAt,
			&interview.Vacancy.Description, &interview.Vacancy.CreatedAt, &interview.Vacancy.UpdatedAt, &interview.Vacancy.DeletedAt,
			&interview.Recruiter.ID, &interview.Recruiter.Name, &interview.Recruiter.Email, &interview.Recruiter.PhoneNumber, &interview.Recruiter.BirthDate, &interview.Recruiter.Gender,
			&interview.Recruiter.Company.CompanyID, &interview.Recruiter.Company.Name, &interview.Recruiter.Company.Location, &interview.Recruiter.Company.Workers, &interview.Recruiter.Company.CreatedAt, &interview.Recruiter.Company.UpdatedAt, &interview.Recruiter.Company.DeletedAt,
			&interview.Recruiter.CreatedAt, &interview.Recruiter.UpdatedAt, &interview.Recruiter.DeletedAt,
			&interview.Date, &interview.CreatedAt, &interview.UpdatedAt, &interview.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		interviews = append(interviews, *interview)
		count++
	}
	return &models.Interviews{Interviews: interviews, Count: count}, nil
}

func (im *InterviewManager) UpdateInterview(interview *models.InterviewUpdated) error {
	tempInterview, err := im.GetInterviewByID(interview.InterviewID)
	if err != nil {
		return err
	}

	if interview.UserID == "" {
		interview.UserID = tempInterview.User.UserID
	}
	if interview.VacancyID == "" {
		interview.VacancyID = tempInterview.Vacancy.VacancyID
	}
	if interview.RecruiterID == "" {
		interview.RecruiterID = tempInterview.Recruiter.ID
	}
	if interview.Date == "" {
		interview.Date = tempInterview.Date
	}

	query := "UPDATE interviews SET user_id = $1, vacancy_id = $2, recruiter_id = $3, interview_date = $4 WHERE id = $5"
	_, err = im.conn.Exec(query, interview.UserID, interview.VacancyID, interview.RecruiterID, interview.Date, interview.InterviewID)
	return err
}

func (im *InterviewManager) DeleteInterview(interviewID string) error {
	query := "UPDATE interviews SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1"
	_, err := im.conn.Exec(query, interviewID)
	return err
}