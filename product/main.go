package main

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Product ...``
type Product struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Price int       `json:"price"`
}

// NewProductStore ...
func NewProductStore() []*Product {
	return []*Product{
		&Product{ID: uuid.New(), Name: "Flash Disk 64", Price: 8000},
		&Product{ID: uuid.New(), Name: "Flash Disk 128", Price: 720000},
		&Product{ID: uuid.New(), Name: "Printer HP", Price: 720000},
		&Product{ID: uuid.New(), Name: "Kertas A4", Price: 50000},
	}
}

// AppProduct ...
type AppProduct struct {
	Store []Product
}

// NewAppProduct ...
func NewAppProduct() *AppProduct {
	return &AppProduct{
		Store: []Product{
			Product{ID: uuid.New(), Name: "Flash Disk 64", Price: 8000},
			Product{ID: uuid.New(), Name: "Flash Disk 128", Price: 720000},
			Product{ID: uuid.New(), Name: "Printer HP", Price: 720000},
			Product{ID: uuid.New(), Name: "Kertas A4", Price: 50000},
		},
	}
}

// Mount ...
func (app *AppProduct) Mount(e *echo.Echo) {
	e.GET("/product", app.List)
	e.GET("/product/:id", app.Detail)
	e.POST("/product", app.Create)
}

// List ...
func (app *AppProduct) List(c echo.Context) error {
	return c.JSON(http.StatusOK, app.Store)
}

// Detail ...
func (app *AppProduct) Detail(c echo.Context) error {
	product := Product{}
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "error",
			"error":  err.Error(),
		})
	}

	for _, item := range app.Store {
		if item.ID == id {
			product = item
			break
		}
	}

	if (Product{}) == product {
		return c.JSON(http.StatusNotFound, map[string]string{
			"status": "error",
			"error":  "product not found",
		})
	}

	return c.JSON(http.StatusOK, product)
}

// Create ...
func (app *AppProduct) Create(c echo.Context) error {
	name := c.FormValue("name")
	price, err := strconv.Atoi("price")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "error",
			"error":  err.Error(),
		})
	}

	app.Store = append(app.Store, Product{
		ID:    uuid.New(),
		Name:  name,
		Price: price,
	})

	return c.JSON(http.StatusOK, app.Store)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	app := NewAppProduct()
	app.Mount(e)

	e.Logger.Fatal(e.Start(":8080"))
}
