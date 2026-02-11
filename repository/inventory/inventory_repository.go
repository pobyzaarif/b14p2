package inventory

import (
	"context"
	"database/sql"

	"github.com/pobyzaarif/b14p2/service/inventory"
)

type SQLRepository struct {
	*sql.DB
}

func NewSQLRepository(db *sql.DB) *SQLRepository {
	return &SQLRepository{
		DB: db,
	}
}

func (r *SQLRepository) GetAll(page int, limit int) (invs []inventory.Inventory, err error) {
	ctx := context.Background()
	rows, err := r.DB.QueryContext(ctx, "SELECT code, name, description, stock FROM inventories LIMIT ?", limit)
	if err != nil {
		return invs, err
	}
	defer rows.Close()

	for rows.Next() {
		var inv inventory.Inventory

		err = rows.Scan(&inv.Code, &inv.Name, &inv.Description, &inv.Stock)
		if err != nil {
			return invs, err
		}
		invs = append(invs, inv)
	}

	return invs, err
}

func (r *SQLRepository) GetByCode(code string) (inv inventory.Inventory, err error) {
	// ctx := context.Background()
	// r.DB.WithContext(ctx).First(&inv, "code = ?", code)
	return
}
