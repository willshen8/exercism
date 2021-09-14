export const keep = <T>(inputList: T[], predicate: (item: T) => boolean): T[] => {
  let output: T[] = []
  inputList.forEach(item => {
    if (predicate(item)) {
      output.push(item)
    }
  })
  return output
}

export function discard<T>(inputList: T[], predicate: (item: T) => boolean): T[] {
  let output: T[] = []
  inputList.forEach(item => {
    if (!predicate(item)) {
      output.push(item)
    }
  })
  return output
}
