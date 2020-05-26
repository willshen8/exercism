export const gigasecond = date => {
  const TeraMilliseconds = 1e12
  return new Date((date.getTime() + TeraMilliseconds))
}