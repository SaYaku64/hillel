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

//func ReadLines(filename string) (out []string, err error) {
//	file, err := os.Open(filename)
//	if err != nil {
//		return []string{}, err
//	}
//	defer file.Close()
//
//	scanner := bufio.NewScanner(file)
//	for scanner.Scan() {
//		out = append(out, scanner.Text())
//	}
//
//	return out, err
//}
//
//func writeLineExample(filename string, lines []string) error {
//	file, err := os.Create(filename)
//	if err != nil {
//		return err
//	}
//	defer file.Close()
//
//	writer := bufio.NewWriter(file)
//
//	for _, line := range lines {
//		fmt.Fprintln(writer, line)
//	}
//	writer.Flush()
//	return nil
//}

func main() {
	var wg sync.WaitGroup

	//ListOfLines, err := ReadLines("input.txt")
	//if err != nil {
	//	fmt.Println("Error reading file:", err)
	//	return
	//}
	//var UppercaseLines []string
	//var ReverseLines []string

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

	//if err := writeLineExample("uppercase_output.txt", UppercaseLines); err != nil {
	//	fmt.Println("Error writing UppercaseLines to file:", err)
	//}
	//
	//if err := writeLineExample("reverse_output.txt", ReverseLines); err != nil {
	//	fmt.Println("Error writing ReverseLines to file:", err)
	//}
}
