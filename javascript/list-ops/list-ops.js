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
        return this.append(new List(listArray.foldl((acc, item) => [...acc, ...item.values], [])))
    }

    // given a predicate and a list, return the list of all items for which predicate(item) is True)
    filter(predicate) {
        return new List(this.foldl((acc, element) => predicate(element) ? [...acc, element] : [...acc], []))
    }

    // given a function and a list, return the list of the results of applying function(item) on all items)
    map(func) {
        return new List(this.foldl((acc, element) => [...acc, func(element)], []))
    }

    // length returns the size of the list
    length() {
        return this.foldl(length => ++length, 0)
    }

    // fold is a generic helper function that works the same as foldl that accepts array as input
    fold(func, acc, arr) {
        const [first, ...rest] = arr
        if (first === undefined)
            return acc
        acc = func(acc, first)
        return this.fold(func, acc, rest)
    }

    // foldl (reduce) each item into the accumulator from the left using function(accumulator, item))
    foldl(func, accumulator) {
        return this.fold(func, accumulator, this.values)
    }

    // foldr (given a function, a list, and an initial accumulator, 
    // fold (reduce) each item into the accumulator from the right using function(item, accumulator))
    foldr(func, accumulator) {
        return this.reverse().foldl(func, accumulator)
    }

    // reverse (given a list, return a list with all the original items, but in reversed order)
    reverse() {
        return new List(this.foldl((acc, element) => [element, ...acc], []))
    }
}