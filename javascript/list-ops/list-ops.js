export class List {
  constructor(values = []) {
    this.values = values
  }

  // append adds a new list to existing list and return a new list
  append(list) {
    let newList = Object.assign({}, this)
    list.values.forEach(value => newList.values.push(value))
    return newList
  }

  // concatenate (given a series of lists, combine all items in all lists into one flattened list)
  concat(listArray) {
    let newList = Object.assign({}, this)
    listArray.values.forEach(list => list.values.forEach(item => newList.values.push(item)))
    return newList
  }

  // given a predicate and a list, return the list of all items for which predicate(item) is True)
  filter(predicate) {
    let newList = new List()
    this.values.forEach(value =>  {
      if (predicate(value)) 
        newList.values.push(value)
      })
    return newList
  }

  // given a function and a list, return the list of the results of applying function(item) on all items)
  map(func) {
    let newList = new List()
    this.values.forEach(value => newList.values.push(func(value)))
    return newList
  }

  length() {
    let count = 0
    this.values.forEach(() => count++)
    return count
  }

  // foldl given a function, a list, and initial accumulator, 
  // fold (reduce) each item into the accumulator from the left using function(accumulator, item))
  foldl(func, accumulator) {
    this.values.forEach(value => accumulator = func(accumulator, value))
    return accumulator
  }

  // foldr (given a function, a list, and an initial accumulator, 
  // fold (reduce) each item into the accumulator from the right using function(item, accumulator))
  foldr(func, accumulator) {
    for (let i=this.values.length-1; i>=0; i--){
      accumulator = func(accumulator, this.values[i])
    }  
    return accumulator
  }

  // reverse (given a list, return a list with all the original items, but in reversed order)
  reverse() {
    let newList = new List()
    // this.values.forEach(() => newList.values.push(popList.values.pop()))
    for (let i=this.values.length-1; i>=0; i--){
      newList.values.push(this.values[i])
    }
    return newList  
  }
}