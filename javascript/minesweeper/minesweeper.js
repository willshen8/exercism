export const annotate = (input) => {
  const mineMatrix = input.map(row => row.split(''))
  const output = mineMatrix.map(
    (row, rowIndex) => {
      row.map((square, columnIndex) => {
        if (square !== '*') {
          console.log("input=", input)
          console.log("Square =", square, "RowIndex=", rowIndex, "ColumnIndex=", columnIndex)
          if (rowIndex - 1 > 0) {
            if (mineMatrix[rowIndex][columnIndex - 1] === '*') {
              return square++
            }
            if (mineMatrix[rowIndex][columnIndex] === '*') {
              return square++
            }
            if (columnIndex + 1 <= mineMatrix[rowIndex - 1].length)
              if (mineMatrix[rowIndex][columnIndex + 1] === '*') {
                square++
              }
          }
          if (rowIndex - 1 < mineMatrix[rowIndex - 1].length) {
            if (columnIndex + 1 <= mineMatrix[rowIndex - 1].length) {
              if (mineMatrix[rowIndex - 1][columnIndex - 1] === '*') {
                return square++
              }
            }
            if (mineMatrix[rowIndex - 1][columnIndex] === '*') {
              return square++
            }
            if (columnIndex + 1 <= mineMatrix[rowIndex - 1].length)
              if (mineMatrix[rowIndex - 1][columnIndex + 1] === '*') {
                return square++
              }
          }
          if (columnIndex - 1 > 0) {
            if (mineMatrix[rowIndex][columnIndex - 1] === '*') {
              return square++
            }
          }
          if (columnIndex + 1 > 0) {
            if (mineMatrix[rowIndex][columnIndex + 1] === '*') {
              return square++
            }
          }
        }
      })
    })
  return output
}