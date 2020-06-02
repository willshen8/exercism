export const compute = (dna1, dna2) => {
  if (dna1.length == 0 && dna2.length == 0) return 0
  if (dna1.length == 0) throw new Error('left strand must not be empty')
  if (dna2.length == 0) throw new Error('right strand must not be empty')
  if (dna1.length != dna2.length) throw new Error('left and right strands must be of equal length')
  let result = 0
  for (let i = 0; i < dna1.length; i++) {
    if (dna1[i] != dna2[i]) result++
  }
  return result
}