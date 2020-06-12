const numberOfLetters = 26

export const isPangram = sentence => {
  let letterSet = new Set()
  sentence.split('').map(letter => {
    if (letter.toLowerCase() >= 'a' && letter.toLowerCase() <= 'z') {
      letterSet.add(letter.toLowerCase())
    }
  })
  return letterSet.size == numberOfLetters
}