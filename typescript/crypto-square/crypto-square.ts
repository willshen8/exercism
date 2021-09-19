const NonAlphanumericRegex = /[^a-zA-Z0-9+]/g

interface Dimension {
  rows: number
  columns: number
}
export class Crypto {
  private _plainText: unknown
  private _cipherText: string = ''
  constructor(plainText: unknown) {
    this._plainText = plainText
  }

  get ciphertext(): unknown {
    // lazy load encoded message - only calculate when requested
    if (!this._cipherText) {
      let cipherText = ''
      if (this._plainText && typeof this._plainText === "string") {
        cipherText = this._plainText.replace(NonAlphanumericRegex, "").toLowerCase();

        const { rows, columns } = this.calculateRectangleDimension(cipherText)
        const normalisedTextArray = this.splitStringToArray(cipherText, rows, columns)

        cipherText = this.encodeMessage(normalisedTextArray, rows, columns)
      }
      this._cipherText = cipherText
    }
    return this._cipherText
  }

  private encodeMessage(input: Array<string>, rows: number, columns: number): string {
    let cipherTextMatrix: Array<string> = []
    for (let i = 0; i < columns; i++) {
      let cipherRow = ''
      for (let j = 0; j < rows; j++) {
        cipherRow += input[j][i]
      }
      cipherTextMatrix[i] = cipherRow
    }
    return cipherTextMatrix.join(' ')
  }

  private calculateRectangleDimension(input: string): Dimension {
    const columns = Math.ceil(Math.sqrt(input.length))
    let rows = Math.floor(Math.sqrt(input.length))
    if (rows * columns < input.length) {
      rows++
    }
    return { rows: rows, columns: columns }
  }

  private splitStringToArray(input: string, rows: number, columns: number): Array<string> {
    if (input.length < rows * columns) {
      input = input.padEnd(rows * columns, ' ')
    }

    let outputArray = Array(rows)
    for (let i = 0; i < rows; i++) {
      outputArray[i] = input.substring(i * columns, (i + 1) * columns)
    }
    return outputArray
  }
}
