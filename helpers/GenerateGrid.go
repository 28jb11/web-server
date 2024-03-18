package helpers

import "28jb11/web-server/models"

func GenerateGrid(width, height int) [][]models.Pixel {
	grid := make([][]models.Pixel, height)
	for i := range grid {
		grid[i] = make([]models.Pixel, width)
		for j := range grid[i] {
			grid[i][j] = models.Pixel{Color: "red"}
		}
	}
	return grid
}
