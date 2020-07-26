package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	e.POST("review/:product_id", app.CreateReview)
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

// CreateReview ...
func (app *AppReview) CreateReview(c echo.Context) error {
	productID, err := uuid.Parse(c.Param("product_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "error",
			"error":  err.Error(),
		})
	}

	comment := c.FormValue("comment")
	rating, err := strconv.Atoi(c.FormValue("rating"))

	review := Review{
		ID:        uuid.New(),
		ProductID: productID,
		Comment:   comment,
		Rating:    rating,
		CreatedAt: time.Now(),
	}

	app.Store = append(app.Store, review)

	return c.JSON(http.StatusOK, review)
}

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","form":"${form}","query":"${query}","status":${status},"error":"${error}","latency":${latency},` +
			`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
			`"bytes_out":${bytes_out}}` + "\n",
	}))

	app := NewAppReview()
	app.Mount(e)

	e.Logger.Fatal(e.Start(":8080"))
}
