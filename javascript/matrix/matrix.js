export class Matrix {
  constructor(data) {
    this._data = data
  }

  static parseRows(input) {
    return input.split('\n').map(row => row.split(' ').map(Number))
  }

  get rows() {
    if (!this._rows) {
      this._rows = Matrix.parseRows(this._data)
    }
    return this._rows
  }

  get columns() {
    if (!this._rows) {
      this._rows = Matrix.parseRows(this._data)
    }
    return this._rows[0].map(
        (_, i) => this._rows.map((_, j) => this._rows[j][i]))
  }
}