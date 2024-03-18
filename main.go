package main

import (
	"28jb11/web-server/handlers"
	"28jb11/web-server/templates"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}
func main() {

	// grid := handlers.GenerateGrid(20, 20)
	//
	// gridComponent := models.GridComponent{Grid: grid}
	//
	// gridHtml := handlers.RenderGrid(gridComponent.Grid)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		render(c, 200, templates.Index("name"))
	})

	r.GET("/grid", func(c *gin.Context) {
		render(c, 200, templates.Grid(handlers.GetGridHtmlString()))
	})

	r.Run()
}
