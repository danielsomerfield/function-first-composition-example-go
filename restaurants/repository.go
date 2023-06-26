package restaurants

import (
	"database/sql"
	"function-first-composition-example-go/review-server/domain"
)

func CreateGetRestaurantById(db *sql.DB) func(string) (*domain.Restaurant, error) {
	return func(id string) (*domain.Restaurant, error) {
		rows, err := db.Query("select id, name from restaurant where id = $1", id)
		defer func(rows *sql.Rows) {
			_ = rows.Close()
		}(rows)
		if err != nil {
			return nil, err
		} else {
			if rows.Next() {
				var id string
				var name string
				err = rows.Scan(&id, &name)
				if err != nil {
					return nil, err
				}
				return &domain.Restaurant{Id: id, Name: name}, nil
			} else {
				return nil, nil
			}
		}
	}
}
