package ratings

type Restaurant struct {
	Id   string
	Name string
}

type Rating int

const (
	Excellent    Rating = 2
	AboveAverage        = 1
	Average             = 0
	BelowAverage        = -1
	Terrible            = -2
)

type User struct {
	Id        string
	IsTrusted bool
}

type RestaurantRating struct {
	User   User
	Rating Rating
}

type RatingsByRestaurant struct {
	RestaurantId string
	Ratings      []RestaurantRating
}
