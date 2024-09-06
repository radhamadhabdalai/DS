/*
 * Copyright (c) 2020 Radhamadhab Dalai
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


package concurrent

import (
	"fmt"
	"sync"
)

func orderDress(wg *sync.WaitGroup) {
	fmt.Println("Order a dress")
	wg.Done()
}
func payCreditCardBills(wg *sync.WaitGroup) {

	fmt.Println(" pay credit cards")
	wg.Done()
}

var listOfTasks = []func(group *sync.WaitGroup){ // list of functions taking argument as sync.WaitGroup of
	orderDress, payCreditCardBills, // comma separated
}

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(listOfTasks))

	for _, task := range listOfTasks {

		task(&waitGroup)
	}

	waitGroup.Wait()
}