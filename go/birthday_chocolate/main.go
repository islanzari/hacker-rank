package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the solve function below.
func solve(s []int32, d int32, m int32) int32 {
	var count int32
	var sum int32
	var countD int32

	for i := 0; i < len(s); i++ {
		//	fmt.Println(i)
		count++
		sum += s[i]
		if m == count {
			if sum == d {
				countD++
				sum = 0
				count = 0
			}
			i = i - (int(m) - 1)
			count = 0
			sum = 0
			//fmt.Println(i)

		}

	}
	return countD
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)

	//defer stdout.Close()

	writer := bufio.NewWriterSize(os.Stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	sTemp := strings.Split(readLine(reader), " ")

	var s []int32

	for i := 0; i < int(n); i++ {
		sItemTemp, err := strconv.ParseInt(sTemp[i], 10, 64)
		checkError(err)
		sItem := int32(sItemTemp)
		s = append(s, sItem)
	}

	dm := strings.Split(readLine(reader), " ")

	dTemp, err := strconv.ParseInt(dm[0], 10, 64)
	checkError(err)
	d := int32(dTemp)

	mTemp, err := strconv.ParseInt(dm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	result := solve(s, d, m)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
