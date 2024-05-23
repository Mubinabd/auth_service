package managers

import (
	"database/sql"
	"fmt"
	"project/models"
	"strconv"
)

type ResumeManager struct {
	conn *sql.DB
}

func NewResumeManager(db *sql.DB) *ResumeManager {
	return &ResumeManager{conn: db}
}

func (rm *ResumeManager) CreateResume(resume *models.ResumeCreated) error {
	query := "INSERT INTO resumes (position, experience, description, user_id) VALUES ($1, $2, $3, $4)"
	_, err := rm.conn.Exec(query, resume.Position, resume.Experience, resume.Description, resume.User)
	return err
}

func (rm *ResumeManager) GetResumeByID(resumeID string) (*models.Resume, error) {
	query := `
	SELECT 
		r.id, r.position, r.experience, r.description, 
		u.id, u.name, u.email, u.phone_number, u.birthday, u.gender, u.created_at, u.updated_at, u.deleted_at 
	FROM resumes r join users u ON r.user_id = u.id WHERE r.id = $1 AND r.deleted_at = 0 and u.deleted_at = 0
	`
	row := rm.conn.QueryRow(query, resumeID)
	resume := &models.Resume{}
	err := row.Scan(
		&resume.ResumeID, &resume.Position, &resume.Experience, &resume.Description,
		&resume.User.UserID, &resume.User.Name, &resume.User.Email, &resume.User.PhoneNumber, &resume.User.BirthDate, &resume.User.Gender, &resume.User.CreatedAt, &resume.User.UpdatedAt, &resume.User.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return resume, nil
}

// Giving userID is optional, if userID is not given it will return all resumes
func (rm *ResumeManager) GetAllResumes(userID, position, min_exp string) (*models.Resumes, error) {
	query := `
	SELECT 
		r.id, r.position, r.experience, r.description, 
		u.id, u.name, u.email, u.phone_number, u.birthday, u.gender, u.created_at, u.updated_at, u.deleted_at 
	FROM resumes r JOIN users u ON r.user_id = u.id WHERE r.deleted_at = 0 AND u.deleted_at = 0`
	var args []interface{}
	paramIndex := 1
	if userID != "" {
		query += fmt.Sprintf(" AND r.user_id = $%d", paramIndex)
		args = append(args, userID)
		paramIndex++
	}
	if position != "" {
		query += fmt.Sprintf(" AND r.position = $%d", paramIndex)
		args = append(args, position)
		paramIndex++
	}
	if min_exp != "" {
		query += fmt.Sprintf(" AND r.experience >= $%d", paramIndex)
		min_exp_int, err := strconv.Atoi(min_exp)
		if err != nil {
			return nil, err
		}
		args = append(args, min_exp_int)
	}

	fmt.Println(query, args)
	rows, err := rm.conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resumes []models.Resume
	var count int
	for rows.Next() {
		resume := &models.Resume{}
		err := rows.Scan(
			&resume.ResumeID, &resume.Position,
			&resume.Experience, &resume.Description,
			&resume.User.UserID, &resume.User.Name,
			&resume.User.Email, &resume.User.PhoneNumber,
			&resume.User.BirthDate, &resume.User.Gender,
			&resume.User.CreatedAt, &resume.User.UpdatedAt,
			&resume.User.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		resumes = append(resumes, *resume)
		count++
	}

	return &models.Resumes{Resumes: resumes, Count: count}, nil
}

func (rm *ResumeManager) UpdateResume(resume *models.ResumeUpdated) error {
	tempResume, err := rm.GetResumeByID(resume.ResumeID)
	if err != nil {
		return err
	}
	if resume.Position == "" {
		tempResume.Position = resume.Position
	}
	query := "UPDATE resumes SET position = $1, experience = $2, description = $3, updated_at = NOW() WHERE id = $4"
	_, err = rm.conn.Exec(query, resume.Position, resume.Experience, resume.Description, resume.ResumeID)
	return err
}

func (rm *ResumeManager) DeleteResume(resumeID string) error {
	query := "UPDATE resumes SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1"
	_, err := rm.conn.Exec(query, resumeID)
	return err
}
