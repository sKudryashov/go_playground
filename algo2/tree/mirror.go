package tree

// public boolean isSymmetric(TreeNode root) {
//     return isMirror(root, root);
// }

// public boolean isMirror(TreeNode t1, TreeNode t2) {
//     if (t1 == null && t2 == null) return true;
//     if (t1 == null || t2 == null) return false;
//     return (t1.val == t2.val)
//         && isMirror(t1.right, t2.left)
//         && isMirror(t1.left, t2.right);
// }

// Let’s implement an in-memory filesystem with the following interface:

// mkdir(path:string)
// write_file(path:string, data:string)
// read_file(path:string) -> string
// The filesystem structure and all data should be stored in the program’s memory.

// ------------- Clarifications --------------

// Path arguments are expected to be well-formed and no validation is required. A well-formed path has a leading slash, and no trailing slash. For example:

// /foo/bar
// /foo/bar/foo/bar/
// /foo/bar/file
// The implementation has several edge cases, which the solution needs to handle:

// The mkdir and write_file operations should fail if the parent directory does not exist.
// If the path in mkdir already exists and is a directory, mkdir should fail
// If the path for mkdir already exists and is a file, mkdir should fail.
// The read_file operation should fail if the file does not exist.
// The read_file and write_file operations should fail if the path points to a directory.
// There’s no notion of relative paths, paths are all absolute.

// package main
// import (
//     "fmt"
//     "strings"
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
// 
// )
