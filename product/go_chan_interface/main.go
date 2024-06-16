package main

import (
	"fmt"
	"strings"
	"sync"
)

type Processor interface {
	Process(data string)
}

type UppercaseProcessor struct {
	bisiv chan string
}

func (up UppercaseProcessor) Process(data string) {
	up.bisiv <- strings.ToUpper(data)
}

type ReverseProcessor struct {
	bisiv chan string
}

func (rp ReverseProcessor) Process(data string) {
	reverseData := []rune(data)
	for i, j := 0, len(reverseData)-1; i < j; i, j = i+1, j-1 {
		reverseData[i], reverseData[j] = reverseData[j], reverseData[i]
	}
	rp.bisiv <- string(reverseData)
}

func main() {
	var wg sync.WaitGroup

	UPChan := make(chan string)
	RPChan := make(chan string)

	up := UppercaseProcessor{bisiv: UPChan}
	rp := ReverseProcessor{bisiv: RPChan}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, line := range ListOfLines {
			up.Process(line)
		}
		close(UPChan)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, line := range ListOfLines {
			rp.Process(line)
		}
		close(RPChan)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for line := range UPChan {
			UppercaseLines = append(UppercaseLines, line)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for line := range RPChan {
			ReverseLines = append(ReverseLines, line)
		}
	}()
	wg.Wait()

	fmt.Println("Caps Lines:")
	for _, line := range UppercaseLines {
		fmt.Printf("%s\n", line)
	}
	fmt.Println()

	fmt.Println("Reverse Lines:")
	for _, line := range ReverseLines {
		fmt.Printf("%s\n", line)
	}
}
