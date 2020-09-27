export const countWords = words => {
  let wordMap = new Map()
  const wordsArray = words.split(/[, ]/)
  wordsArray.map(word => {
    let parsedWord = word.trim().toLowerCase().replace(/(^\W*)|(\W*$)/g, '')
    if (parsedWord === '')
      return
    if (!wordMap.has(parsedWord))
      wordMap.set(parsedWord, 1)
    else
      wordMap.set(parsedWord, wordMap.get(parsedWord) + 1)
  })
  const wordFreq = {}
  for (const [key, value] of wordMap.entries()) {
    wordFreq[key] = value
  }
  return wordFreq
}