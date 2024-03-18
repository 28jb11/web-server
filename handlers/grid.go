package handlers

import (
	"fmt"

	"28jb11/web-server/models"

	"github.com/a-h/templ"
)

func RenderGrid(grid [][]models.Pixel) templ.Component {
	html := "<div class=\"grid-container\">\n"
	for _, row := range grid {
		for _, pixel := range row {
			html += fmt.Sprintf("<div class=\"pixel\" style=\"background-color: %s;\"></div>\n", pixel.Color)
		}
	}
	html += "</div>"
	return templ.Raw(html)
}
