package Barang

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"Inventory_Project/config"
)

func ShowMaster(ctx echo.Context) error {
	data := &config.M{
		"title":"Data Barang",
		"path": "barang",
	}
	ctx.Render(http.StatusOK,"header",data)
	ctx.Render(http.StatusOK,"sidenav",data)
	return ctx.Render(http.StatusOK,"barang.html",data)
}

func Options(ctx echo.Context) error {
	data := &config.M{
		"title":"Options",
		"path": "barang",
	}
	ctx.Render(http.StatusOK,"header",data)
	ctx.Render(http.StatusOK,"sidenav",data)
	return ctx.Render(http.StatusOK,"options.html",data)
}
func Masuk(ctx echo.Context) error {
	data := &config.M{
		"title": "Barang Masuk",
		"path": "barang-masuk",
	}
	ctx.Render(http.StatusOK,"header",data)
	ctx.Render(http.StatusOK,"sidenav",data)
	return ctx.Render(http.StatusOK,"barang_masuk.html",data)
}
func Keluar(ctx echo.Context) error {
	data := &config.M{
		"title": "Barang Keluar",
		"path": "barang-keluar",
	}
	ctx.Render(http.StatusOK,"header",data)
	ctx.Render(http.StatusOK,"sidenav",data)
	return ctx.Render(http.StatusOK,"barang_keluar.html",data)
}
func Supplier(ctx echo.Context) error {
	data := &config.M{
		"title": "Supplier",
		"path": "supplier",
	}
	ctx.Render(http.StatusOK,"header",data)
	ctx.Render(http.StatusOK,"sidenav",data)
	return ctx.Render(http.StatusOK,"supplier.html",data)
}