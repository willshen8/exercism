export class Matrix {
  constructor(data) {
    this._data = data
  }

  static parseRows(input) {
    return input.split('\n').map(row => row.split(' ').map(Number))
  }

  static transpose(input) {
    return input[0].map((_, i) => input.map(row => row[i]))
  }

  get rows() {
    if (!this._rows) {
      this._rows = Matrix.parseRows(this._data)
    }
    return this._rows
  }

  get columns() {
    if (!this._columns) {
      if (!this._rows) this.rows
      return Matrix.transpose(this._rows)
    }
    return this._columns
  }
}