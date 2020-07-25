package main

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// Review ...
type Review struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"product_id"`
	Comment   string    `json:"comment"`
	Rating    int       `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
}

// AppReview ....
type AppReview struct {
	Store []Review
}

// NewAppReview ...
func NewAppReview() *AppReview {
	return &AppReview{}
}

// Mount ...
func (app *AppReview) Mount(e *echo.Echo) {
	e.GET("review/:product_id", app.GetReview)
}

// GetReview ...
func (app *AppReview) GetReview(c echo.Context) error {
	reviews := []Review{}
	productID, err := uuid.Parse(c.Param("product_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "error",
			"error":  err.Error(),
		})
	}

	for _, item := range app.Store {
		if item.ProductID == productID {
			reviews = append(reviews, item)
		}
	}

	return c.JSON(http.StatusOK, reviews)
}

func main() {
	e := echo.New()
	app := NewAppReview()
	app.Mount(e)

	e.Logger.Fatal(e.Start(":8080"))
}
