export class LinkedList {
  constructor() {
    this.head = null
    this.tail = null
  }

  // push insert an element at the back of the list
  push(nodeValue) {
    this.insertNode(nodeValue, this.count)
  }

  // Pop removes the element at back
  pop() {
    if (!this.tail) return
    const poppedData = this.tail.data
    this.deleteNode(this.tail)
    return poppedData
  }

  // shift removes the element from the beginning of the list
  shift() {
    if (!this.head) return
    const shiftedData = this.head.data
    this.deleteNode(this.head)
    return shiftedData
  }

  // unshift inserts a new value at front
  unshift(nodeValue) {
    this.insertNode(nodeValue, 1)
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

  // deleteNode removes a node from the likedList
  deleteNode(node) {
    if (node == this.head) {
      this.head = node.next
    } else if (node == this.tail) {
      this.tail.prev.next = null
      this.tail = this.tail.prev
    } else {
      node.prev.next = node.next
      node.next.prev = node.prev
    }
  }

  // insertNode adds a node insert the linked list
  insertNode(nodeValue, index) {
    const newNode = new Node(nodeValue)
    if (!this.head) {
      this.head = newNode
      this.tail = newNode
    } else if (index == 1) {
      this.head.prev = newNode
      newNode.next = this.head
      this.head = newNode
    } else if (index == this.count) {
      newNode.prev = this.tail
      this.tail.next = newNode
      this.tail = newNode
    } else {
      let currNode = this.head
      let incr = 1
      while (currNode.next !== null && incr < index) {
        currNode = currNode.next
        incr++
      }
      newNode.next = currNode.next
      currNode.next = newNode
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