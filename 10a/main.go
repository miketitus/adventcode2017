package main

import (
	"bytes"
	"fmt"
	"strconv"
)

const size = 5

var list circularList
var input = []int{3}
var input1 = []int{
	83, 0, 193, 1, 254, 237, 187, 40, 88, 27, 2, 255, 149, 29, 42, 100,
}

type cNode struct {
	value int
	next  *cNode
	prev  *cNode
}

type circularList struct {
	zeroNode *cNode
}

func (cl *circularList) addNode(cn *cNode) {
	if cl.zeroNode == nil {
		cl.zeroNode = cn
		cn.next = cn
		cn.prev = cn
	} else {
		// append to "end" of list (before zeroNode)
		oldEnd := cl.zeroNode.prev
		oldEnd.next = cn
		cn.next = cl.zeroNode
		cn.prev = oldEnd
		cl.zeroNode.prev = cn
	}
}

func (cl *circularList) getNodeAt(index int) *cNode {
	node := cl.zeroNode
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}

func (cl *circularList) swapSection(start, length int) {
	if length <= 1 {
		fmt.Printf("skipping length %d at start %d", length, start)
		return
	}
	left := cl.getNodeAt(start)
	right := cl.getNodeAt(start + length - 1)
	fmt.Printf("left: %v, right: %v\n", left, right)
	origLnext := left.next
	origLprev := left.prev
	origRnext := right.next
	origRprev := right.prev
	// update target nodes
	left.next = origRnext
	left.prev = origRprev
	right.prev = origLprev
	right.next = origLnext
	// update "exterior" nodes
	origLprev.next = right
	origRnext.prev = left
	// update "interior" nodes, if any
	if length > 2 {
		origLnext.prev = right
		origRprev.next = left
	}
}

func (cl circularList) String() string {
	buf := bytes.NewBufferString("[")
	if cl.zeroNode == nil {
		// write nothing
	} else {
		buf.WriteRune(' ')
		for start := cl.zeroNode; ; start = start.next {
			buf.WriteString(strconv.Itoa(start.value))
			buf.WriteRune(' ')
			if start.next == cl.zeroNode {
				break
			}
		}
	}
	buf.WriteString("]")
	return buf.String()
}

func main() {
	// initialize list
	list = circularList{}
	for i := 0; i < size; i++ {
		node := cNode{value: i}
		list.addNode(&node)
	}
	fmt.Println(list)
	list.swapSection(0, 3)
	fmt.Println(list)
	// process input
	/*
		position := 0
		skipSize := 0
		for _, length := range input {
			reverseSection(position, length)
			position += length + skipSize
			skipSize++
			fmt.Printf("pos:%d, ss:%d\n", position, skipSize)
			fmt.Printf("list:%v\n", list)
		}
	*/
}
