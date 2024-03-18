package handlers

import (
	"fmt"

	"28jb11/web-server/models"
)

func RenderGrid(grid [][]models.Pixel) string {
	html := "<div style=\"font-size: 0;\">\n"
	for _, row := range grid {
		for _, pixel := range row {
			html += fmt.Sprintf("<div style=\"display: inline-block; background-color: %s; width: 10px; height: 10px; margin: 1px;\"></div>\n", pixel.Color)
		}
		html += "<div style=\"clear: both;\"></div>\n" // This acts as a line break
	}
	html += "</div>\n"
	return html
}
