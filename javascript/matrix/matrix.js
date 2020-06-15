export class Matrix {
  constructor(input) {
    this.matrix = input.split('\n').map(row => row.split(' ').map(Number))
  }

  get rows() {
    return this.matrix
  }

  get columns() {
    if (this.matrix.length < 2) return this.matrix
    const columns = []
    for (let i = 0; i < this.matrix[0].length; i++) {
      let column = []
      for (let j = 0; j < this.matrix.length; j++) {
        column.push(this.matrix[j][i])
      }
      columns.push(column)
    }
    return columns
  }
}