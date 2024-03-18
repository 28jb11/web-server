package handlers

import (
	"fmt"

	"28jb11/web-server/models"
)

func RenderGrid(grid [][]models.Pixel) string {
	html := "<div>\n"
	for _, row := range grid {
		for _, pixel := range row {
			html += fmt.Sprintf("<div class=\"pixel-container\"><div class=\"pixel\"></div><span class=\"pixel-coords\">%d,%d</span></div>\n", pixel.PosX, pixel.PosY)
		}
		html += "<div style=\"clear: both;\"></div>\n" // This acts as a line break
	}
	html += "</div>\n"
	return html
}

func GenerateGrid(width, height int) [][]models.Pixel {
	grid := make([][]models.Pixel, height)
	for i := range grid {
		grid[i] = make([]models.Pixel, width)
		for j := range grid[i] {
			grid[i][j] = models.Pixel{Color: "grey"}
			// assign grid coordinates to each pixel in the grid
			grid[i][j].PosX = j
			grid[i][j].PosY = i
		}
	}
	return grid
}

func GetGridHtmlString() string {
	var s string
	gridComponent := models.GridComponent{Grid: GenerateGrid(20, 20)}
	s = RenderGrid(gridComponent.Grid)
	return s
}
