package matrix

func rotate(mtrx [][]int) {
	n := len(mtrx)
	till := (n+1)/2 + n%2
	for i := 0; i < till; i++ {
		for j := 0; j < n/2; j++ {
			tmp := make([]int, 4)
			row := i
			col := j
			for k := 0; k < 4; k++ {
				tmp[k] = mtrx[row][col]
				x := row
				row = col
				col = n - 1 - x
			}
			for k := 0; k < 4; k++ {
				till2 := (k + 3) % 4
				mtrx[row][col] = tmp[till2]
				x := row
				row = col
				col = n - 1 - x
			}
		}
	}
}

// class Solution {
// 	public void rotate(int[][] matrix) {
// 	  int n = matrix.length;
// 	  for (int i = 0; i < n / 2 + n % 2; i++) {
// 		for (int j = 0; j < n / 2; j++) {
// 		  int[] tmp = new int[4];
// 		  int row = i;
// 		  int col = j;
// 		  for (int k = 0; k < 4; k++) {
// 			tmp[k] = matrix[row][col];
// 			int x = row;
// 			row = col;
// 			col = n - 1 - x;
// 		  }
// 		  for (int k = 0; k < 4; k++) {
// 			matrix[row][col] = tmp[(k + 3) % 4];
// 			int x = row;
// 			row = col;
// 			col = n - 1 - x;
// 		  }
// 		}
// 	  }
// 	}
//   }
