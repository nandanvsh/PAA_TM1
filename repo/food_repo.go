package repo

import (
	"database/sql"
	"foods/model"
)

type FoodRepo interface {
	GetFoodByUserId(user_id string) ([]model.GetFood, error)
	GetFoodById(id string) (model.GetFood, error)
	AddFood(food *model.Food) error
	UpdateFood(id string, food model.Food) error
	DeleteFood(id string) error
}

type foodRepo struct {
	db *sql.DB
}

func NewFoodRepo(db *sql.DB) FoodRepo {
	return &foodRepo{db}
}

func (f *foodRepo) GetFoodByUserId(user_id string) ([]model.GetFood, error) {
	var foods []model.GetFood

	query := `select id, name, price, description from foods where user_id = $1`
	rows, err := f.db.Query(query, user_id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var item model.GetFood
		rows.Scan(
			&item.ID,
			&item.Name,
			&item.Price,
			&item.Description,
		)

		foods = append(foods, item)
	}

	return foods, nil
}

func (f *foodRepo) GetFoodById(id string) (model.GetFood, error) {
	var food model.GetFood

	query := `select id, name, price, description from foods where id = $1`
	rows, err := f.db.Query(query, id)
	if err != nil {
		return model.GetFood{}, err
	}

	for rows.Next() {
		rows.Scan(
			&food.ID,
			&food.Name,
			&food.Price,
			&food.Description,
		)
	}

	return food, nil
}

func (f *foodRepo) AddFood(food *model.Food) error {
	var id int
	query := `insert into foods (user_id, name, price, description) values ($1, $2, $3, $4) returning id`
	err := f.db.QueryRow(query, food.User_ID, food.Name, food.Price, food.Description).Scan(&id)

	if err != nil {
		return err
	}

	food.ID = id
	return nil
}

func (f *foodRepo) UpdateFood(id string, food model.Food) error {
	query := `update foods set name=$1, price=$2, description=$3 where id = $4`
	_, err := f.db.Exec(query, food.Name, food.Price, food.Description, id)
	if err != nil {
		return err
	}

	return nil
}

func (f *foodRepo) DeleteFood(id string) error {
	query := `delete from foods where id = $1`
	_, err := f.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
