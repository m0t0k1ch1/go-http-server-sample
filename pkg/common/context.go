package common

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type emptyResponse struct{}

// Context represents the context of the current HTTP request.
type Context struct {
	echo.Context
}

// Success sends a JSON response with status code 200.
func (c *Context) Success(v interface{}) error {
	return c.JSON(http.StatusOK, v)
}

// SuccessWithEmpty sends an empty JSON response with status code 200.
func (c *Context) SuccessWithEmpty() error {
	return c.Success(emptyResponse{})
}

// NotFound sends an error response with status code 404.
func (c *Context) NotFound(message string) error {
	return echo.NewHTTPError(http.StatusNotFound, message)
}

// InternalServerError sends an error response with status code 500.
func (c *Context) InternalServerError(err error) error {
	c.Logger().Error(err)
	return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
}
