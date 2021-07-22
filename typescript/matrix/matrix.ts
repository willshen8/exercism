class Matrix {
  private _rows: number[][] = Matrix.parseRows(this._data)
  private _columns: number[][] = Matrix.transpose(this._rows);
  constructor(private _data: string) { }

  static parseRows(input: string): number[][] {
    return input.split('\n').map(row => row.split(' ').map(Number))
  }

  static transpose(input: number[][]): number[][] {
    return input[0].map((_, i) => input.map(row => row[i]))
  }

  public get rows(): number[][] {
    return this._rows
  }

  public get columns(): number[][] {
    return this._columns
  }
}

export default Matrix
