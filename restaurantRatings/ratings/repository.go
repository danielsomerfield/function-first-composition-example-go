package ratings

import (
	"database/sql"
)

func CreateFindRatingsByRestaurant(db *sql.DB) func(string) ([]RatingsByRestaurant, error) {
	return func(city string) (rbr []RatingsByRestaurant, err error) {
		rows, err := db.Query("select rr.restaurant_id, rr.rating, rr.rated_by_user_id, u.trusted as is_trusted from restaurant_rating rr inner join \"user\" u on rr.rated_by_user_id = u.id and city = $1", city)
		defer func(rows *sql.Rows) {
			_ = rows.Close()
		}(rows)
		if err != nil {
			return nil, err
		} else {
			m := make(map[string][]RestaurantRating)
			for rows.Next() {
				var restaurantId string
				var rating string
				var userId string
				var trusted bool
				err = rows.Scan(&restaurantId, &rating, &userId, &trusted)
				if err != nil {
					return nil, err
				}
				if m[restaurantId] == nil {
					m[restaurantId] = make([]RestaurantRating, 0)
				}
				m[restaurantId] = append(m[restaurantId], RestaurantRating{
					User: User{
						Id:        userId,
						IsTrusted: trusted,
					},
					Rating: ratings[rating],
				})
			}
			rbr = make([]RatingsByRestaurant, 0)
			for k, v := range m {
				rbr = append(rbr, RatingsByRestaurant{
					RestaurantId: k,
					Ratings:      v,
				})
			}
			return rbr, nil
		}

	}
}
