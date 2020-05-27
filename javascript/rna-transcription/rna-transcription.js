export const toRna = dna => {
  let result = ''
  for (let strand of dna) {
    result += DnaToRna(strand)
  }
  return result
}

const DnaToRna = strand => {
  switch (strand) {
    case 'G':
      return 'C'
    case 'C':
      return 'G'
    case 'T':
      return 'A'
    case 'A':
      return 'U'
    default:
      return ''
  }
}