const lettersInEnglishAlphabet = 26

export const isPangram = sentence =>
    new Set(sentence.toLowerCase().replace(/[^A-Za-z]/g, '')).size ==
    lettersInEnglishAlphabet
