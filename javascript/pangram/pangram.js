const lettersInEnglishAlphabet = 26

export const isPangram = sentence => {
  let alphabetSet = new Set(sentence.toLowerCase().replace(/[^A-Za-z]/g, ''))
  return alphabetSet.size == lettersInEnglishAlphabet
}