export const keep = <T>(inputList: T[], predicate: (item: T) => boolean): T[] => {
  let output: T[] = []
  for (let listItem of inputList) {
    if (predicate(listItem)) {
      output.push(listItem)
    }
  }
  return output
}

export const discard = <T>(inputList: T[], predicate: (item: T) => boolean): T[] => {
  return keep(inputList, function (item: T) { return !predicate(item) })
}
