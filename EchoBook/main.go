package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/MoonSeoJun/EchoBook/controllers"
	"github.com/MoonSeoJun/EchoBook/models"
	"golang.org/x/crypto/acme/autocert"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	models.ConnectDataBase()

	e := echo.New()

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer

	e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/bookstore", controllers.GetAllBooks)
	e.GET("/bookstore/:id", controllers.GetBook)
	e.GET("/bookstore/Create", func(c echo.Context) error {
		return c.Render(http.StatusOK, "createpage.html", nil)
	})
	e.GET("/bookstore/Update/:id", controllers.GetBookToUpdate)
	e.POST("/bookstore/Create/Commplete", controllers.CreateBook)
	e.POST("/bookstore/Update/Complete/:id", controllers.UpdateBook)
	e.POST("/bookstore/Delete/:id", controllers.DeleteBook)

	e.Logger.Fatal(e.Start(":8080"))
}
