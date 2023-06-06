package testutil

import "log"

type User struct {
	Id      string
	Name    string
	Trusted bool
}

type Restaurant struct {
	Id   string
	Name string
}

type RatingForRestaurant struct {
	Id         string
	User       User
	Restaurant Restaurant
	Rating     string
}

func CreateUser(db DB, user User) {
	err := db.Exec("INSERT INTO \"user\"(id, name, trusted) VALUES ($1, $2, $3)", user.Id, user.Name, user.Trusted)
	if err != nil {
		log.Fatalf("Failed to create user %v", err)
	}
}

func CreateRestaurant(db DB, restaurant Restaurant) {
	err := db.Exec("INSERT INTO restaurant(id, name) VALUES ($1, $2)", restaurant.Id, restaurant.Name)
	if err != nil {
		log.Fatalf("Failed to create restaurant %v", err)
	}
}

func CreateRatingByUserForRestaurant(db DB, rating RatingForRestaurant) {
	err := db.Exec(
		"insert into restaurant_rating (id, rated_by_user_id, restaurant_id, rating, city) VALUES ($1, $2, $3, $4, 'vancouverbc')",
		rating.Id, rating.User.Id, rating.Restaurant.Id, rating.Rating,
	)
	if err != nil {
		log.Fatalf("Failed to create rating %v", err)
	}
}
