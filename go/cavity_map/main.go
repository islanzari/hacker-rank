package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func searchMatrixforCordinate(n []int) {

	for i := 1; i < len(n)-1; i++ {
		for j := 1; j < len(n)-1; j++ {
			if n[i-1][j] != 0 && n[i+1][j] != 0 && n[i][j-1] != 0 && n[i][j+1] != 0 {

			}
		}
	}
}

// Complete the cavityMap function below.
func cavityMap(grid []string) []string {

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	writer := bufio.NewWriterSize(os.Stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var grid []string

	for i := 0; i < int(n); i++ {
		gridItem := readLine(reader)
		grid = append(grid, gridItem)
	}

	result := cavityMap(grid)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
