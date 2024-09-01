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
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func EchoString(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintf(c, "\t", strings.ToUpper(shout))

}

func handleConn3(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {

		echo(c, input.Text(), 1*time.Second)

	}
	c.Close()
}
func mustCopy3(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {

		log.Fatal(err)

	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8005")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy3(os.Stdout, conn)

}
