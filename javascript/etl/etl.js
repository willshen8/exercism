export const transform = input => {
  let result = new Object()
  Object.entries(input).forEach(item => {
    let resultValue = item[0]
    let resultKeys = item[1]
    resultKeys.map(item => result[item.toLowerCase()] = Number(resultValue))
  })
  return result
}