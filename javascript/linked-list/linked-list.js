export class LinkedList {
  constructor() {
    this.head = null 
    this.tail = null
  }

  // push insert an element at the back of the list
  push(nodeValue) {
    const newNode = new Node(nodeValue)
    const oldTail = this.tail
    if (!this.head) {
      this.head = newNode 
      this.tail = newNode
    }
    else {
      oldTail.next = newNode
      newNode.prev = oldTail 
      this.tail = newNode
    }
  }

  // Pop removes the element at back
  pop() {
    if (!this.tail) return 
    const popNode = this.tail
      if (this.head !== this.tail) {
        const newTail = this.tail.prev
        newTail.next = null 
        this.tail = newTail
      }
    else {
      this.head = null 
      this.tail = null
    }
    return popNode.data
  }

  // shift removes the element from the beginning of the list
  shift() {
    const shiftedNode = this.head
    const newHead = this.head.next
    if (!this.head) return 
    if (!newHead) this.tail = newHead 
    this.head = newHead
    return shiftedNode.data
  }

  // unshift inserts a new value at front
  unshift(nodeValue) {
    const newNode = new Node(nodeValue)
    const oldHead = this.head
    if (!this.head) {
      this.head = newNode 
      this.tail = newNode
    }
    else {
      oldHead.prev = newNode
      newNode.next = oldHead 
      this.head = newNode
    }
  }

  // delete the first occurrence of a specified value
  delete(deleteValue) {
    const deleteNode = this.findNodeByValue(deleteValue)
    if (deleteNode) this.deleteNode(deleteNode)
  }

  // count output the number of nodes in the linked list
  count() {
    let count = 0
    if (!this.head) return 0
    let currentNode = this.head
    while (currentNode !== null) {
      count++
      currentNode = currentNode.next
    }
    return count
  }

  // findNodeByValue returns the first node with the specific value
  findNodeByValue(value) {
    let searchNode = null
    if (this.head) {
      searchNode = this.head
      while (searchNode.data !== value) {
        if (searchNode.next === null) {
            return
        }
        searchNode = searchNode.next
      }
      return searchNode
    }
  }

  // deleteNode removes deleteNode from the likedList
  deleteNode(deleteNode){
    let nextNode = null
    let prevNode = null
    if (deleteNode == this.head) {
      this.head = deleteNode.next 
      this.tail = deleteNode.next
    } else if (deleteNode == this.tail) {
      prevNode = deleteNode.prev
      prevNode.next = null 
      this.tail = prevNode
    } else {
      nextNode = deleteNode.next
      prevNode = deleteNode.prev
      prevNode.next = nextNode
      nextNode.prev = prevNode
    }
  }
}

class Node {
  constructor(data, prev = null, next = null) {
    this.data = data 
    this.prev = prev 
    this.next = next
  }
}