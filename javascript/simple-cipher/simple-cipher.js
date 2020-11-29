const FIRST_ALPHABET_CODE = 97
const LAST_ALPHABET_CODE = 122
const KEY_LENGTH = 100
const NUM_OF_ENG_ALPHABET = 26
const generateRandomLetter = () => String.fromCharCode(Math.floor(Math.random() * (LAST_ALPHABET_CODE - FIRST_ALPHABET_CODE) + FIRST_ALPHABET_CODE))

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
      const finalPosition = char.charCodeAt(0) + this._key[index % this._key.length].charCodeAt(0) - FIRST_ALPHABET_CODE
      if (finalPosition > LAST_ALPHABET_CODE) { // wrap if go beyond Z
        return String.fromCharCode(finalPosition - NUM_OF_ENG_ALPHABET)
      }
      return String.fromCharCode(finalPosition)
    }).join('')
  }

  decode(encodedMessage) {
    return encodedMessage.split('').map((char, index) => {
      const finalPosition = char.charCodeAt(0) - this._key[index % this._key.length].charCodeAt(0) + FIRST_ALPHABET_CODE 
      if (finalPosition < FIRST_ALPHABET_CODE) { // wrap if go beyond A
        return String.fromCharCode(finalPosition + NUM_OF_ENG_ALPHABET)
      }
      return String.fromCharCode(finalPosition)
    }).join('')
  }

  get key() {
    return this._key
  }
}
