package controller
import (
  "github.com/labstack/echo/v4"
  "net/http"
  "Inventory_Project/config"
  "fmt"
  "time"
)
func Login (c echo.Context) error {
  // cookie, err := c.Cookie(config.COOKIE_LOGIN_KEY)
  // if err != nil {
  //   fmt.Println(err)
  //   return err
  // }
  // if cookie.Value == config.COOKIE_LOGIN_VAL {
  //   return c.Redirect(http.StatusOK,"/index")
  // }
  data := &M{
    "message":"Halaman Login",
    "title":"Sign In",
  }
  return c.Render(http.StatusOK,"login",data)
}

func AuthLogin (c echo.Context) error {
  username := c.FormValue("username")
  password := c.FormValue("password")
  if username == "admin" && password == "admin" {
    cookie := new(http.Cookie)
    cookie.Name = config.COOKIE_LOGIN_KEY
    cookie.Value = config.COOKIE_LOGIN_VAL
    cookie.Expires = time.Now().Add(48 * time.Hour)
    c.SetCookie(cookie)
    _, err := c.Cookie(config.COOKIE_LOGIN_KEY)
    if err == nil {
      data := &M{
        "message":"Ini halaman home",
        "title" : "HOME",
      }
      return c.Render(http.StatusOK,"index",data)
    }
    fmt.Println("error:",err)
  }
  return c.String(http.StatusOK,"gagal Login")
}

func CheckCookie(next echo.HandlerFunc) echo.HandlerFunc {
  return func (c echo.Context) error {
    fmt.Println("Halo, saya ada di middleware")
    cookie, err := c.Cookie(config.COOKIE_LOGIN_KEY)
    if err != nil {
      fmt.Println("Cookie tidak ada")
      data := &M{
        "title":"Login",
      }
      return c.Render(http.StatusOK,"login",data)
    }
    if cookie.Value == config.COOKIE_LOGIN_VAL {
      fmt.Println("Sudah Login")
      return c.Redirect(http.StatusFound,"/tes")
    }
    fmt.Println("gak jalan")
    return next(c)
  }
}
