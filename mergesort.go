package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

var inversions uint

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

func splitArray(array []int) (a1 []int, a2 []int) {
	mid := (len(array) / 2)

	a1 = make([]int, mid)
	copy(a1, array[0:mid])

	if len(array)%2 == 0 {
		a2 = make([]int, mid)
	} else {
		a2 = make([]int, mid+1)
	}

	copy(a2, array[mid:])

	return
}

func xsort(numbers []int) (ret []int) {

	ret = make([]int, len(numbers))

	a1, a2 := splitArray(numbers)

	if len(a1) > 1 {
		a1 = xsort(a1)
	}

	if len(a2) > 1 {
		a2 = xsort(a2)
	}

	// merge
	i := 0
	j := 0

	//5305 6711 6804
	for k := 0; k < len(numbers); k++ {

		if i == len(a1) {
			ret[k] = a2[j]
			j++
			continue
		}

		if j == len(a2) {
			ret[k] = a1[i]
			i++
			continue
		}

		if a1[i] < a2[j] {
			ret[k] = a1[i]
			i++
		} else {
			inversions += uint(len(a1) - i)
			ret[k] = a2[j]
			j++
		}

	}

	return
}

func main() {
	lines, err := readLines("src/IntegerArray.txt")
	if err != nil {
		log.Fatal(err)
	}

	//	lines := []int{1, 3, 5, 2, 4, 6}
	//
	//	a1, a2 := splitArray(lines)
	//
	//	fmt.Println(a1, a2)
	started := time.Now()
	fmt.Println(started.UnixNano())
	lines = xsort(lines)
	duration := (time.Now().UnixNano() - started.UnixNano()) / 1000000

	//	fmt.Println("LINES", lines)
	fmt.Println("Duration", duration)
	fmt.Println("inversions", inversions)

}
