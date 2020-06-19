package listops

type IntList []int
type binFunc func(int, int) int
type unaryFunc func(x int) int
type predFunc func(n int) bool

// Append will append the list appendList onto list l and returns a new list
func (l IntList) Append(appendList IntList) IntList {
	resultList := make(IntList, len(l)+len(appendList))
	for i := 0; i < len(l); i++ {
		resultList[i] = l[i]
	}
	for j := 0; j < len(appendList); j++ {
		resultList[j+len(l)] = appendList[j]
	}
	return resultList
}

func (l IntList) Concat(lists []IntList) IntList {
	var resultList = l
	for _, list := range lists {
		resultList = append(resultList, list...)
	}
	return resultList
}

// Filter removes items in list that do not satisfy the filterfunc
func (l IntList) Filter(filterFunc predFunc) IntList {
	resultList := make(IntList, 0)
	for i := 0; i < len(l); i++ {
		if filterFunc(l[i]) {
			resultList = append(resultList, l[i])
		}
	}
	return resultList
}

func (l IntList) Length() int {
	var length int
	for i, _ := range l {
		length = i + 1
	}
	return length
}

func (l IntList) Map(mapFunc unaryFunc) IntList {
	resultList := make(IntList, len(l))
	for i := 0; i < len(l); i++ {
		resultList[i] = mapFunc(l[i])
	}
	return resultList
}

func (l IntList) Foldl(foldFunc binFunc, initial int) int {
	var accumulator = initial
	for i := 0; i < len(l); i++ {
		accumulator = foldFunc(accumulator, l[i])
	}
	return accumulator
}

func (l IntList) Foldr(foldFunc binFunc, initial int) int {
	var accumulator = initial
	for i := len(l) - 1; i > -1; i-- {
		accumulator = foldFunc(l[i], accumulator)
	}
	return accumulator
}

func (l IntList) Reverse() IntList {
	resultList := make(IntList, len(l))
	var resultIndex int
	for i := len(l) - 1; i > -1; i-- {
		resultList[resultIndex] = l[i]
		resultIndex++
	}
	return resultList
}
