package main

import (
	"net/http"
	"os"

	"github.com/davecgh/go-spew/spew"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

type Body struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func main() {
	e := echo.New()
	e.GET("/", DefaultHandler)

	e.GET("/private-endpoint", DefaultHandler, MiddlewareAPIKey)

	groupProduct := e.Group("/products")
	groupProduct.GET("", func(c echo.Context) error {
		page := c.QueryParam("page")
		pageSize := c.QueryParam("pageSize")

		return c.JSON(http.StatusOK, echo.Map{"message": "products endpoint" + page + pageSize})
	})
	groupProduct.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")

		return c.JSON(http.StatusOK, echo.Map{"message": "products endpoint" + id})
	})
	groupProduct.POST("", func(c echo.Context) error {
		var payload Body

		err := c.Bind(&payload)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "payload error")
		}

		spew.Dump(payload)
		return c.JSON(http.StatusCreated, echo.Map{"message": "Created"})
	})

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}

func DefaultHandler(c echo.Context) error {
	userID, _ := c.Get("user_id").(string)
	return c.JSON(http.StatusOK, echo.Map{"hello": "world" + userID})
}

func MiddlewareAPIKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		apiKey := c.Request().Header.Get("x-api-key")
		if apiKey != os.Getenv("APP_API_KEY") {
			return echo.NewHTTPError(http.StatusForbidden, http.StatusText(http.StatusForbidden))
		}

		c.Set("user_id", "1")

		return next(c)
	}
}
