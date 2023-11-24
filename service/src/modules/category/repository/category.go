package repository

import (
	"context"
	"database/sql"
	"fmt"

	_errors "github.com/jabardigitalservice/super-app-services/event/src/error"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/category/entity"
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

func (r *Repository) Getcategories(ctx context.Context, params request.QueryParam) ([]entity.Category, error) {
	binds := make([]interface{}, 0)

	var query = fmt.Sprintf(`
        SELECT id, name
        FROM categories
        WHERE 1 = 1 %s
    `, r.filterCategoryQuery(params, &binds))

	result, err := r.getCategories(ctx, query, binds...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) getCategories(ctx context.Context, query string, args ...interface{}) ([]entity.Category, error) {
	rows, err := r.db.Slave.QueryContext(ctx, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, _errors.ErrNotFound
		}
		return nil, err
	}
	defer rows.Close()

	result := make([]entity.Category, 0)

	for rows.Next() {
		category := new(entity.Category)

		if err := rows.Scan(
			&category.ID,
			&category.Name,
		); err != nil {
			return nil, err
		}

		result = append(result, *category)
	}

	return result, nil
}

func (r *Repository) filterCategoryQuery(params request.QueryParam, binds *[]interface{}) string {
	var query string
	counter := 1

	if params.Keyword != "" {
		*binds = append(*binds, "%"+params.Keyword+"%")
		query = fmt.Sprintf(`%s AND (name ILIKE $%d)`, query, counter)
		counter++
	}

	sortBy := " ORDER BY name DESC"
	if params.SortBy != "" {
		sortBy = fmt.Sprintf(" ORDER BY %s %s", params.SortBy, params.SortOrder)
	}

	query += sortBy

	if params.PageSize > 0 {
		*binds = append(*binds, params.PageSize)
		query = fmt.Sprintf(`%s LIMIT $%d`, query, counter)
		counter++
	}

	if params.Offset > 0 {
		*binds = append(*binds, params.Offset)
		query = fmt.Sprintf(`%s OFFSET $%d`, query, counter)
	}

	return query
}

func (r *Repository) CountFilteredCategories(ctx context.Context, params request.QueryParam) (int, error) {
	binds := make([]interface{}, 0)

	query := fmt.Sprintf(`SELECT COUNT(1) FROM categories WHERE 1 = 1 %s`, r.filterCategoryCountQuery(params, &binds))

	count, err := r.countFilteredCategories(ctx, query, binds...)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *Repository) countFilteredCategories(ctx context.Context, query string, args ...interface{}) (int, error) {
	var count int
	err := r.db.Slave.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *Repository) filterCategoryCountQuery(params request.QueryParam, binds *[]interface{}) string {
	var query string
	counter := 1

	if params.Keyword != "" {
		*binds = append(*binds, "%"+params.Keyword+"%")
		query = fmt.Sprintf(`%s AND (name ILIKE $%d)`, query, counter)
		counter++
	}

	return query
}
