export const commands = secret => {
  const secretHandshake = ['wink', 'double blink', 'close your eyes', 'jump']
  const convertToBinaryString = num => (num).toString(2)
  const reverseBitPosition = 4

  let result = []
  convertToBinaryString(secret).split("").reverse().map((digit, index) => {
    if (digit === '1' && secretHandshake[index]) result.push(secretHandshake[index])
    if (digit === '1' && index === reverseBitPosition) result.reverse()
  })
  return result
}
