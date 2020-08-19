package matrix

import "fmt"

/*

Imagine we have an image. We'll represent this image as a simple 2D array where every pixel is a 1 or a 0. The image you get is known to have a single rectangle of 0s on a background of 1s.

Write a function that takes in the image and returns one of the following representations of the rectangle of 0's: top-left coordinate and bottom-right coordinate OR top-left coordinate, width, and height.

image1 = [
  [1, 1, 1, 1, 1, 1, 1],
  [1, 1, 1, 1, 1, 1, 1],
  [1, 1, 1, 0, 0, 0, 1],
  [1, 1, 1, 0, 0, 0, 1],
  [1, 1, 1, 1, 1, 1, 1],
]

Sample output variations (only one is necessary):

findRectangle(image1) =>
  x: 3, y: 2, width: 3, height: 2
  2,3 3,5 -- row,column of the top-left and bottom-right corners

Other test cases:

image2 = [
  [1, 1, 1, 1, 1, 1, 1],
  [1, 1, 1, 1, 1, 1, 1],
  [1, 1, 1, 1, 1, 1, 1],
  [1, 1, 1, 1, 1, 1, 1],
  [1, 1, 1, 1, 1, 1, 0],
]

findRectangle(image2) =>
  x: 6, y: 4, width: 1, height: 1
  4,6 4,6

image3 = [
  [1, 1, 1, 1, 1, 1, 1],
  [1, 1, 1, 1, 1, 1, 1],
  [1, 1, 1, 1, 1, 1, 1],
  [1, 1, 1, 1, 1, 0, 0],
  [1, 1, 1, 1, 1, 0, 0],
]

findRectangle(image3) =>
  x: 5, y: 3, width: 2, height: 2
  3,5 4,6

image4 = [
  [0, 1, 1, 1, 1, 1, 1],
  [1, 1, 1, 1, 1, 1, 1],
  [1, 1, 1, 1, 1, 1, 1],
  [1, 1, 1, 1, 1, 1, 1],
  [1, 1, 1, 1, 1, 1, 1],
]

findRectangle(image4) =>
  x: 0, y: 0, width: 1, height: 1
  0,0 0,0

image5 = [
  [0],
]

findRectangle(image5) =>
  x: 0, y: 0, width: 1, height: 1
  0,0 0,0

n: number of rows in the input image
m: number of columns in the input image

*/

func Launch() {
	// image2 := [][]int{
	// 	[]int{1, 1, 1, 1, 1, 1, 1},
	// 	[]int{1, 1, 1, 1, 1, 1, 1},
	// 	[]int{1, 1, 1, 1, 1, 1, 1},
	// 	[]int{1, 1, 1, 1, 1, 1, 1},
	// 	[]int{1, 1, 1, 1, 1, 1, 0},
	// }

	image3 := [][]int{
		[]int{1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 0, 0},
		[]int{1, 1, 1, 1, 1, 0, 0},
	}
	x, y := FindOutput(image3)
	fmt.Printf("image3 %d %d %d %d ", x[0], x[1], y[0], y[1])

	image4 := [][]int{
		[]int{0, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1, 1},
	}

	x, y = FindOutput(image4)
	fmt.Printf(" %d %d %d %d ", x[0], x[1], y[0], y[1])

	image5 := [][]int{
		[]int{0},
	}

	x, y = FindOutput(image5)
	fmt.Printf(" %d %d %d %d ", x[0], x[1], y[0], y[1])

	image1 := [][]int{
		[]int{1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 0, 0, 0, 1},
		[]int{1, 1, 1, 0, 0, 0, 1},
		[]int{1, 1, 1, 1, 1, 1, 1},
	}

	x, y = FindOutput(image1)
	fmt.Printf(" %d %d %d %d ", x[0], x[1], y[0], y[1])
}

func FindOutput(image1 [][]int) ([]int, []int) {
	var topLeftX, topLeftY int
	var bottomX, bottomY int
	// loop:
	for i := 0; i < len(image1); i++ {
		for j := 0; j < len(image1[i]); j++ {
			if image1[i][j] == 0 {
				topLeftX = j
				topLeftY = i
				x, y := searchRoutine(image1, i, j+1)
				println("x , y ", x, y)
			}
		}
	}
	return []int{topLeftX, topLeftY}, []int{bottomX, bottomY}
}

func searchRoutine(image1 [][]int, i, j int) (int, int) {
	var x, y int
	if j+1 < len(image1[i]) && image1[i][j+1] == 0 {
		x, y = searchRoutine(image1, i, j+1)
		return x, y
	}
	if i+1 < len(image1) && image1[i+1][j] == 0 {
		searchRoutine(image1, i+1, j)
		x, y = i+1, j
		return x, y
	}
	return x, y
}
