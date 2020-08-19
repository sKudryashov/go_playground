package arr

// rotate 1

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

// rotate in single loop 2
// class Solution {
// 	public void rotate(int[][] matrix) {
// 	  int n = matrix.length;
// 	  for (int i = 0; i < (n + 1) / 2; i ++) {
// 		for (int j = 0; j < n / 2; j++) {
// 		  int temp = matrix[n - 1 - j][i];
// 		  matrix[n - 1 - j][i] = matrix[n - 1 - i][n - j - 1];
// 		  matrix[n - 1 - i][n - j - 1] = matrix[j][n - 1 -i];
// 		  matrix[j][n - 1 - i] = matrix[i][j];
// 		  matrix[i][j] = temp;
// 		}
// 	  }
// 	}
//   }
