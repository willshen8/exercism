const ASCII_a = 97
const ASCII_z = 122
const KEY_LENGTH = 100
const NUM_OF_ENG_ALPHABET = 26
const generateRandomLetter = () => String.fromCharCode(Math.floor(Math.random() * (ASCII_z - ASCII_a) + ASCII_a))

export class Cipher {
  constructor(key) {
    if (key == undefined) {
      for (let i=0; i<KEY_LENGTH; i++) {
        this._key += generateRandomLetter()
      }
      return
    }
    this._key = key
  }

  encode(input) {
    return input.split("").map((char, index) => {
      if (char.charCodeAt(0) + this._key[index % this._key.length].charCodeAt(0) - ASCII_a > ASCII_z) {
        return String.fromCharCode(char.charCodeAt(0) + this._key[index % this._key.length].charCodeAt(0) - ASCII_a - NUM_OF_ENG_ALPHABET)
      }
      return String.fromCharCode(char.charCodeAt(0) - ASCII_a + this._key[index % this._key.length].charCodeAt(0))
    }).join('')
  }

  decode(encodedMessage) {
    return encodedMessage.split('').map((char, index) => {
      if (char.charCodeAt(0) - this._key[index % this._key.length].charCodeAt(0) + ASCII_a < ASCII_a) {
        return String.fromCharCode(char.charCodeAt(0) - this._key[index % this._key.length].charCodeAt(0) + ASCII_a + NUM_OF_ENG_ALPHABET)
      }
      return String.fromCharCode(char.charCodeAt(0) - this._key[index % this._key.length].charCodeAt(0) + ASCII_a)
    }).join('')
  }

  get key() {
    return this._key
  }
}
