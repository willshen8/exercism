export const rows = (number) => {
  const pascalTriangles = []
  if (number <= 0) return pascalTriangles
  pascalTriangles.push([1])
  for (let i = 2; i < number + 1; i++) {
    const pascalRow = []
    pascalRow[0] = 1
    pascalRow[i - 1] = 1
    for (let j = 1; j < i - 1; j++) {
      pascalRow[j] = pascalTriangles[i - 2][j - 1] + pascalTriangles[i - 2][j]
    }
    pascalTriangles.push(pascalRow)
  }
  return pascalTriangles
}