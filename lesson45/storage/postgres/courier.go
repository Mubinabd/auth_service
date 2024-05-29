package postgres

import (
	"database/sql"
	pb "github.com/husanmusa/NT_Golang_10/lesson45/genproto/courier"
)

type CourierRepo struct {
	db *sql.DB
}

func NewCourierRepo(db *sql.DB) *CourierRepo {
	return &CourierRepo{db: db}
}

func (c *CourierRepo) Deliver(or *pb.TakeOrder) (*pb.DeliverOrder, error) {
	rows, err := c.db.Query(`select delivered_time, p.id, p.name
								from orders o
							left join selected_product sp on sp.order_id=o.id
							left join products p on sp.product_id=p.id
							where o.name = $1`, or.GetName())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*pb.Product
	var tm string
	for rows.Next() {
		var pr pb.Product
		err = rows.Scan(&tm, &pr.Id, &pr.Name)
		if err != nil {
			return nil, err
		}
		products = append(products, &pr)
	}

	return &pb.DeliverOrder{Time: tm, Products: products}, nil
}
