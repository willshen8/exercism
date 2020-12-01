const FIRST_ALPHABET_CODE = 'a'.charCodeAt(0)
const LAST_ALPHABET_CODE = 'z'.charCodeAt(0)
const KEY_LENGTH = 100
const NUM_OF_ENG_ALPHABET = 26
const RIGHT = 1
const LEFT = 0
const generateRandomLetter = () => String.fromCharCode(Math.floor(Math.random() * (LAST_ALPHABET_CODE - FIRST_ALPHABET_CODE + 1) + FIRST_ALPHABET_CODE))

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
    const shiftChar = (char, index) => {
      const charIndex = char.charCodeAt(0) - FIRST_ALPHABET_CODE
      const shiftAmount = this._key[index % this._key.length].charCodeAt(0) - FIRST_ALPHABET_CODE
      const shiftedTo = (direction === RIGHT) ? charIndex + shiftAmount : charIndex - shiftAmount
      const finalPosition = (shiftedTo + NUM_OF_ENG_ALPHABET) % NUM_OF_ENG_ALPHABET + FIRST_ALPHABET_CODE
      return String.fromCharCode(finalPosition)
    }
    return input.split("").map(shiftChar).join('')
  }

  get key() {
    return this._key
  }
}
