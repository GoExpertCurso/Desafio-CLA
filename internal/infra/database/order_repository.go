package database

import (
	"database/sql"

	"github.com/GoExpertCurso/Desafio-CLA/internal/entity"
	"github.com/hashicorp/go-multierror"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) ListOrders() ([]entity.Order, error) {
	rows, err := r.Db.Query(`select id, price, tax, final_price from orders`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	orders := []entity.Order{}
	var scanErr error

	for rows.Next() {
		var id string
		var price, tax, finalPrice float64

		if err := rows.Scan(&id, &price, &tax, &finalPrice); err != nil {
			scanErr = multierror.Append(scanErr, err)
			continue
		}

		orders = append(orders, entity.Order{
			ID:         id,
			Price:      price,
			Tax:        tax,
			FinalPrice: finalPrice,
		})
	}
	if err := rows.Err(); err != nil {
		return nil, multierror.Append(scanErr, err)
	}

	return orders, nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("Select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
