package ratings

type Rating int

const (
	Excellent    Rating = 2
	AboveAverage        = 1
	Average             = 0
	BelowAverage        = -1
	Terrible            = -2
)

func (r Rating) value() int {
	return int(r)
}

var ratings = map[string]Rating{
	"Excellent":    Excellent,
	"AboveAverage": AboveAverage,
	"Average":      Average,
	"BelowAverage": BelowAverage,
	"Terrible":     Terrible,
}

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
