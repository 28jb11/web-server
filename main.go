package main

import (
	"28jb11/web-server/templates"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}
func main() {
	// grid := helpers.GenerateGrid(10, 10)
	// pixelGridComponent := models.GridComponent{Grid: grid}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		render(c, 200, templates.Index("name"))
	})
	r.Run()
}
