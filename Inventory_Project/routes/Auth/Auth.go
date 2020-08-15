package Auth

import (
	"Inventory_Project/config"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)
type M map[string]interface {}
func Login(ctx echo.Context) error {
	alert,err := ctx.Cookie(config.COOKIE_ALERT)
	data := &M{
		"title":"Login",
		"alert":"",
	}
	if err == nil{
		data = &M{"alert":alert.Value}
	}
	return ctx.Render(http.StatusOK,"login.html",data)
}
func LoginPost(ctx echo.Context) error {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")
	if username == "admin" && password == "admin" {
		cookie := &http.Cookie{
			Name: config.COOKIE_LOGIN_KEY,
			Value: config.COOKIE_LOGIN_VAL,
			Expires: time.Now().Add(config.COOKIE_LOGIN_EXPIRES * time.Hour),
		}
		ctx.SetCookie(cookie)
	}
	return ctx.Redirect(http.StatusFound,"/login")
}
