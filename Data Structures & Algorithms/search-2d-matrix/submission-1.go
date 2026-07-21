func searchMatrix(matrix [][]int, target int) bool {
    m, n := len(matrix), len(matrix[0])
    r, c := 0, n - 1

    for r < m && c >= 0 {
        if matrix[r][c] > target {
            c--
        } else if matrix[r][c] < target {
            r++
        } else {
            return true
        }
    }
    return false
}
