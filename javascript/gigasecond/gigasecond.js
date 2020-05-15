export const gigasecond = date => {
  const newDate = new Date()
  newDate.setTime(date.getTime() + 1000000000000)
  newDate.toISOString()
  return newDate
}