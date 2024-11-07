package product

import (
	"database/sql"
	"errors"

	"github.com/lucasti79/bgw4-put-patch-delete/internal/domain"
	"github.com/lucasti79/bgw4-put-patch-delete/pkg/apperrors"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) domain.Repository {
	return &repository{db}
}

func (r repository) Get() (p []domain.Product, err error) {
	var products []domain.Product

	rows, err := r.db.Query(GetAll)

	if err != nil {
		return products, err
	}

	for rows.Next() {
		var product domain.Product

		if err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Quantity,
			&product.Type,
		); err != nil {
			return products, err
		}

		products = append(products, product)
	}
	return products, nil
}

func (r repository) GetByID(id int) (p *domain.Product, err error) {
	// execute query
	row := r.db.QueryRow(
		GetById,
		id,
	)
	if err := row.Err(); err != nil {
		return nil, err
	}
	// scan result
	var product domain.Product
	if err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Price,
		&product.Quantity,
		&product.Type,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.ErrNotFound
		}
		return nil, err
	}

	return &product, nil
}

func (r repository) UpdateOrCreate(p *domain.Product) (err error) {
	result, err := r.db.Exec(
		Create,
		p.Name, p.Price, p.Quantity, p.Type,
	)
	if err != nil {
		return err
	}
	// get last inserted id
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	// set user id
	p.Id = int(lastInsertId)
	return nil
}

func (r repository) Update(id int, p *domain.Product) (err error) {
	_, err = r.db.Exec(
		"UPDATE products SET name=?, price=?, quantity=?, type=? WHERE id=?",
		p.Name, p.Price, p.Quantity, p.Type, p.Id,
	)
	if err != nil {
		return err
	}
	return nil
}
func (r repository) Delete(id int) (err error) {
	// executando um statement
	stmt, err := r.db.Prepare(Delete)

	// executando uma transaction
	// transaction, err := r.db.BeginTx(context.Background(), nil)
	// transaction.Exec()
	// transaction.Commit()
	// transaction.Rollback()

	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(id)

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return apperrors.ErrNotFound
	}

	return err
}
