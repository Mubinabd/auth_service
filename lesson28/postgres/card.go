package postgres

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/husanmusa/NT_Golang_10/lesson28/helper"
	"github.com/husanmusa/NT_Golang_10/lesson28/models"
)

type cardRepo struct {
	db *sql.DB
}

func NewCardRepo(db *sql.DB) *cardRepo {
	return &cardRepo{db}
}

func (c *cardRepo) Create(req *models.Card) error {
	id := uuid.NewString()

	_, err := c.db.Exec(`INSERT INTO card(id, name, card_num, expired_at, password, card_type, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		id,
		req.Name,
		req.CardNum,
		req.ExpiredAt,
		req.Password,
		req.CardType,
		req.UserID,
	)

	return err
}

func (c *cardRepo) GetById(id string) (*models.Card, error) {
	card := &models.Card{ID: id}
	err := c.db.QueryRow("SELECT name, card_num, expired_at, amount, card_type, user_id, created_at FROM card WHERE id=$1 AND deleted_at=0", id).
		Scan(&card.Name,
			&card.CardNum,
			&card.ExpiredAt,
			&card.Amount,
			&card.CardType,
			&card.UserID,
			&card.CreatedAt,
		)
	if err != nil {
		return nil, err
	}

	return card, nil
}

func (c *cardRepo) GetAll(card models.GetAllCardReq) (*models.GetAllCardsResp, error) {
	cards := []models.Card{}
	params := make(map[string]interface{}, 0)
	var arr []interface{}

	query := `SELECT name, card_num, expired_at, amount, 
			card_type, user_id, created_at FROM card `
	filter := " WHERE deleted_at = 0 "

	if len(card.CardType) > 0 {
		params["card_type"] = card.CardType
		filter += " and card_type=:card_type"
	}

	if len(card.UserId) > 0 {
		params["user_id"] = card.UserId
		filter += " and user_id=:user_id"
	}

	query = query + filter
	query, arr = helper.ReplaceQueryParams(query, params)
	fmt.Println(query, arr)

	rows, err := c.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		card := models.Card{}
		err := rows.Scan(&card.Name,
			&card.CardNum,
			&card.ExpiredAt,
			&card.Amount,
			&card.CardType,
			&card.UserID,
			&card.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		cards = append(cards, card)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	var count int
	err = c.db.QueryRow("SELECT count(1) FROM card").Scan(&count)
	if err != nil {
		return nil, err
	}

	return &models.GetAllCardsResp{Cards: &cards, Count: count}, nil
}

func (c *cardRepo) UpdatePassword(card models.UpdatePassReq) error {
	_, err := c.db.Exec(`UPDATE card SET password = $1, updated_at = now() WHERE id = $2 AND deleted_at = 0`,
		card.Pass, card.ID)

	return err
}

func (c *cardRepo) Delete(id string) error {
	_, err := c.db.Exec("UPDATE card SET deleted_at = date_part('epoch', current_timestamp)::INT WHERE id = $1 AND deleted_at = 0", id)

	return err
}
