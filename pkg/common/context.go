package common

import "github.com/labstack/echo/v4"

// Context represents the context of the current HTTP request.
type Context struct {
	echo.Context
}
