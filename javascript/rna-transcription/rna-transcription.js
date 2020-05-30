const transcription = {
  G: 'C',
  C: 'G',
  T: 'A',
  T: 'A',
  A: 'U'
}

export const toRna = dna =>
    dna.split('').map(nucleotides => transcription[nucleotides]).join('')
