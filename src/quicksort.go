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

var comparisions, xx uint

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

type  IndexValue struct {
   index int
       value int
      }

func medianPivot(array []int) (medianIndex int) {
  mid := (len(array) / 2)

  var (median IndexValue)
  first := IndexValue{0, array[0]}
  last := IndexValue{len(array) - 1, array[len(array) - 1]}

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

  // fmt.Println("arraylen", len(array), "sorted three", three)

  medianIndex = three[1].index

  return
}


func partition(array []int) (j int) {

  // First Pivot
//  pivot := 0

  // Last Pivot
  //pivot := len(array) - 1

  // Median Pivot
  pivot := medianPivot(array)



//  comparisions = comparisions + uint(len(array)) - 1

	pivotValue := array[pivot]

	array[0], array[pivot] = array[pivot], array[0]

	j = 1
	for i := 1; i < len(array); i++ {
    comparisions = comparisions + 1
		if array[i] < pivotValue {
			array[i], array[j] = array[j], array[i]
			j = j + 1
		}
	}
	array[0], array[j-1] = array[j-1], array[0]
	return
}

func quicksort(data []int) {
	if len(data) <= 1 {
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
//   s := [...]int{0,9,8,7,6,5,4,3,2,1}
	 //s, err := readLines("IntegerArray.txt")
	 s, err := readLines("/usr/local/google/home/joetoth/Downloads/datax.txt")

	if err != nil {
			log.Fatal(err)
		}
  log.Println("cat")
	quicksort(s[0:])
	// fmt.Println(s)
	fmt.Println(comparisions)
  //x := [...]int{-4, 3,1 ,5,1, -4}
  //medianPivot(x[0:])
}
