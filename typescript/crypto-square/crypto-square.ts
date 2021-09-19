export class Crypto {
  private _plainText: unknown
  constructor(plainText: unknown) {
    this._plainText = plainText
  }

  get ciphertext(): unknown {
    let cipherText = ''
    if (this._plainText && typeof this._plainText === "string") {
      cipherText = this._plainText.replace(/[^a-zA-Z0-9+]/g, "").toLowerCase();
      const columns = Math.ceil(Math.sqrt(cipherText.length))
      let rows = Math.floor(Math.sqrt(cipherText.length))
      if(rows * columns < cipherText.length) {
        rows++
      }
      //add padding to the end of the text
      if (cipherText.length < rows * columns) {
        cipherText = cipherText.padEnd(rows*columns, ' ')
      }

      let normalisedTextMatrix = Array(rows)
      for(let i=0; i<rows; i++) {
        normalisedTextMatrix[i] = cipherText.substring(i*columns, (i+1)*columns)
      }
      let cipherTextMatrix: Array<string> = []
      for(let i=0; i<columns; i++) {
        let cipher = ''
        for(let j=0; j<rows; j++) {
          cipher += normalisedTextMatrix[j][i]
        }
        cipherTextMatrix[i] = cipher
      }
     cipherText = cipherTextMatrix.join(' ')
    }
    return cipherText
  }
}
