package repository

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/category/transport/handler/http/request"
)

func (r *Repository) Createcategory(ctx context.Context, obj request.Category) (*request.Category, error) {
	insertQuery := `
        INSERT INTO categories (name)
        VALUES ($1)
    `

	_, err := r.db.Slave.ExecContext(ctx, insertQuery,
		obj.Name)

	if err != nil {
		return nil, err
	}

	selectQuery := `
        SELECT id FROM categories
        WHERE name = $1
    `

	err = r.db.Slave.QueryRowContext(ctx, selectQuery, obj.Name).Scan(&obj.ID)

	if err != nil {
		return nil, err
	}

	return &obj, nil
}
