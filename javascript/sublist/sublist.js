export class List {
  constructor(arr) {
    this._list = arr
  }

  get getList() {return this._list}

  compare(list) {
    if (this._list === undefined && list.getList === undefined) {
      return 'EQUAL'
    }
    if (this._list === undefined && list.getList.length > 0) {
      return 'SUBLIST'
    }
    if (this._list.length > 0 && list.getList === undefined) {
      return 'SUPERLIST'
    }

    if(this._list.length === list.getList.length) {
      for (let i=0; i<this._list.length; i++){
        if(this._list[i] !== list.getList[i]) return 'UNEQUAL'
      }
      return 'EQUAL'
    }

    if(this._list.length < list.getList.length && isSubArray(list.getList, this._list)){
      return 'SUBLIST'
    } else if (this._list.length > list.getList.length && isSubArray(this._list, list.getList)) {
      return 'SUPERLIST'
    }

    return 'UNEQUAL'
  }
}

// returns true if arrB is a subArray of arrA
const isSubArray = (arrA, arrB) => {
  // find the index of the first element of A
  let startIndex = 0;
  for(let i=0; i < arrA.length; i++) {
    if (arrA[i] === arrB[0]) {
      startIndex = i
    }
  }

  for (let i=startIndex, j=0; i < arrB.length; i++, j++) {
    if(arrA[i] !== arrB[j]) {
      return false
    }
  }
  return true
}