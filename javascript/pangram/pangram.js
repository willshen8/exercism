export const isPangram = sentence => {
  let letterMap = new Map()
  sentence.split('').map(letter => {
    if (letter.toLowerCase() >= 'a' && letter.toLowerCase() <= 'z') {
      letterMap.set(letter.toLowerCase(), true)
    }
  })
  if (letterMap.size == 26) return true
  return false
}