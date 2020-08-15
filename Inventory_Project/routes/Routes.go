package routes

import (
	"Inventory_Project/config"
	"Inventory_Project/routes/Auth"
	"Inventory_Project/routes/Barang"
	"Inventory_Project/routes/Dashboard"
	"Inventory_Project/routes/ErrorPage"
	"Inventory_Project/routes/Laporan"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"net/http"
	"strings"
	"time"
)

type Renderer struct {
	template *template.Template
	debug    bool
	location string
}
func NewRenderer(debug bool) *Renderer {
	tpl := new(Renderer)
	tpl.debug = debug

	tpl.ReloadTemplates()

	return tpl
}
func (t *Renderer) ReloadTemplates() {
	t.template = template.Must(template.ParseFiles(
		//root directory
		"./templates/index.html",

		//tmpl directory
		"./templates/tmpl/header_src.html",
		"./templates/tmpl/footer_src.html",
		"./templates/tmpl/side_nav.html",

		//auth directory
		"./templates/auth/login.html",

		//error directory
		"./templates/error/404.html",

		//barang directory
		"./templates/barang/barang.html",
		"./templates/barang/barang_keluar.html",
		"./templates/barang/barang_masuk.html",
		"./templates/barang/laporan.html",
		"./templates/barang/options.html",
		"./templates/barang/supplier.html",

		//action directory
		"./templates/action/action_barang.html",
		"./templates/action/action_barang_keluar.html",
		"./templates/action/action_barang_masuk.html",
		"./templates/action/action_laporan.html",
		"./templates/action/action_options.html",
		"./templates/action/action_supplier.html",
	))

}

func (t *Renderer) Render(
	w io.Writer,
	name string,
	data interface{},
	c echo.Context,
) error {
	if t.debug {
		t.ReloadTemplates()
	}

	return t.template.ExecuteTemplate(w, name, data)
}
func Server() *echo.Echo {
	e := echo.New()
	e.Static("/assets","assets")
	e.Renderer = NewRenderer(true)
	Routes(e)
	return e
}
func Routes(e *echo.Echo)  {
	e.Use(CheckLogin)
	e.HTTPErrorHandler = ErrorHandler
	//root
	e.GET("/index",Dashboard.Home)
	e.GET("/",Dashboard.Home)
	//auth
	e.GET("/login", Auth.Login)
	e.POST("/loginpost", Auth.LoginPost)
	e.GET("/logout",logout)
	//error
	e.GET("/404",ErrorPage.Err404)
	//barang
	e.GET("/barang",Barang.ShowMaster)
	e.GET("/options",Barang.Options)
	e.GET("/barang-masuk",Barang.Masuk)
	e.GET("/barang-keluar",Barang.Keluar)
	e.GET("/supplier",Barang.Supplier)
	//Laporan
	e.GET("/laporan",Laporan.Laporan)
}
func CheckLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		dir := context.Path()
		mthd := context.Request().Method
		if len(context.Request().Header["Accept"]) < 1{
			return next(context)
		}
		cont := context.Request().Header["Accept"][0]
		cont = strings.Split(cont,",")[0]
		cookie, err := context.Cookie(config.COOKIE_LOGIN_KEY)
		if cont == "text/html" {
			if strings.ToLower(dir) != "/login" && mthd == "GET"{
				if err != nil {
					return context.Redirect(http.StatusFound,"/login")
				}else {
					if cookie.Value != config.COOKIE_LOGIN_VAL {
						return context.Redirect(http.StatusFound,"/login")
					}
				}
			}else {
				if err == nil {
					if cookie.Value == config.COOKIE_LOGIN_VAL {
						return context.Redirect(http.StatusFound,"/index")
					}
				}
			}
		}
		return next(context)
	}
}
func ErrorHandler(err error, ctx echo.Context)  {
	error,ok := err.(*echo.HTTPError)
	if ok {
		switch error.Code {
		case http.StatusNotFound:
			ctx.Redirect(http.StatusFound,"/404")
		}
	}
}
func logout(ctx echo.Context) error {
	cookie, err := ctx.Cookie(config.COOKIE_LOGIN_KEY)
	if err != nil {
		return ctx.Redirect(http.StatusFound,"/login")
	}
	cookie.Expires = time.Unix(0, 0)
	ctx.SetCookie(cookie)
	return ctx.Redirect(http.StatusFound,"/login")
}