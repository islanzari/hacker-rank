package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func max(arr []int32) int32 {
	max := arr[0]

	for _, v := range arr {

		if max < v {
			max = v
		}
	}
	return max
}

// Complete the hurdleRace function below.
func hurdleRace(k int32, height []int32) int32 {
	var potion int32

	if max(height) > k {

		potion = max(height) - k

	}
	return potion
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)

	//defer stdout.Close()

	writer := bufio.NewWriterSize(os.Stdout, 1024*1024)

	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	heightTemp := strings.Split(readLine(reader), " ")

	var height []int32

	for i := 0; i < int(n); i++ {
		heightItemTemp, err := strconv.ParseInt(heightTemp[i], 10, 64)
		checkError(err)
		heightItem := int32(heightItemTemp)
		height = append(height, heightItem)
	}

	result := hurdleRace(k, height)

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
