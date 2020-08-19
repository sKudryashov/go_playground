package matrix

func GetNumOfIslands(islands [][]string) int {
	var islandsCount int
	for v := 0; v < len(islands); v++ {
		for h := 0; h < len(islands[v]); h++ {
			if islands[v][h] == "1" {
				islandsCount++
				dfs(islands, v, h)
			}
		}
	}
	println(" islandsCount ", islandsCount)
	return islandsCount
}

func dfs(islands [][]string, v, h int) {
	// 1 if we can lookup right and it's 1
	// 2 if we can lookup down and it's 1
	lenV := len(islands)
	lenH := len(islands[v])
	islands[v][h] = "0"
	if h+1 < lenH && islands[v][h+1] == "1" {
		dfs(islands, v, h+1)
	}
	// println(" 1 ")
	if v+1 < lenV && islands[v+1][h] == "1" {
		dfs(islands, v+1, h)
	}
	// println(" 2 ")
	if h-1 >= 0 && islands[v][h-1] == "1" {
		dfs(islands, v, h-1)
	}
	// println(" 3 ")
	if v-1 >= 0 && islands[v-1][h] == "1" {
		dfs(islands, v-1, h)
	}
}
