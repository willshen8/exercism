export const square = squareNum => {
  if (squareNum < 1 || squareNum > 64) throw new Error('square must be between 1 and 64')
  return BigInt(Math.pow(2, squareNum - 1))
}

export const total = () => BigInt(2 ** 64) - BigInt(1)