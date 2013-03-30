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
package stack_test

import (
	"fmt"
	"stack"
)

// Simple usage example that inserts the numbers 1, 2, 3 into a stack and then
// removes them one by one, printing them to the standard output.
func Example_usage() {
	// Create a stack and push some data in
	s := stack.New()
	for i := 0; i < 3; i++ {
		s.Push(i)
	}
	// Pop out the stack contents and display them
	for !s.Empty() {
		fmt.Println(s.Pop())
	}
	// Output:
	// 2
	// 1
	// 0
}
