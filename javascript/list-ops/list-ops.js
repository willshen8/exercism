export class List {
    constructor(values = []) {
        this.values = values
    }

    // append adds a new list to existing list and return a new list
    append(list) {
        return new List([...this.values, ...list.values])
    }

    // concatenate (given a series of lists, combine all items in all lists into one flattened list)
    concat(listArray) {
        let newList = Object.assign({}, this)
        for (let i = 0; listArray.values[i] !== undefined; i++) {
            newList.values = [...newList.values, ...listArray.values[i].values]
        }
        return newList
    }

    // given a predicate and a list, return the list of all items for which predicate(item) is True)
    filter(predicate) {
        let newList = new List()
        for (let i = 0; i < this.length(); i++) {
            if (predicate(this.values[i]))
                newList.values = [...newList.values, this.values[i]]
        }
        return newList
    }

    // given a function and a list, return the list of the results of applying function(item) on all items)
    map(func) {
        let newList = new List()
        for (let i = 0; i < this.length(); i++) {
            newList.values = [...newList.values, func(this.values[i])]
        }
        return newList
    }

    // length returns the size of the list
    length() {
        let count = 0
        for (let i = 0; this.values[i] !== undefined; i++) {
            count++
        }
        return count
    }

    // foldl given a function, a list, and initial accumulator, 
    // fold (reduce) each item into the accumulator from the left using function(accumulator, item))
    foldl(func, accumulator) {
        for (let i = 0; i < this.length(); i++) {
            accumulator = func(accumulator, this.values[i])
        }
        return accumulator
    }

    // foldr (given a function, a list, and an initial accumulator, 
    // fold (reduce) each item into the accumulator from the right using function(item, accumulator))
    foldr(func, accumulator) {
        for (let i = this.length() - 1; i >= 0; i--) {
            accumulator = func(accumulator, this.values[i])
        }
        return accumulator
    }

    // reverse (given a list, return a list with all the original items, but in reversed order)
    reverse() {
        let newList = new List()
        for (let i = this.length() - 1; i >= 0; i--) {
            newList.values = [...newList.values, this.values[i]]
        }
        return newList
    }
}