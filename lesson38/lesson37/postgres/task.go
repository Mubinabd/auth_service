package postgres

import (
	"context"
	"errors"
	"github.com/husanmusa/NT_Golang_10/lesson38/lesson37/models"
	"time"

	"github.com/jackc/pgx/v5"
)

type TaskDB struct {
	Db *pgx.Conn
}

func NewTask(db *pgx.Conn) *TaskDB {
	return &TaskDB{db}
}

func (taskdb *TaskDB) Create(ctx context.Context, task *models.Task) error {
	query := `INSERT INTO task (name, type, user_id, description, deadline) VALUES ($1, $2, $3, $4, $5)`
	_, err := taskdb.Db.Exec(ctx, query, task.Name, task.Type, task.UserID, task.Description, task.DeadLine)
	if err != nil {
		return err
	}
	return nil
}

func (taskdb *TaskDB) Update(ctx context.Context, task *models.Task) error {
	query := `UPDATE task SET name = $1, type = $2, user_id = $3, description = $4, deadline = $5, updated_at = $6 WHERE id = $7`
	_, err := taskdb.Db.Exec(ctx, query, task.Name, task.Type, task.UserID, task.Description, task.DeadLine, time.Now(), task.ID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return errors.New("no rows found")
		}
		return err
	}
	return nil
}

func (taskdb *TaskDB) Delete(ctx context.Context, id *string) error {
	query := `
				UPDATE
				 	task 
				SET 
					deleted_at = $1 
				WHERE 
					id = $2`
	_, err := taskdb.Db.Exec(ctx, query, time.Now().Unix(), *id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return errors.New("no rows found")
		}
		return err
	}
	return nil
}

func (taskdb *TaskDB) GetById(ctx context.Context, id *string) (*models.Task, error) {
	var task models.Task
	query := `SELECT 
					id, 
					name, 
					type, 
					user_id, 
					description, 
					deadline, 
					created_at, 
					updated_at, 
					deleted_at 
				FROM 
					task 
				WHERE 
					id = $1
				AND
					deleted_at = 0`

	err := taskdb.Db.QueryRow(ctx, query, *id).Scan(
		&task.ID,
		&task.Name,
		&task.Type,
		&task.UserID,
		&task.Description,
		&task.DeadLine,
		&task.CreatedAt,
		&task.UpdatedAt,
		&task.DeletedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("no rows found")
		}
		return nil, err
	}
	return &task, nil
}

func (taskdb *TaskDB) GetAll(ctx context.Context) (*models.AllTasks, error) {
	var tasks []models.Task
	var count uint
	query := `
	SELECT 
		COUNT(1) 
	FROM 
		task 
	WHERE 
		deleted_at = 0`
	err := taskdb.Db.QueryRow(ctx, query).Scan(&count)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("no rows found")
		}
		return nil, err
	}
	query = `SELECT 
					id, 
					name, 
					type, 
					user_id, 
					description, 
					deadline, 
					created_at, 
					updated_at, 
					deleted_at 
				FROM 
					task 
				WHERE 
					deleted_at = 0`
	rows, err := taskdb.Db.Query(ctx, query)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("no rows found")
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var task models.Task
		err = rows.Scan(
			&task.ID,
			&task.Name,
			&task.Type,
			&task.UserID,
			&task.Description,
			&task.DeadLine,
			&task.CreatedAt,
			&task.UpdatedAt,
			&task.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return &models.AllTasks{Count: count, Tasks: tasks}, nil
}

func (taskdb *TaskDB) GetByUserId(ctx context.Context, userID *string) ([]models.Task, error) {
	var tasks []models.Task

	query := `SELECT 
				id, 
				name, 
				type, 
				user_id, 
				description, 
				deadline, 
				created_at, 
				updated_at, 
				deleted_at 
			FROM 
				task 
			WHERE 
				user_id = $1
			AND
				deleted_at = 0`

	rows, err := taskdb.Db.Query(ctx, query, userID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("no rows found")
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		err = rows.Scan(
			&task.ID,
			&task.Name,
			&task.Type,
			&task.UserID,
			&task.Description,
			&task.DeadLine,
			&task.CreatedAt,
			&task.UpdatedAt,
			&task.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (taskdb *TaskDB) GetByUserIdAndDateRange(ctx context.Context, userID string, fromDate, toDate time.Time) ([]models.Task, error) {
	var tasks []models.Task

	query := `SELECT 
				id, 
				name, 
				type, 
				user_id, 
				description, 
				deadline, 
				created_at, 
				updated_at, 
				deleted_at 
			FROM 
				task 
			WHERE 
				user_id = $1
			AND
				deleted_at = 0
			AND
				created_at >= $2
			AND
				created_at <= $3`

	rows, err := taskdb.Db.Query(ctx, query, userID, fromDate, toDate)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("no rows found")
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		err = rows.Scan(
			&task.ID,
			&task.Name,
			&task.Type,
			&task.UserID,
			&task.Description,
			&task.DeadLine,
			&task.CreatedAt,
			&task.UpdatedAt,
			&task.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (taskdb *TaskDB) GetFinishedTasks(ctx context.Context, userID string) ([]models.Task, error) {
	var tasks []models.Task

	query := `SELECT 
				id, 
				name, 
				type, 
				user_id, 
				description, 
				deadline, 
				created_at, 
				updated_at, 
				deleted_at 
			FROM 
				task 
			WHERE 
				user_id = $1
			AND
				deleted_at = 0
			AND
				deadline < NOW()
			`

	rows, err := taskdb.Db.Query(ctx, query, userID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("no rows found")
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		err = rows.Scan(
			&task.ID,
			&task.Name,
			&task.Type,
			&task.UserID,
			&task.Description,
			&task.DeadLine,
			&task.CreatedAt,
			&task.UpdatedAt,
			&task.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (taskdb *TaskDB) GetTasksByType(ctx context.Context, userID string, types models.TaskType) ([]models.Task, error) {
	var tasks []models.Task

	query := `SELECT 
				id, 
				name, 
				type, 
				user_id, 
				description, 
				deadline, 
				created_at, 
				updated_at, 
				deleted_at 
			FROM 
				task 
			WHERE 
				user_id = $1
			AND
				deleted_at = 0
			AND
				deadline < NOW()
			`

	rows, err := taskdb.Db.Query(ctx, query, userID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("no rows found")
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		err = rows.Scan(
			&task.ID,
			&task.Name,
			&task.Type,
			&task.UserID,
			&task.Description,
			&task.DeadLine,
			&task.CreatedAt,
			&task.UpdatedAt,
			&task.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
