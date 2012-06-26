package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	//	"math/rand"
	"os"
	"strconv"
)

var comparisions uint

// Read a whole file into the memory and store it as array of lines
func readLines(path string) (lines []int, err error) {
	var (
		file   *os.File
		part   []byte
		prefix bool
	)
	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 0))
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			intLine, _ := strconv.Atoi(buffer.String())
			lines = append(lines, intLine)
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}

func partition(array []int) (j int) {

	comparisions = comparisions + uint(len(array)) - 1

	//	pivot := rand.Int() % len(array)
	pivot := len(array) - 1

	//	fmt.Println("pivot...", pivot)
	pivotValue := array[pivot]
	//	fmt.Println("pivotvalue...", pivotValue)
	array[0], array[pivot] = array[pivot], array[0]
	//	fmt.Println("move to front...", array)

	j = 1
	for i := 1; i < len(array); i++ {
		if array[i] < pivotValue {
			//			fmt.Println("before swap...", array, "swapping", array[i], array[j])
			array[i], array[j] = array[j], array[i]
			j = j + 1
			//			fmt.Println("j...", j, "swapped...", array)
		}
	}
	//	fmt.Println("before final swap...", array)
	array[0], array[j-1] = array[j-1], array[0]
	//	fmt.Println("paritioned...", array)
	//os.Exit(0)
	return
}

func quicksort(data []int) {
	if len(data) <= 1 {
		return
	}

	//	fmt.Println("before parition...", data)
	newPivotIndex := partition(data)
	quicksort(data[0:newPivotIndex])
	begin := newPivotIndex
	end := len(data)
	quicksort(data[begin:end])
}

func main() {
	//	s := [...]int{1, 5, 8, 3, 2, 9, 0, 4, 6, 7}
	s, err := readLines("src/IntegerArray.txt")
	if err != nil {
		log.Fatal(err)
	}
	quicksort(s[0:])
	fmt.Println(comparisions)
}
