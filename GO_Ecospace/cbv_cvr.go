/*
 * Copyright (c) 2015 Radhamadhab Dalai
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */

package main

import (
	"fmt"
	"unsafe"
)

func callByValue(a [10]int, name byte) {
	for i, elem := range a {
		fmt.Printf("%c[%d] in %d", name, i, elem)
	}
}

func callByReference(b *[10]int, name byte) {
	for i := 0; i < len(b); i++ {

		fmt.Printf("%c[%d] is %d", name, i, b[i])
		fmt.Println("")
	}
}

func callByReference2D(b *[9]int, name byte) {
	for i := 0; i < len(b); i++ {

		fmt.Printf("%c[%d] is %d", name, i, b[i])
		fmt.Println("")
	}
}

func main() {

	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	callByReference(&a, 'a')

	b := new([10]int)
	b[1] = 1
	b[2] = 1

	callByValue(*b, 'b')

	var c [3][3]int
	c[1][1] = 1

	callByReference2D((*[9]int)(unsafe.Pointer(&c)), 'c')

}
