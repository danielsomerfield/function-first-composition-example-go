package ratings

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"go/types"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController_RespondsWithJSONRatings(t *testing.T) {
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = []gin.Param{
		{
			Key:   "city",
			Value: "nyc",
		},
	}

	restaurant := Restaurant{
		Id:   "cafegloucesterid",
		Name: "Cafe Gloucester",
	}

	vancouverRestaurants := []Restaurant{
		restaurant,
	}

	dependenciesStub := Dependencies{
		GetTopRestaurants: func() ([]Restaurant, error) {
			return vancouverRestaurants, nil
		},
	}

	controller := createController(&dependenciesStub)
	controller(c)
	assert.Equal(t, http.StatusOK, recorder.Code)
	responseBody := struct {
		Restaurants []struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"restaurants"`
	}{}
	_ = json.Unmarshal(recorder.Body.Bytes(), &responseBody)
	assert.Equal(t, len(responseBody.Restaurants), 1)
	assert.Equal(t, responseBody.Restaurants[0].Id, restaurant.Id)
	assert.Equal(t, responseBody.Restaurants[0].Name, restaurant.Name)
}

func TestController_RespondsWith500OnUnexpectedError(t *testing.T) {
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	dependenciesStub := Dependencies{
		GetTopRestaurants: func() ([]Restaurant, error) {
			return nil, types.Error{}
		},
	}

	controller := createController(&dependenciesStub)
	controller(c)
	assert.Equal(t, recorder.Code, http.StatusInternalServerError)

}
