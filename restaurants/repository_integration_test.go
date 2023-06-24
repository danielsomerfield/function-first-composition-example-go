package restaurants

import (
	"context"
	"function-first-composition-example-go/review-server/testutil"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetRestaurantByIdWithData(t *testing.T) {

	ctx := context.Background()

	db := testutil.StartDB(ctx, "../db/init.sql")
	t.Cleanup(func() {
		if err := db.Stop(ctx); err != nil {
			t.Fatalf("failed to terminate db: %s", err.Error())
		}
	})

	id := "cafegloucesterid"
	name := "Cafe Gloucester"
	testutil.CreateRestaurant(db, testutil.Restaurant{Id: id, Name: name})
	getRestaurantById := CreateGetRestaurantById(db.(testutil.DBImplementation).GetDB())
	restaurant, err := getRestaurantById(id)
	if err != nil {
		t.Fatalf("Received error %v", err)
	}
	assert.Equal(t, restaurant.Id, id)
	assert.Equal(t, restaurant.Name, name)
}

func TestGetRestaurantByIdNotFound(t *testing.T) {

	ctx := context.Background()

	db := testutil.StartDB(ctx, "../db/init.sql")
	t.Cleanup(func() {
		if err := db.Stop(ctx); err != nil {
			t.Fatalf("failed to terminate db: %s", err.Error())
		}
	})

	getRestaurantById := CreateGetRestaurantById(db.(testutil.DBImplementation).GetDB())
	restaurant, err := getRestaurantById("does not exist")
	if err != nil {
		t.Fatalf("Received error %v", err)
	}
	if restaurant != nil {
		t.Fatalf("Received unexpected value %v", restaurant)
	}
}
