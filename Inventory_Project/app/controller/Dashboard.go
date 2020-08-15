package controller
import (
  "github.com/labstack/echo/v4"
  "net/http"
)
func Dashboard (c echo.Context) error {
  data := M{
    "message":"Hello World",
    "title":"Dashboard",
  }
  return c.Render(http.StatusOK,"index",data)
}
