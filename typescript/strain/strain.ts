export function keep<T>(inputList: T[], predicate: (item: T) => boolean): T[] {
  return inputList.filter(listItem => predicate(listItem))
}

export function discard<T>(inputList: T[], predicate: (item: T) => boolean): T[] {
  return inputList.filter(listItem => !predicate(listItem))
}
