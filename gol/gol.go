package main

func calculateNextState(p golParams, world [][]byte) [][]byte {
	// Initialize newWorld with the same dimensions as world
	newWorld := make([][]byte, p.imageHeight)
	for i := range newWorld {
		newWorld[i] = make([]byte, p.imageWidth)
	}

	for i := 0; i < p.imageHeight; i++ {
		for j := 0; j < p.imageWidth; j++ {
			// Count living neighbors
			var sum int = 0
			//🌹⚠️这里之前用的是老师用c语言写的转换成go结果是报错，以后可以细学是什么情况
			// Check all 8 neighbors using modular arithmetic for wrapping
			for di := -1; di <= 1; di++ {
				for dj := -1; dj <= 1; dj++ {
					if di == 0 && dj == 0 {
						continue // Skip the current cell
					}
					// Use modular arithmetic to handle wrapping
					ni := (i + di + p.imageHeight) % p.imageHeight
					nj := (j + dj + p.imageWidth) % p.imageWidth
					if world[ni][nj] == 255 {
						sum++
					}
				}
			}

			// Apply Game of Life rules
			if world[i][j] == 255 { // Current cell is alive
				if sum < 2 || sum > 3 {
					newWorld[i][j] = 0 // Dies (underpopulation or overpopulation)
				} else {
					newWorld[i][j] = 255 // Survives
				}
			} else { // Current cell is dead
				if sum == 3 {
					newWorld[i][j] = 255 // Becomes alive (reproduction)
				} else {
					newWorld[i][j] = 0 // Stays dead
				}
			}
		}
	}
	return newWorld
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	var alivecells []cell
	for i := 0; i < p.imageHeight; i++ {
		for j := 0; j < p.imageWidth; j++ {
			if world[i][j] == 255 {
				alivecells = append(alivecells, cell{j, i}) // Note: x=j, y=i
			}
		}
	}
	return alivecells
}
