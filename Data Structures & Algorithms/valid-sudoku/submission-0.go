import "slices"
func isValidSudoku(board [][]byte) bool {
    cols := make(map[byte][]byte, 9)
    rows := make(map[byte][]byte, 9)
    boxes := make(map[byte][]byte, 9)

    for row := 0; row < len(board); row++ {
        for col := 0; col < len(board[row]); col++ {

            cell := board[row][col]

            if cell == '.' {
                continue
            }

            boxKey := byte((row/3)*3 + (col/3))

            if slices.Contains(rows[byte(row)],cell) || 
                slices.Contains(cols[byte(col)],cell) || 
                slices.Contains(boxes[byte(boxKey)],cell) {
               
                return false
            }

            rows[byte(row)] = append(rows[byte(row)],cell)
            cols[byte(col)] = append(cols[byte(col)],cell)
            boxes[byte(boxKey)] = append(boxes[byte(boxKey)],cell)
        }
    }

    return true
}