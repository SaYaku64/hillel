package main


import (
	"fmt"
	"strings"
	"sync"
)

var ListOfLines = []string{
	"Hello Everyone!",
	"Let's make THIS task",
	"a bit cooler than others",
	"YoU WiLL NeeD tO Make",
	"Simple BUT interesting program",
	"Do Your BEST",
	"GoOd LuCk!",
}

var (
	UppercaseLines = []string{}
	ReverseLines   = []string{}
)

// interface Processor
type Processor interface {
	Process(data string) string
}

type UppercaseProcessor struct {
}

func (p UppercaseProcessor) Process(data string) string {
	return strings.ToUpper(data)
}

type ReverseProcessor struct {
}

func (p ReverseProcessor) Process(data string) string {
	runes := []rune(data)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func readLines(lines []string, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, line := range lines {
		out <- line
	}
	close(out)
}

// Функція для обробки рядків
func processLines(processor Processor, in <-chan string, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for line := range in {
		out <- processor.Process(line)
	}
	close(out)
}

// Функція для запису оброблених рядків у змінні
func collectLines(in <-chan string, wg *sync.WaitGroup, result *[]string) {
	defer wg.Done()
	for line := range in {
		*result = append(*result, line)
	}
}

func main() {
	var wg sync.WaitGroup

	bufferSize := len(ListOfLines)

	lines := make(chan string, bufferSize)
	uppercaseLines := make(chan string, bufferSize)
	reverseLines := make(chan string, bufferSize)

	// Запуск горутини для читання рядків зі змінної
	wg.Add(1)
	go readLines(ListOfLines, lines, &wg)

	// Запуск горутин для обробки рядків
	wg.Add(1)
	go processLines(UppercaseProcessor{}, lines, uppercaseLines, &wg)

	wg.Add(1)
	go processLines(ReverseProcessor{}, lines, reverseLines, &wg)

	// Запуск горутин для збору оброблених рядків у змінні
	wg.Add(1)
	go collectLines(uppercaseLines, &wg, &UppercaseLines)

	wg.Add(1)
	go collectLines(reverseLines, &wg, &ReverseLines)

	wg.Wait()

	// Виведення результатів для перевірки
	fmt.Println("Uppercase Lines:")
	for _, line := range UppercaseLines {
		fmt.Println(line)
	}

	fmt.Println("\nReverse Lines:")
	for _, line := range ReverseLines {
		fmt.Println(line)
	}
}
