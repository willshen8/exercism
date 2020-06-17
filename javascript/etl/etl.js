export const transform = input => {
  let result = new Object()
  Object.entries(input).forEach(([key, values]) => {
    values.map(item => result[item.toLowerCase()] = Number(key))
  })
  return result
}