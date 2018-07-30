package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the migratoryBirds function below.
func migratoryBirds(ar []int32) int32 {
	var fin int32
	var zmi int32
	most := make([]int32, 5)
	for i := 0; i < len(ar); i++ {

		switch {
		case ar[i] == 1:
			most[0]++
		case ar[i] == 2:
			most[1]++
		case ar[i] == 3:
			most[2]++
		case ar[i] == 4:
			most[3]++
		case ar[i] == 5:
			most[4]++
		}

	}
	for j := 0; j < len(most); j++ {
		if most[j] > zmi {
			zmi = most[j]
			fin = int32(j + 1)
		}
	}

	return fin

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	writer := bufio.NewWriterSize(os.Stdout, 1024*1024)

	arCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	arTemp := strings.Split(readLine(reader), " ")

	var ar []int32

	for i := 0; i < int(arCount); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := migratoryBirds(ar)

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
