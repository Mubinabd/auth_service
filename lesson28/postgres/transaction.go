package postgres

import (
	"database/sql"

	"github.com/husanmusa/NT_Golang_10/lesson28/models"

	"github.com/google/uuid"
)

type TransRepo struct {
	db *sql.DB
}

func NewTransRepo(DB *sql.DB) *TransRepo {
	return &TransRepo{DB}
}

func (tr *TransRepo) Create(req *models.Transaction) error {
	id := uuid.NewString()

	_, err := tr.db.Exec("insert into transaction(id, type, description, from_card, to_card, Payment_Name, created_at) values($1, $2, $3, $4, $5, $6, $7)",
		id, req.Type, req.Description, req.FromCard, req.ToCard, req.PaymentName, req.CreatedAt)

	return err
}

func (tr *TransRepo) GetBydId(id string) (*models.Transaction, error) {
	transaction := &models.Transaction{ID: id}
	err := tr.db.QueryRow("select type, description, from_card, to_card, Payment_Name, created_at from transaction where from_card=$1 and deleted_at=0", id).
		Scan(transaction.Type, transaction.Description, transaction.FromCard, transaction.ToCard, transaction.PaymentName, transaction.CreatedAt)
	if err != nil {
		return nil, err
	}

	return transaction, err
}

func (tr *TransRepo) GetAll() (*models.GetAllTransResp, error) {
	transes := []models.Transaction{}

	rows, err := tr.db.Query("select * from transaction where deleted_at = 0")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		trans := models.Transaction{}

		err = rows.Scan(&trans.ID, &trans.Amount, &trans.Type, &trans.Description, &trans.FromCard, &trans.ToCard, &trans.PaymentName, &trans.CreatedAt, &trans.UpdatedAt, &trans.DeletedAt)
		if err != nil {
			return nil, err
		}

		transes = append(transes, trans)
	}

	var count int
	err = tr.db.QueryRow("select count(1) from users").Scan(&count)
	if err != nil {
		return nil, err
	}

	return &models.GetAllTransResp{Trans: &transes, Count: count}, nil

}

// func (tr *TransRepo) Update(trans *models.Transaction) error {
// 	_, err := tr.db.Exec("update transaction set type = $1, description = $2, from_card = $3, to_card = $4, Payment_Name = $5, updated_at = now() where id = $6",
// 		trans.Type, trans.Description, trans.FromCard, trans.ToCard, trans.PaymentName, trans.ID)

// 	return err
// }

func (tr *TransRepo) Delete(trans *models.Transaction) error {
	_, err := tr.db.Exec("update transaction set deleted_at = now() where id = $1", trans.ID)
	if err != nil {
		return err
	}

	return nil
}

func (tr *TransRepo) GetBalance(cardId string) (int, error) {
	debit, credit := 0, 0
	err := tr.db.QueryRow(`select sum(amount) filter(where type='debit') as debit, 
	sum(amount) filter(where type='credit') as credit 
		from transaction where from_card = $1 and deleted_at=0`, cardId).Scan(&debit, &credit)

	if err != nil {
		return 0, err
	}

	return debit - credit, nil
}

func (tr *TransRepo) 