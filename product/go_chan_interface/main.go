package main

import (
	"fmt"
	"github.com/SaYaku64/hillel/packages/bufioFiles"
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

func savePlase(ch chan string, wg *sync.WaitGroup, slise1 *[]string) {
	defer wg.Done()
	for line := range ch {
		*slise1 = append(*slise1, line)
	}
}

func savePlasePlease(ch chan string, wg *sync.WaitGroup, both Processor, ListOfLines []string) {
	defer wg.Done()
	for _, line := range ListOfLines {
		both.Process(line)
	}
	close(ch)
}

func main() {
	var wg sync.WaitGroup

	ListOfLines, err := bufiofiles.ReadLines("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var UppercaseLines []string
	var ReverseLines []string

	UPChan := make(chan string)
	RPChan := make(chan string)

	up := UppercaseProcessor{bisiv: UPChan}
	rp := ReverseProcessor{bisiv: RPChan}

	wg.Add(4)
	go savePlasePlease(UPChan, &wg, up, ListOfLines)
	go savePlasePlease(RPChan, &wg, rp, ListOfLines)
	go savePlase(UPChan, &wg, &UppercaseLines)
	go savePlase(RPChan, &wg, &ReverseLines)

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

	if err := bufiofiles.WriteLineExample("uppercase_output.txt", UppercaseLines); err != nil {
		fmt.Println("Error writing UppercaseLines to file:", err)
	}

	if err := bufiofiles.WriteLineExample("reverse_output.txt", ReverseLines); err != nil {
		fmt.Println("Error writing ReverseLines to file:", err)
	}
}
