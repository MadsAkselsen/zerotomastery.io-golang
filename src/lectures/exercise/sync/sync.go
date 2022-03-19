//--Summary:
//  Create a program that can read text from standard input and count the
//  number of letters present in the input.
//
//--Requirements:
//* Count the total number of letters in any chosen input
//* The input must be supplied from standard input
//* Input analysis must occur per-word, and each word must be analyzed
//  within a goroutine
//* When the program finishes, display the total number of letters counted
//
//--Notes:
//* Use CTRL+D (Mac/Linux) or CTRL+Z (Windows) to signal EOF, if manually
//  entering data
//* Use `cat FILE | go run ./exercise/sync` to analyze a file
//* Use any synchronization techniques to implement the program:
//  - Channels / mutexes / wait groups

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	// "strconv"
	// "time"
)

func getText(rd bufio.Reader) string {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, &rd)
	if err != nil {
		//* Report any errors to the terminal
		fmt.Println("Error:", err)
	}
		
	return buf.String()
}

func main() {
	var wg sync.WaitGroup
	sum := 0
	textFile := "num1.txt"
	file, err := os.Open(textFile)
	if err != nil {
		//* Report any errors to the terminal
		fmt.Println("Error:", err)
		return
	}

	rd := bufio.NewReader(file)
	text := getText(*rd)
	words := strings.Fields(text)
	fmt.Println(words)

	for _, word := range words {
		wg.Add(1)
		
		go func() {
			runesAmount := 0
			for i:=0; i< len(word); i++ {
				runesAmount ++
			}
			sum += runesAmount
			wg.Done()
		}()
		

		
	}
	wg.Wait()
	fmt.Println(sum)

}
