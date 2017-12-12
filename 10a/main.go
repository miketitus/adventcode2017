package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

const size = 5

var list circularList
var input = []int{3, 4}
var input1 = []int{
	83, 0, 193, 1, 254, 237, 187, 40, 88, 27, 2, 255, 149, 29, 42, 100,
}

type cNode struct {
	index int
	value int
	next  *cNode
	prev  *cNode
}

type circularList struct {
	zeroNode *cNode
	length   int
}

func (cl *circularList) addNode(cn *cNode) {
	if cl.zeroNode == nil {
		cl.zeroNode = cn
		cn.next = cn
		cn.prev = cn
		cn.index = 0
	} else {
		// append to "end" of list (before zeroNode)
		oldEnd := cl.zeroNode.prev
		oldEnd.next = cn
		cn.next = cl.zeroNode
		cn.prev = oldEnd
		cn.index = cn.prev.index + 1
		cl.zeroNode.prev = cn
	}
	cl.length++
}

func (cl *circularList) getNodeAt(index int) *cNode {
	node := cl.zeroNode
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}

func (cl *circularList) len() int {
	return cl.length
}

func (cl *circularList) swapSection(start, length int) {
	if length <= 1 {
		fmt.Printf("skipping length %d at start %d\n", length, start)
		return
	}
	left := cl.getNodeAt(start)
	right := cl.getNodeAt(start + length - 1)
	fmt.Printf("left:  %p  %v\n", left, left)
	fmt.Printf("right:  %p  %v\n", right, right)
	origLindex := left.index
	origLnext := left.next
	origLprev := left.prev
	origRindex := right.index
	origRnext := right.next
	origRprev := right.prev
	// update target nodes
	left.index = origRindex
	left.next = origRnext
	if length == 2 {
		left.prev = right
		right.next = left
	} else {
		left.prev = origRprev
		right.next = origLnext
	}
	right.index = origLindex
	right.prev = origLprev
	// update "exterior" nodes
	origLprev.next = right
	origRnext.prev = left
	// update "interior" nodes, if any
	if length > 2 {
		origLnext.prev = right
		origRprev.next = left
	}
	// update zeroNode, if necessary
	if left.index == 0 {
		cl.zeroNode = left
	} else if right.index == 0 {
		cl.zeroNode = right
	}
	fmt.Printf("left:  %p  %v\n", left, left)
	fmt.Printf("right:  %p  %v\n", right, right)
	if length > 3 {
		cl.swapSection(start+1, length-2)
	}
}

func (cl circularList) String() string {
	buf := bytes.NewBufferString("[")
	if cl.zeroNode == nil {
		// write nothing
	} else {
		buf.WriteRune(' ')
		orig := cl.zeroNode
		for start := cl.zeroNode; ; start = start.next {
			//buf.WriteString(strconv.Itoa(start.index))
			//buf.WriteRune(':')
			buf.WriteString(strconv.Itoa(start.value))
			buf.WriteRune(' ')
			if start.next == cl.zeroNode {
				break
			} else if start.next == orig {
				fmt.Println("Found infinite loop in list")
				os.Exit(1)
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
	// process input
	position := 0
	skipSize := 0
	for _, length := range input {
		list.swapSection(position, length)
		position += length + skipSize
		position = position % list.len()
		skipSize++
		fmt.Printf("pos:%d, ss:%d\n", position, skipSize)
		fmt.Printf("list:%v\n", list)
	}
}
