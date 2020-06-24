export class Matrix {
  constructor(input) {
    this.matrix = input.split('\n').map(row => row.split(' ').map(Number))
  }

  get rows() {
    return this.matrix
  }

  get columns() {
    if (this.matrix.length < 2) return this.matrix
    const columns = new Array(this.matrix[0].length).fill(0).map(() => new Array(this.matrix.length).fill(0));
    this.matrix.map((row, i) => {
      row.map((e, j) => {
        columns[j][i] = e
      })
    })
    return columns
  }
}