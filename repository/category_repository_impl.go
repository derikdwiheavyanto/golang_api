package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang_restful_api/helper"
	"golang_restful_api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "insert into category(name) values (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	category.Id = int(id)
	return category
}
func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQl := "update category set name = ? where id= ?"

	_, err := tx.ExecContext(ctx, SQl, category.Name, category.Id)
	helper.PanicIfErr(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "delete from category where id = ?"

	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfErr(err)
}

func (reposityory *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {

	SQL := "select id,name from category where id = ?"

	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfErr(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfErr(err)
		return category, nil
	} else {
		return category, errors.New("category not found")
	}
}

func (reposityory *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "select id,name from category "

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfErr(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfErr(err)

		categories = append(categories, category)
	}

	return categories
}
