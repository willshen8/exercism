const FIRST_ALPHABET_CODE = 97
const LAST_ALPHABET_CODE = 122
const KEY_LENGTH = 100
const NUM_OF_ENG_ALPHABET = 26
const RIGHT = 1
const LEFT = 0
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
    return this.shift(input, RIGHT)
  }

  decode(encodedMessage) {
    return this.shift(encodedMessage, LEFT)
  }

  shift(input, direction) {
    const rightShiftChar = (char, index) => char.charCodeAt(0) + this._key[index % this._key.length].charCodeAt(0) - FIRST_ALPHABET_CODE
    const leftShiftChar = (char, index) => char.charCodeAt(0) - this._key[index % this._key.length].charCodeAt(0) + FIRST_ALPHABET_CODE
    const shiftChar = (char, index) => {
      const finalPosition = (direction === RIGHT) ? rightShiftChar(char, index) : leftShiftChar(char, index)
        if (direction === RIGHT && finalPosition > LAST_ALPHABET_CODE) {
            return String.fromCharCode(finalPosition - NUM_OF_ENG_ALPHABET)
        }
        if (finalPosition < FIRST_ALPHABET_CODE) { // left
            return String.fromCharCode(finalPosition + NUM_OF_ENG_ALPHABET)
        }
      return String.fromCharCode(finalPosition)
    }
    return input.split("").map(shiftChar).join('')
  }

  get key() {
    return this._key
  }
}
