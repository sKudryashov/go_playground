package arr

// def findSignatureCounts(arr):
//   output = [-1] * len(arr)

//   def countSignatures(student_id: int, next_student: int, total_signatures: int = 1) -> int:
//     if student_id == next_student:
//       return total_signatures
//     return countSignatures(student_id, arr[next_student-1], 1 + total_signatures)

//   for i, e in enumerate(arr):
//     if i == e-1: # they are passing to themselves
//       output[i] = 1
//     else:
//       output[i] = countSignatures(i+1, arr[i])
//   return output
