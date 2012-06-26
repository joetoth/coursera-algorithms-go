package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
		"log"
	"math/rand"
	"os"
	"strconv"
	"time"
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

func medianPivot(array []int) (medianIndex int) {
  mid := (len(array) / 2)
  

var IndexValue struct {
  index int
  value int
}

var (first, median, last IndexValue)

  if len(array)%2 == 0 {
    median.value = array[mid-1]
    median.index = mid-1
  } else {
    median.value = array[mid]
    median.index = mid-1
  }

  three :=[...]IndexValue{first,median,last}
  // sort
  for i := 0; i < 2; i++ {
    if three[i].value > three[i+1].value {
      three[i], three[i+1] = three[i+1], three[i]
    }
  }

  if three[1].value < three[0].value {
      three[0], three[1] = three[1], three[0]
  }

  fmt.Println("sorted three", three)

  medianIndex = three[1].index

  return
}


func partition(array []int) (j int) {
  comparisions = comparisions + uint(len(array)) - 1
	// r := rand.Int()
	//pivot := r % len(array)
  pivot := medianPivot(array)
	pivotValue := array[pivot]
	//fmt.Println("lenarray", len(array), "pivot", pivot, "pivotValue", pivotValue)
  //os.Exit(1)
	//	pivot := len(array) - 1

	//	fmt.Println("pivot...", pivot)

	//	fmt.Println("pivotvalue...", pivotValue)
	array[0], array[pivot] = array[pivot], array[0]
	//	fmt.Println("move to front...", array)

	j = 1
	for i := 1; i < len(array); i++ {
		if array[i] < pivotValue {
			//			fmt.Println("before swap...", array, "swapping", array[i], array[j])
			array[i], array[j] = array[j], array[i]
			//fmt.Println( "i...", i, "j...", j, "swapped...", array)
			j = j + 1
		}
	}
		//fmt.Println("before final swap...", array)
	array[0], array[j-1] = array[j-1], array[0]
		//fmt.Println("paritioned...", array)
	//os.Exit(0)
	return
}

func quicksort(data []int) {
	if len(data) <= 2 {
		return
	}

	// fmt.Println("before parition...", data)
	newPivotIndex := partition(data)
   //fmt.Println("beg0", "end", newPivotIndex- 1)
	quicksort(data[0:newPivotIndex - 1])
		begin := newPivotIndex
		end := len(data)
    //fmt.Println("begin", begin, "end", end)
		quicksort(data[begin:end])
}

func main() {
	rand.Seed(time.Now().Unix())
	 s := [...]int{1, 5, 8, 3, 2, 9, 0, 7, 6, 4}
	//	s, err := readLines("IntegerArray.txt")
	//if err != nil {
	//		log.Fatal(err)
	//	}
  log.Println("cat")
	quicksort(s[0:])
	fmt.Println(s)
	fmt.Println(comparisions)
  //x := [...]int{-4, 3,1 ,5,1, -4}
  //medianPivot(x[0:])
}
