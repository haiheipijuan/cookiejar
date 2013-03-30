// CookieJar - A contestant's algorithm toolbox
// Copyright 2013 Peter Szilagyi. All rights reserved.
//
// CookieJar is dual licensed: you can redistribute it and/or modify it under
// the terms of the GNU General Public License as published by the Free Software
// Foundation, either version 3 of the License, or (at your option) any later
// version.
//
// The toolbox is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for
// more details.
//
// Alternatively, the CookieJar toolbox may be used in accordance with the terms
// and conditions contained in a signed written agreement between you and the
// author(s).
//
// Author: peterke@gmail.com (Peter Szilagyi)

// Package queue implements a FIFO (first in first out) data structure supporting
// arbitrary types (even a mixture).
//
// Internally it uses a dynamically growing circular slice of blocks, resulting
// in faster resizes than a simple dynamic array/slice would allow.
package queue

// The size of a block of data
const blockSize = 4096

// First in, first out data structure.
type Queue struct {
	tailIdx int
	headIdx int
	tailOff int
	headOff int

	blocks [][]interface{}
	head   []interface{}
	tail   []interface{}
}

// Creates a new, empty queue.
func New() *Queue {
	result := new(Queue)
	result.blocks = [][]interface{}{make([]interface{}, blockSize)}
	result.head = result.blocks[0]
	result.tail = result.blocks[0]
	return result
}

// Pushes a new element into the queue, expanding it if necessary.
func (q *Queue) Push(data interface{}) {
	q.tail[q.tailOff] = data
	q.tailOff++
	if q.tailOff == blockSize {
		q.tailOff = 0
		q.tailIdx = (q.tailIdx + 1) % len(q.blocks)

		// If we wrapped over to the end, insert a new block and update indices
		if q.tailIdx == q.headIdx {
			buffer := make([][]interface{}, len(q.blocks)+1)
			copy(buffer[:q.tailIdx], q.blocks[:q.tailIdx])
			buffer[q.tailIdx] = make([]interface{}, blockSize)
			copy(buffer[q.tailIdx+1:], q.blocks[q.tailIdx:])
			q.blocks = buffer
			q.headIdx++
			q.head = q.blocks[q.headIdx]
		}
		q.tail = q.blocks[q.tailIdx]
	}
}

// Pops out an element from the queue. Note, no bounds checking are done.
func (q *Queue) Pop() interface{} {
	res := q.head[q.headOff]
	q.headOff++
	if q.headOff == blockSize {
		q.headOff = 0
		q.headIdx = (q.headIdx + 1) % len(q.blocks)
		q.head = q.blocks[q.headIdx]
	}
	return res
}

// Checks whether the queue is empty.
func (q *Queue) Empty() bool {
	return q.headIdx == q.tailIdx && q.headOff == q.tailOff
}

// Clears out the contents of the queue.
func (q *Queue) Reset() {
	q.headIdx = 0
	q.tailIdx = 0
	q.headOff = 0
	q.tailOff = 0
	q.blocks = [][]interface{}{make([]interface{}, blockSize)}
	q.head = q.blocks[0]
	q.tail = q.blocks[0]
}
