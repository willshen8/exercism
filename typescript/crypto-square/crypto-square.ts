const NonAlphanumericRegex = /[^a-zA-Z0-9+]/g

interface Dimension {
  rows: number
  columns: number
}
export class Crypto {
  private _plainText: string
  private _cipherText: string = ''
  constructor(plainText: string) {
    this._plainText = plainText
  }

  get ciphertext(): string {
    // lazy load encoded message - only calculate when requested
    if (!this._cipherText) {
      let cipherText = this._plainText.replace(NonAlphanumericRegex, "").toLowerCase();
      const { rows, columns } = this.calculateRectangleDimension(cipherText)
      const normalisedTextArray = this.splitStringToArray(cipherText, rows, columns)
      cipherText = this.encodeMessage(normalisedTextArray, rows, columns)
      this._cipherText = cipherText
    }
    return this._cipherText
  }

  private encodeMessage(input: Array<string>, rows: number, columns: number): string {
    const cipherTextMatrix: Array<string> = []
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
    const sqrRootOfInputLength = Math.sqrt(input.length)
    const columns = Math.ceil(sqrRootOfInputLength)
    let rows = Math.floor(sqrRootOfInputLength)
    if (rows * columns < input.length) {
      rows++
    }
    return { rows, columns }
  }

  private splitStringToArray(input: string, rows: number, columns: number): Array<string> {
    let paddedInput = input.slice()
    if (input.length < rows * columns) {
      paddedInput = paddedInput.padEnd(rows * columns, ' ')
    }

    const outputArray = Array(rows)
    for (let i = 0; i < rows; i++) {
      outputArray[i] = paddedInput.substring(i * columns, (i + 1) * columns)
    }
    return outputArray
  }
}
