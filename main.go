package main

import (
	gochaninterface "github.com/SaYaku64/hillel/product/go_chan_interface"
	"strings"
	"sync"
)

type Processor interface {
	Process(data string) string
}

type UppercaseProcessor struct{}

func (up UppercaseProcessor) Process(data string) string {
	return strings.ToUpper(data)
}

type ReverseProcessor struct{}

func (rp ReverseProcessor) Process(data string) string {
	reverseData := []rune(data)
	for i, j := 0, len(reverseData)-1; i < j; i, j = i+1, j-1 {
		reverseData[i], reverseData[j] = reverseData[j], reverseData[i]
	}
	return string(reverseData)
}

//func processLines(lines []string, processor Processor, wg *sync.WaitGroup) {
//	for _, line := range lines {
//		go func(line string) {
//			defer wg.Done()
//			fmt.Println(processor.Process(line))
//		}(line)
//	}
//}

func main() {
	var wg sync.WaitGroup
	up := UppercaseProcessor{}
	pr := ReverseProcessor{}

	wg.Add(len(gochaninterface.ListOfLines))
	for _, line := range gochaninterface.ListOfLines {
		go func(line string) {
			defer wg.Done()
			upperLine := up.Process(line)
			gochaninterface.UppercaseLines = append(gochaninterface.UppercaseLines, upperLine)
		}(line)
	}
	wg.Wait()

	wg.Add(len(gochaninterface.ListOfLines))
	for _, line := range gochaninterface.ListOfLines {
		go func(line string) {
			defer wg.Done()
			reverseLine := pr.Process(line)
			gochaninterface.ReverseLines = append(gochaninterface.ReverseLines, reverseLine)
		}(line)
	}
	wg.Wait()

	//wg.Add(len(gochaninterface.ListOfLines))
	//processLines(gochaninterface.ListOfLines, up, &wg)
	//wg.Wait()
	//
	//wg.Add(len(gochaninterface.ListOfLines))
	//processLines(gochaninterface.ListOfLines, pr, &wg)
	//wg.Wait()
}
