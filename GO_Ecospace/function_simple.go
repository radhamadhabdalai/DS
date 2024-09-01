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
	"sync"
	"time"
)

type Job struct {
	i    int
	max  int
	text string
}

func outText(j *Job, goGroup *sync.WaitGroup) {
	//text1 string
	if j.i < j.max {
		var text1 = stringToText(j.text)
		time.Sleep(1 * time.Millisecond)
		fmt.Println(text1)
	}
	goGroup.Done()

}

func stringToText(text string) string {
	return text
}

func main() {
	//myText := stringToText("hello world")
	goGroup := new(sync.WaitGroup)

	hello := new(Job)
	hello.i = 100
	hello.text = "Hello"
	hello.max = 200
	go outText(hello, goGroup)

	hello2 := new(Job)
	hello2.i = 100
	hello2.text = "World"
	hello2.max = 200
	go outText(hello2, goGroup)

	goGroup.Add(2)
	goGroup.Wait()

}
