export const countWords = words => {
  const result = {}
  words.split(/[, ]/).map(word => {
    let parsedWord = word.trim().toLowerCase().replace(/(^\W*)|(\W*$)/g, '')
    if (parsedWord === '')
      return
    if (result[parsedWord])
      result[parsedWord]++
    else
      result[parsedWord] = 1
  })
  return result
}