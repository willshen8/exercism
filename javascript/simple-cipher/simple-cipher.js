const ASCII_a = 97
const ASCII_z = 122
const MIN_NUM = 100
const generateRandomLetter = () => String.fromCharCode(Math.floor(Math.random() * (ASCII_z - ASCII_a) + ASCII_a))

export class Cipher {
  constructor(key) {
    if (!key || key == undefined) {
      for (let i=0; i<MIN_NUM; i++) {
        this._key += generateRandomLetter()
      }
      return
    }
    this._key = key
  }

  encode(input) {
    return input.split("").map((char, index) => String.fromCharCode(char.charCodeAt(0) - ASCII_a + this._key[index].charCodeAt(0))).join('')
  }

  decode(encodedMessage) {
    return encodedMessage.split('').map((char, index) => String.fromCharCode(char.charCodeAt(0) - this._key[index].charCodeAt(0) + ASCII_a)).join('')
  }

  generateRandomLetter() {

  }

  get key() {
    return this._key
  }
}
